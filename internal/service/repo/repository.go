package repo

import (
	"github.com/gitscan/rules"
	gitLib "gopkg.in/src-d/go-git.v4"
)

type Interface interface {
	Clone() error
	Name() string
	URL() string
	Rules() rules.Interface
	Repo() *gitLib.Repository
	GetHeadCommitHash() (string, error)
}

type Repo struct {
	name       string
	url        string
	repository *gitLib.Repository
	rules      rules.Interface
}

func New(name, url string, rules rules.Interface) Interface {
	return &Repo{
		name:  name,
		url:   url,
		rules: rules,
	}
}

func (r *Repo) Repo() *gitLib.Repository {
	return r.repository
}

func (r *Repo) Name() string {
	return r.name
}

func (r *Repo) URL() string {
	return r.url
}

func (r *Repo) Rules() rules.Interface {
	return r.rules
}

func (r *Repo) GetHeadCommitHash() (string, error) {
	ref, err := r.repository.Head()
	if err != nil {
		return "", err
	}

	return ref.Hash().String(), nil
}
