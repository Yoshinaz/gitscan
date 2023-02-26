package ulid_test

import (
	"github.com/gitscan/internal/utils/ulid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("return a uuid with prefix", func(t *testing.T) {
		result := ulid.New("prefix")
		assert.True(t, strings.HasPrefix(result, "prefix"))
	})
}
