package git

import (
	"github.com/rs/zerolog/log"
	gitLib "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gorm.io/gorm"
	"guardrail/gitscan/constants/status"
	"guardrail/gitscan/internal/database"
	"guardrail/gitscan/internal/report"
	"guardrail/gitscan/rules"
	"time"
)

func (r *Repo) Scan(db database.DB, working chan bool) (string, error) {
	if r.Repo == nil {
		r.Clone()
	}
	ref, err := r.Repo.Head()
	if err != nil {
		return "", err
	}
	//if the scan information of the latest commit exist, don't need to rescan again
	infoDB, err := r.GetInfoDB(db)

	if err == nil && ref.Hash().String() == infoDB.Commit {
		return infoDB.Status, nil
	}

	//database error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	scanAllCommit := false
	//create a new info record for the first time
	if err == gorm.ErrRecordNotFound {
		infoDB, err = db.Info().Create(database.Info{
			Name:        r.Name,
			URL:         r.URL,
			Status:      status.QUEUED.String(),
			Commit:      ref.Hash().String(),
			Description: "",
			EnqueuedAt:  time.Now(),
			StartedAt:   nil,
			FinishedAt:  nil,
		})
		scanAllCommit = true
	}

	if err != nil {
		return "", err
	}
	working <- true
	//go process(*r, infoDB, db, working)
	process(*r, infoDB, db, scanAllCommit, working)

	return infoDB.Status, nil
}

func process(repo Repo, infoDB database.Info, db database.DB, scanAllCommit bool, working chan bool) {
	log.Info().Msgf("process begin for %s", repo.Name)
	currentTime := time.Now()
	infoDB.StartedAt = &currentTime
	// ... retrieves the branch pointed by HEAD
	ref, err := repo.Repo.Head()
	if err != nil {
		processFailure(infoDB, db, "could not retrieve head commit")
		return
	}

	// ... retrieves the commit history
	cIter, err := repo.Repo.Log(&gitLib.LogOptions{From: ref.Hash()})
	if err != nil {
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
			findings, err := rules.Process(c)
			if err != nil {
				return err
			}

			err = report.SaveFindings(infoDB.ID, c.Hash.String(), findings, db)
			if err != nil {
				return err
			}
		}
		return nil
	})

	infoDB.Status = status.SUCCESS.String()
	if err != nil {
		infoDB.Status = status.FAILED.String()
	}

	currentTime = time.Now()
	infoDB.FinishedAt = &currentTime
	infoDB.Commit = ref.Hash().String()
	db.Info().Update(infoDB)

	<-working
}

func processFailure(reportInfo database.Info, db database.DB, description string) {
	reportInfo.Status = status.FAILED.String()
	currentTime := time.Now()
	reportInfo.FinishedAt = &currentTime
	reportInfo.Description = description
	db.Info().Update(reportInfo)
}

func (r *Repo) GetInfoDB(db database.DB) (database.Info, error) {
	infoDB, err := db.Info().FindByURL(r.URL, status.SUCCESS)
	if err != nil {
		return database.Info{}, err
	}

	return infoDB, err
}
