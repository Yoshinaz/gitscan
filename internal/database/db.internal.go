package database

import "gorm.io/gorm"

func (d db) Ping() error {
	sqlDb, err := d.RawDB().DB()
	if err != nil {
		return err
	}

	return sqlDb.Ping()
}

func (d db) UseTransaction(fn TransactionFunction) error {
	return d.gorm.Transaction(func(tx *gorm.DB) error {
		return fn(FromDB(tx))
	})
}
