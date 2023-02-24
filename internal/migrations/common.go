package migrations

import (
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"gorm.io/gorm"
)

func getDbDriver(db *gorm.DB, noLock bool) (database.Driver, error) {
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	return mysql.WithInstance(sqlDb, &mysql.Config{NoLock: noLock})
}
