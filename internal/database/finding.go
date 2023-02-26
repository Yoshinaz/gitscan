package database

import (
	"gorm.io/gorm"
)

type Finding struct {
	Model
	InfoID string
	Type   string
	RuleID string
	Commit string
}

type FindingInterface interface {
	Create(input Finding) (Finding, error)
	Creates(input []Finding) ([]Finding, error)
	FindByInfoID(infoID string) ([]Finding, error)
	FindByInfoIDAndCommit(infoID, commit string) (Finding, error)
}

type finding struct {
	gorm *gorm.DB
}

func (c finding) Create(input Finding) (Finding, error) {
	tx := c.gorm.Create(&input)

	return input, tx.Error
}

func (c finding) Creates(inputs []Finding) ([]Finding, error) {
	if len(inputs) == 0 {
		return inputs, nil
	}
	
	tx := c.gorm.Create(&inputs)

	return inputs, tx.Error
}

func (c finding) FindByInfoID(infoID string) ([]Finding, error) {
	var models []Finding
	tx := c.gorm.Model(&Finding{}).Where(Finding{InfoID: infoID}).Find(&models)

	return models, tx.Error
}

func (c finding) FindByInfoIDAndCommit(infoID, commit string) (Finding, error) {
	var model Finding
	tx := c.gorm.Model(&Finding{}).Where(Finding{InfoID: infoID, Commit: commit}).First(&model)

	return model, tx.Error
}

func (d db) Finding() FindingInterface {
	return finding(d)
}
