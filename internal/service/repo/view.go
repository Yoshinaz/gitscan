package repo

import (
	"github.com/gitscan/constants/reqError"
	"github.com/gitscan/constants/status"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/report"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strings"
)

// ViewReport retrieve the repository's scan information
func ViewReport(r Interface, db database.DB) (report.Info, error) {
	name := r.Name()
	url := r.URL()
	l := log.With().Str("url", url).Logger()
	l.Info().Msg("view report begin")
	infoDB, err := db.Info().Find(database.Info{URL: url})
	if err != nil {
		l.Warn().Msgf("database error: %s", err.Error())

		return report.Info{
			Name:        name,
			URL:         url,
			Status:      reqError.URL_ERROR.String(),
			Description: "URL might be wrong or didn't scan yet",
		}, err
	}

	if infoDB.Status != status.SUCCESS.String() {
		return report.Info{
			Name:        infoDB.Name,
			URL:         infoDB.URL,
			Status:      infoDB.Status,
			EnqueuedAt:  infoDB.EnqueuedAt,
			StartedAt:   infoDB.StartedAt,
			FinishedAt:  infoDB.FinishedAt,
			CreatedAt:   infoDB.CreatedAt,
			Description: infoDB.Description,
			Findings:    []report.Finding{},
		}, nil
	}

	l.Info().Msg("get finding information")
	findingDB, err := db.Finding().FindByInfoID(infoDB.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Warn().Msgf("database error: %s", err.Error())

		return report.Info{
			Name:        name,
			URL:         url,
			Status:      reqError.INTERNAL_ERROR.String(),
			Description: "INTERNAL ERROR",
		}, err
	}

	l.Info().Msg("get location for each finding information")
	findings, err := mappingFinding(r, findingDB, db)
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Warn().Msgf("database error: %s", err.Error())

		return report.Info{
			Name:        name,
			URL:         url,
			Status:      reqError.INTERNAL_ERROR.String(),
			Description: "INTERNAL ERROR",
		}, err
	}

	return report.Info{
		Name:        infoDB.Name,
		URL:         infoDB.URL,
		Status:      infoDB.Status,
		EnqueuedAt:  infoDB.EnqueuedAt,
		StartedAt:   infoDB.StartedAt,
		FinishedAt:  infoDB.FinishedAt,
		CreatedAt:   infoDB.CreatedAt,
		Description: infoDB.Description,
		Findings:    findings,
	}, nil
}

// mappingFinding mapping finding database to report.Finding object
func mappingFinding(r Interface, findingDB []database.Finding, db database.DB) ([]report.Finding, error) {
	findings := make([]report.Finding, 0)
	for _, v := range findingDB {
		locationDB, err := db.Location().FindByFindingID(v.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return findings, err
		}

		finding := report.Finding{
			Type:     v.Type,
			RuleId:   v.RuleID,
			Commit:   v.Commit,
			Metadata: getMetadata(r, v.RuleID),
		}
		if locationDB != nil {
			location := mappingLocation(locationDB)
			finding.Location = location
		}
		findings = append(findings, finding)
	}

	return findings, nil
}

// getMetadata get description and severity for ruleID
func getMetadata(r Interface, ruleID string) report.Metadata {
	return r.Rules().GetMetaData(ruleID)
}

// mappingLocation mapping location database to report.Location object
func mappingLocation(locationDB []database.Location) []report.Location {
	location := make([]report.Location, 0)
	for _, v := range locationDB {
		lines := strings.Split(v.Lines, ",")
		begins := make([]report.Begin, 0)
		for _, v := range lines {
			begins = append(begins, report.Begin{Line: v})
		}
		l := report.Location{
			Path: v.Path,
			Positions: report.Position{
				Begin: begins,
			},
		}
		location = append(location, l)
	}

	return location
}
