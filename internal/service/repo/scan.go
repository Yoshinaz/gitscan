package repo

import (
	"errors"
	"fmt"
	"github.com/gitscan/constants/reqError"
	"github.com/gitscan/constants/status"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/report"
	"github.com/rs/zerolog/log"
	gitLib "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gorm.io/gorm"
	"time"
)

// Scan for a request github repository url to find a secret in all commit history
// don't scan for the duplicate url with the same head commit
func Scan(r Interface, db database.DB, working chan bool) (string, error) {
	url := r.URL()
	name := r.Name()

	l := log.With().Str("url", url).Logger()
	l.Info().Msg("scan begin")

	if r.Repo() == nil {
		err := r.Clone()

		if err != nil {
			l.Error().Msg("url might be wrong")

			return reqError.URL_ERROR.String(), err
		}
	}
	headCommit, err := r.GetHeadCommitHash()
	if err != nil {
		l.Error().Msg("could not retrieve head")

		return reqError.HEAD_ERROR.String(), err
	}

	// if the scan information of the latest commit exist, don't need to rescan again
	// if it was failed will try to rescan again
	infoDB, err := db.Info().Find(database.Info{URL: url})
	if err == nil && headCommit == infoDB.Commit && infoDB.Status != status.FAILED.String() {
		l.Info().Msgf("the scan information exist with latest commit %s", infoDB.Commit)

		return infoDB.Status, nil
	}

	//database error
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Error().Msg("database error: could not query info db")

		return reqError.INTERNAL_ERROR.String(), err
	}

	scanAllCommit := false
	//create a new info record for the first time
	if err == gorm.ErrRecordNotFound {
		infoDB, err = db.Info().Create(database.Info{
			Name:        name,
			URL:         url,
			Status:      status.QUEUED.String(),
			Commit:      headCommit,
			Description: "",
			EnqueuedAt:  time.Now(),
			StartedAt:   nil,
			FinishedAt:  nil,
		})
		scanAllCommit = true
	}

	if err != nil {
		l.Error().Msg("database error: could not create info db")

		return reqError.INTERNAL_ERROR.String(), err
	}

	l.Info().Msgf("enqueued for the latest commit %s", infoDB.Commit)
	infoDB.Status = status.QUEUED.String()
	// create go routine to wait until a free slot exist
	go func() {
		working <- true
		go process(r, infoDB, db, scanAllCommit, working)
	}()

	return infoDB.Status, nil
}

// process scan a request github repository url
// finding for a secret from commit history
func process(r Interface, infoDB database.Info, db database.DB, scanAllCommit bool, working chan bool) {
	url := r.URL()

	l := log.With().Str("url", url).Logger()
	l.Info().Msg("process begin")
	currentTime := time.Now()
	infoDB.StartedAt = &currentTime
	// ... retrieves the branch pointed by HEAD
	ref, err := r.Repo().Head()
	if err != nil {
		l.Error().Msg("could not retrieve head")
		processFailure(infoDB, db, "could not retrieve head commit")

		return
	}

	// ... retrieves the commit history
	cIter, err := r.Repo().Log(&gitLib.LogOptions{From: ref.Hash()})
	if err != nil {
		l.Error().Msg("retrieve commit history failed")
		processFailure(infoDB, db, "retrieve commit history failed")

		return
	}
	infoDB.Status = status.INPROGRESS.String()
	db.Info().Update(infoDB)

	skip := false
	// iterates over the commits
	err = cIter.ForEach(func(c *object.Commit) error {
		// stop when meet the latest scan commit
		if c.Hash.String() == infoDB.Commit {
			skip = true
		}
		if !skip || scanAllCommit {
			findings, err := r.Rules().Process(c)
			if err != nil {
				errWrap := errors.New(fmt.Sprintf("rules processing error: %s", err.Error()))
				l.Error().Msgf(errWrap.Error())

				return errWrap
			}

			err = report.SaveFindings(infoDB.ID, c.Hash.String(), findings, db)
			if err != nil {
				errWrap := errors.New(fmt.Sprintf("could not save finding result: %s", err.Error()))
				l.Error().Msgf("could not save finding result: %s", err.Error())

				return errWrap
			}
		}
		return nil
	})

	infoDB.Status = status.SUCCESS.String()
	if err != nil {
		infoDB.Status = status.FAILED.String()
		infoDB.Description = err.Error()

		l.Warn().Msgf("process failed: %s", err.Error())
	}

	currentTime = time.Now()
	infoDB.FinishedAt = &currentTime
	infoDB.Commit = ref.Hash().String()
	db.Info().Update(infoDB)

	l.Info().Msg("process finish")
	<-working
}

func processFailure(reportInfo database.Info, db database.DB, description string) {
	reportInfo.Status = status.FAILED.String()
	currentTime := time.Now()
	reportInfo.FinishedAt = &currentTime
	reportInfo.Description = description
	db.Info().Update(reportInfo)
}
