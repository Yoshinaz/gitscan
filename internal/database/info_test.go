//go:build integration_test
// +build integration_test

package database_test

import (
	"github.com/gitscan/constants/status"
	"github.com/gitscan/internal/database"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInfo_Create(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("create", func(t *testing.T) {
		info := database.Info{
			Model:       database.Model{},
			Name:        "Test",
			URL:         "",
			RuleSet:     "",
			Status:      "",
			Commit:      "",
			Description: "",
			AllCommit:   "",
			EnqueuedAt:  time.Now(),
			StartedAt:   nil,
			FinishedAt:  nil,
		}
		result, err := db.Info().Create(info)

		assert.Nil(t, err)
		assert.Equal(t, "Test", result.Name)
	})
}

func TestInfo_Update(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("update", func(t *testing.T) {
		info := database.Info{
			Model: database.Model{
				ID: "1234",
			},
			Name:        "Test",
			URL:         "testurl",
			RuleSet:     "",
			Status:      "",
			Commit:      "",
			Description: "",
			AllCommit:   "",
			EnqueuedAt:  time.Now(),
			StartedAt:   nil,
			FinishedAt:  nil,
		}
		result, err := db.Info().Create(info)
		assert.Nil(t, err)
		assert.Equal(t, "Test", result.Name)

		info.Name = "Test2"
		result, err = db.Info().Update(info)
		assert.Nil(t, err)
		assert.Equal(t, "Test2", result.Name)
	})
}

func TestInfo_Find(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("find", func(t *testing.T) {
		info := database.Info{
			Model:       database.Model{},
			Name:        "Test",
			URL:         "Test url",
			RuleSet:     "",
			Status:      "Success",
			Commit:      "",
			Description: "",
			AllCommit:   "",
			EnqueuedAt:  time.Now(),
			StartedAt:   nil,
			FinishedAt:  nil,
		}
		_, err := db.Info().Create(info)

		result, err := db.Info().Find(database.Info{URL: "Test url"})
		assert.Nil(t, err)
		assert.Equal(t, "Test", result.Name)
	})
}

func TestInfo_FindRecoveryInfo(t *testing.T) {
	tx := MustGetRawDB(t).Begin()
	defer tx.Rollback()
	db := database.FromDB(tx)
	t.Run("create", func(t *testing.T) {
		info := database.Info{
			Model:      database.Model{},
			Name:       "Test",
			Status:     status.SUCCESS.String(),
			EnqueuedAt: time.Now(),
			StartedAt:  nil,
			FinishedAt: nil,
		}
		_, err := db.Info().Create(info)

		info = database.Info{
			Model:      database.Model{},
			Name:       "Test",
			URL:        "Test url",
			Status:     status.QUEUED.String(),
			EnqueuedAt: time.Now(),
			StartedAt:  nil,
			FinishedAt: nil,
		}
		_, err = db.Info().Create(info)

		info = database.Info{
			Model:      database.Model{},
			Name:       "Test",
			URL:        "Test url",
			Status:     status.INPROGRESS.String(),
			EnqueuedAt: time.Now(),
			StartedAt:  nil,
			FinishedAt: nil,
		}
		_, err = db.Info().Create(info)

		info = database.Info{
			Model:      database.Model{},
			Name:       "Test",
			URL:        "Test url",
			Status:     status.FAILED.String(),
			EnqueuedAt: time.Now(),
			StartedAt:  nil,
			FinishedAt: nil,
		}
		_, err = db.Info().Create(info)

		result, err := db.Info().FindRecoveryInfo()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})
}
