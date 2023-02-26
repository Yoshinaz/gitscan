//go:build integration_test
// +build integration_test

package database_test

import (
	"github.com/gitscan/internal/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFinding_Create(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("create", func(t *testing.T) {
		finding := database.Finding{
			Model:  database.Model{},
			InfoID: "123",
			Type:   "",
			RuleID: "",
			Commit: "",
		}
		result, err := db.Finding().Create(finding)

		assert.Nil(t, err)
		assert.Equal(t, "123", result.InfoID)
	})
}

func TestFinding_Creates(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("creates", func(t *testing.T) {
		findings := []database.Finding{
			{
				Model:  database.Model{},
				InfoID: "123",
				Type:   "",
				RuleID: "",
				Commit: "",
			},
			{
				Model:  database.Model{},
				InfoID: "124",
				Type:   "",
				RuleID: "",
				Commit: "",
			},
		}
		result, err := db.Finding().Creates(findings)

		assert.Nil(t, err)
		assert.Equal(t, "123", result[0].InfoID)
		assert.Equal(t, "124", result[1].InfoID)
	})
}

func TestFinding_FindByInfoID(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	infoID := "123"
	t.Run("create", func(t *testing.T) {
		findings := []database.Finding{
			{
				Model:  database.Model{},
				InfoID: infoID,
				Type:   "",
				RuleID: "",
				Commit: "",
			},
			{
				Model:  database.Model{},
				InfoID: infoID,
				Type:   "",
				RuleID: "",
				Commit: "",
			},
		}
		_, err := db.Finding().Creates(findings)

		result, err := db.Finding().FindByInfoID(infoID)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})
}

func TestFinding_FindByInfoIDAndCommit(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	infoID := "123"
	t.Run("create", func(t *testing.T) {
		findings := []database.Finding{
			{
				Model:  database.Model{},
				InfoID: infoID,
				Type:   "",
				RuleID: "1",
				Commit: "1234",
			},
			{
				Model:  database.Model{},
				InfoID: infoID,
				Type:   "",
				RuleID: "2",
				Commit: "1235",
			},
		}
		_, err := db.Finding().Creates(findings)

		result, err := db.Finding().FindByInfoIDAndCommit(infoID, "1234")
		assert.Nil(t, err)
		assert.Equal(t, "1", result.RuleID)
	})
}
