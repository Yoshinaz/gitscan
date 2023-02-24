package git

import (
	"gorm.io/gorm"
	"guardrail/gitscan/constants/status"
	"guardrail/gitscan/internal/database"
	"guardrail/gitscan/internal/report"
	"guardrail/gitscan/rules"
	"strings"
)

func (r *Repo) ViewReport(db database.DB) (report.Info, error) {
	infoDB, err := db.Info().FindByURL(r.URL, status.SUCCESS)
	if err != nil {
		return report.Info{}, err
	}
	findingDB, err := db.Finding().FindByInfoID(infoDB.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return report.Info{}, err
	}
	findings := make([]report.Finding, 0)
	for _, v := range findingDB {
		locationDB, _ := db.Location().FindByFindingID(v.ID)
		finding := report.Finding{
			Type:     v.Type,
			RuleId:   v.RuleID,
			Commit:   v.Commit,
			Metadata: getMetadata(v.RuleID),
		}
		if locationDB != nil {
			location := toLocation(locationDB)
			finding.Location = location
		}
		findings = append(findings, finding)
	}

	return report.Info{
		Name:        infoDB.Name,
		URL:         infoDB.URL,
		Status:      infoDB.Status,
		EnqueuedAt:  infoDB.EnqueuedAt,
		StartedAt:   *infoDB.StartedAt,
		FinishedAt:  *infoDB.FinishedAt,
		CreatedAt:   infoDB.CreatedAt,
		Description: infoDB.Description,
		Findings:    findings,
	}, nil
}

func getMetadata(ruleID string) report.Metadata {
	return rules.GetMetaData(ruleID)
}

func toLocation(locationDB []database.Location) []report.Location {
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
