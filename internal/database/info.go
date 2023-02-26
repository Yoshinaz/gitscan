package database

import (
	"github.com/gitscan/constants/status"
	"time"

	"gorm.io/gorm"
)

type Info struct {
	Model
	Name        string
	URL         string
	RuleSet     string
	Status      string
	Commit      string
	Description string
	AllCommit   string
	EnqueuedAt  time.Time
	StartedAt   *time.Time
	FinishedAt  *time.Time
}

type InfoInterface interface {
	Create(input Info) (Info, error)
	Update(input Info) (Info, error)
	Find(filter Info) (Info, error)
	FindRecoveryInfo() ([]Info, error)
}

type info struct {
	gorm *gorm.DB
}

func (c info) Create(input Info) (Info, error) {
	tx := c.gorm.Create(&input)

	return input, tx.Error
}

func (c info) Update(input Info) (Info, error) {
	tx := c.gorm.Model(&Info{}).Where(&Info{URL: input.URL}).Updates(input)

	return input, tx.Error
}

func (c info) Find(filter Info) (Info, error) {
	var models Info
	tx := c.gorm.Model(&Info{}).Where(&filter).First(&models)

	return models, tx.Error
}

func (c info) FindRecoveryInfo() ([]Info, error) {
	var models []Info
	tx := c.gorm.Where("status IN ?", []string{status.QUEUED.String(), status.INPROGRESS.String()}).Find(&models)

	return models, tx.Error
}

func (d db) Info() InfoInterface {
	return info(d)
}
