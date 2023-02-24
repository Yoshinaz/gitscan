package report

import (
	"guardrail/gitscan/constants/status"
	"guardrail/gitscan/internal/database"
	"strings"
)

func SaveFindings(infoID string, commit string, findings []Finding, db database.DB) error {
	dbFindings := make([]database.Finding, 0)
	for _, v := range findings {
		record := database.Finding{
			InfoID: infoID,
			Type:   v.Type,
			RuleID: v.RuleId,
			Commit: commit,
		}

		dbFindings = append(dbFindings, record)
	}

	_, err := db.Finding().Creates(dbFindings)
	if err != nil {
		return err
	}

	err = saveLocation(infoID, commit, findings, db)

	return err
}

func saveLocation(infoID string, commit string, findings []Finding, db database.DB) error {
	dbLocations := make([]database.Location, 0)
	errCount := 0
	for _, finding := range findings {
		f, err := db.Finding().FindByInfoIDAndCommit(infoID, commit)
		if err != nil {
			errCount++
			continue
		}

		for _, location := range finding.Location {
			lines := toLines(location.Positions.Begin)
			record := database.Location{
				FindingID: f.ID,
				Path:      location.Path,
				Lines:     lines,
				Status:    status.ACTIVE.String(),
			}

			dbLocations = append(dbLocations, record)
		}
	}

	_, err := db.Location().Creates(dbLocations)

	return err
}

func toLines(begins []Begin) string {
	tmp := make([]string, 0)
	for _, v := range begins {
		tmp = append(tmp, v.Line)
	}

	return strings.Join(tmp, ",")
}
