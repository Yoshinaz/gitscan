package repo

import (
	"errors"
	gitLib "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func (r *Repo) Clone() error {
	if r.url == "" {
		return errors.New("must be init before use")
	}

	repo, err := gitLib.Clone(memory.NewStorage(), nil, &gitLib.CloneOptions{
		URL: r.url,
	})
	r.repository = repo

	if err != nil {
		return err
	}
	return nil
}
