package report

import (
	"github.com/gitscan/constants/status"
	"github.com/gitscan/internal/database"
	"gorm.io/gorm"
	"strings"
)

func SaveFindings(infoID string, commit string, findings []Finding, db database.DB) error {
	dbFindings := make([]database.Finding, 0)
	for _, v := range findings {
		if isFindingExist(infoID, commit, db) {
			continue
		}

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

func isFindingExist(infoID string, commit string, db database.DB) bool {
	_, err := db.Finding().FindByInfoIDAndCommit(infoID, commit)
	if err != gorm.ErrRecordNotFound {
		return true
	}

	return false
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
	if len(dbLocations) == 0 {
		return nil
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
