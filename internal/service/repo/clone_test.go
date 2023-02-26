package repo_test

import (
	"github.com/gitscan/internal/service/repo"
	"github.com/gitscan/rules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_Clone(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rule := rules.New("")
		r := repo.New("test", "https://github.com/Yoshinaz/test_secret", rule)
		err := r.Clone()

		assert.Nil(t, err)
	})
}
