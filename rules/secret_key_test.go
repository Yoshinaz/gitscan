package rules_test

import (
	"github.com/gitscan/rules"
	"github.com/stretchr/testify/assert"
	gitLib "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"testing"
)

func TestSecretKey_GetRuleInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rule := rules.NewSecretKey()
		info := rule.GetRuleInfo()

		assert.Equal(t, "G401", info.RuleID)
	})
}

func TestSecretKey_Process(t *testing.T) {
	repository, err := gitLib.Clone(memory.NewStorage(), nil, &gitLib.CloneOptions{
		URL: "https://github.com/Yoshinaz/test_secret",
	})
	assert.Nil(t, err)
	t.Run("success", func(t *testing.T) {
		rule := rules.NewSecretKey()
		cIter, _ := repository.Log(&gitLib.LogOptions{})
		err = cIter.ForEach(func(c *object.Commit) error {
			fIter, err := c.Files()
			if err != nil {
				return err
			}
			err = fIter.ForEach(func(file *object.File) error {
				_, _, err := rule.Process(file)

				assert.Nil(t, err)

				return nil
			})
			
			return nil
		})

	})
}
