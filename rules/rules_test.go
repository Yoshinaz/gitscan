package rules_test

import (
	"github.com/gitscan/rules"
	"github.com/stretchr/testify/assert"
	gitLib "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"testing"
)

func TestRules_Process(t *testing.T) {
	repository, err := gitLib.Clone(memory.NewStorage(), nil, &gitLib.CloneOptions{
		URL: "https://github.com/Yoshinaz/test_secret",
	})
	assert.Nil(t, err)
	t.Run("processing rules", func(t *testing.T) {
		ref, err := repository.Head()
		assert.Nil(t, err)
		rule := rules.New("1234")
		cIter, err := repository.Log(&gitLib.LogOptions{From: ref.Hash()})
		err = cIter.ForEach(func(c *object.Commit) error {
			finding, err := rule.Process(c)
			assert.Nil(t, err)
			assert.NotEmpty(t, finding)

			return nil
		})
	})
}
