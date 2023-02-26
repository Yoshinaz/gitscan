//go:build integration_test
// +build integration_test

package database_test

import (
	"github.com/gitscan/internal/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocation_Create(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("create", func(t *testing.T) {
		location := database.Location{
			Model:     database.Model{},
			FindingID: "123",
			Path:      "",
			Lines:     "",
			Status:    "",
		}
		result, err := db.Location().Create(location)

		assert.Nil(t, err)
		assert.Equal(t, "123", result.FindingID)
	})
}

func TestLocation_Creates(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("creates", func(t *testing.T) {
		locations := []database.Location{
			{
				Model:     database.Model{},
				FindingID: "123",
				Path:      "",
				Lines:     "",
				Status:    "",
			},
			{
				Model:     database.Model{},
				FindingID: "124",
				Path:      "",
				Lines:     "",
				Status:    "",
			},
		}
		result, err := db.Location().Creates(locations)

		assert.Nil(t, err)
		assert.Equal(t, "123", result[0].FindingID)
		assert.Equal(t, "124", result[1].FindingID)
	})
}

func TestLocation_FindByFindingID(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	infoID := "123"
	t.Run("create", func(t *testing.T) {
		locations := []database.Location{
			{
				Model:     database.Model{},
				FindingID: "123",
				Path:      "",
				Lines:     "",
				Status:    "",
			},
			{
				Model:     database.Model{},
				FindingID: "123",
				Path:      "",
				Lines:     "",
				Status:    "",
			},
		}
		_, err := db.Location().Creates(locations)

		result, err := db.Location().FindByFindingID(infoID)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})
}
