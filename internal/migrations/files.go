package migrations

import (
	"embed"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

//go:embed sql/*.sql
var schemaFs embed.FS

func New(db *gorm.DB, noLock bool) (*migrate.Migrate, error) {
	dbDriver, err := getDbDriver(db, noLock)
	if err != nil {
		return nil, err
	}
	d, err := iofs.New(schemaFs, "sql")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return migrate.NewWithInstance("iofs", d, "postgres", dbDriver)
}
