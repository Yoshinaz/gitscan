package database_test

import (
	"github.com/gitscan/config"
	"github.com/gitscan/internal/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func MustGetRawDB(t *testing.T) *gorm.DB {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Failed()
	}
	db, err := database.New(cfg.DB)
	if err != nil {
		t.Failed()
	}

	return db.RawDB()
}

func TestDb_Ping(t *testing.T) {
	tx := MustGetRawDB(t)
	db := database.FromDB(tx)
	err := db.Ping()
	assert.NoError(t, err)
}
