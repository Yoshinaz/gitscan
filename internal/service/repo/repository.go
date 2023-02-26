package repo

import (
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/report"
	"github.com/gitscan/rules"
	gitLib "gopkg.in/src-d/go-git.v4"
)

type Interface interface {
	Init(name, url string, rules rules.Interface)
	Clone() error
	Scan(db database.DB, working chan bool) (string, error)
	ViewReport(db database.DB) (report.Info, error)
}

type Repo struct {
	Name  string
	URL   string
	Repo  *gitLib.Repository
	Rules rules.Interface
}

func New() Interface {
	return &Repo{
		Name: "",
		URL:  "",
		Repo: nil,
	}
}
