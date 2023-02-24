package git

import (
	gitLib "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func (r *Repo) Clone() error {
	repo, err := gitLib.Clone(memory.NewStorage(), nil, &gitLib.CloneOptions{
		URL: r.URL,
	})
	r.Repo = repo

	if err != nil {
		return err
	}
	return nil
}
