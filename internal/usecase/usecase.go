package usecase

import (
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/report"
	"github.com/gitscan/internal/service/repo"
	"github.com/gitscan/rules"
)

type Interface interface {
	Scan(name, url, ruleID string, scanAll bool, working chan bool) (string, error)
	View(name, url, ruleID string) report.Info
	Recovery(name, url, ruleID string, working chan bool)
}

type UseCase struct {
	DB database.DB
}

func New(db database.DB) UseCase {
	return UseCase{DB: db}
}

func (u UseCase) Scan(name, url, ruleSet string, scanAll bool, working chan bool) (string, error) {
	rule := rules.New(ruleSet)
	r := repo.New(name, url, rule)

	return repo.Scan(r, scanAll, u.DB, working)
}

func (u UseCase) View(name, url, ruleSet string) (report.Info, error) {
	rule := rules.New(ruleSet)
	r := repo.New(name, url, rule)

	return repo.ViewReport(r, u.DB)
}

func (u UseCase) Recovery(infoDB database.Info, working chan bool) {
	rule := rules.New(infoDB.RuleSet)
	r := repo.New(infoDB.Name, infoDB.URL, rule)
	r.Clone()
	repo.ProcessScan(r, infoDB, u.DB, true, working)
}
