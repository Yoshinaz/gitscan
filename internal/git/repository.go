package git

import (
	gitLib "gopkg.in/src-d/go-git.v4"
	"guardrail/gitscan/internal/database"
	"guardrail/gitscan/internal/report"
)

type Interface interface {
	Clone() error
	Scan(db database.DB, working chan bool) (string, error)
	ViewReport(db database.DB) (report.Info, error)
}

type Repo struct {
	Name string
	URL  string
	Repo *gitLib.Repository
}

func New(name, url string) Interface {
	return &Repo{
		Name: name,
		URL:  url,
		Repo: nil,
	}
}
