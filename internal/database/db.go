package database

import (
	"fmt"
	"github.com/gitscan/config"
	"github.com/gitscan/internal/utils/ulid"
	"github.com/rs/zerolog/log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TransactionFunction func(DB) error

var database *db

type DB interface {
	RawDB() *gorm.DB
	UseTransaction(fn TransactionFunction) error
	Ping() error
	Info() InfoInterface
	Finding() FindingInterface
	Location() LocationInterface
}

type Model struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *Model) BeforeUpdate(g *gorm.DB) error {
	d.UpdatedAt = time.Now()
	g.Statement.Omit("created_at")
	g.Statement.Omit("id")

	return nil
}

type db struct {
	gorm *gorm.DB
}

func (d *Model) BeforeCreate(g *gorm.DB) error {
	name := g.Statement.Table
	if d.ID != "" {
		return nil
	}
	d.ID = ulid.New(name)

	return nil
}

func (d db) RawDB() *gorm.DB {
	return d.gorm
}

func FromDB(d *gorm.DB) DB {
	return db{gorm: d}
}

func New(c config.Database) (DB, error) {
	if database != nil {
		return database, nil
	}

	if err := newDatabase(c); err != nil {
		return nil, err
	}

	return database, nil
}

func newDatabase(dbConfig config.Database) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	log.Info().Msgf("connecting to %v db host %v port %v", dbConfig.Name, dbConfig.Host, dbConfig.Port)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//	Logger:         logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return err
	}

	if database == nil {
		database = &db{
			gorm: gormDB,
		}

		return database.Ping()
	}

	internal, err := database.RawDB().DB()
	if err != nil {
		log.Warn().Msgf("cannot get the raw db instance %v", err.Error())
	}
	if internal != nil {
		_ = internal.Close()
	}
	*database = db{
		gorm: gormDB,
	}

	return database.Ping()
}
