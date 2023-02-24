package database

import "gorm.io/gorm"

type Location struct {
	Model
	FindingID string
	Path      string
	Lines     string
	Status    string
}

type LocationInterface interface {
	Create(input Location) (Location, error)
	Creates(input []Location) ([]Location, error)
	FindByFindingID(findingID string) ([]Location, error)
}

type location struct {
	gorm *gorm.DB
}

func (c location) Create(input Location) (Location, error) {
	tx := c.gorm.Create(&input)

	return input, tx.Error
}

func (c location) Creates(input []Location) ([]Location, error) {
	tx := c.gorm.Create(&input)

	return input, tx.Error
}

func (c location) FindByFindingID(findingID string) ([]Location, error) {
	var models []Location
	tx := c.gorm.Model(&Location{}).Where(Location{FindingID: findingID}).Find(&models)

	return models, tx.Error
}

func (d db) Location() LocationInterface {
	return location(d)
}
