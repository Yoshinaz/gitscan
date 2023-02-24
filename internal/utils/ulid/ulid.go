package ulid

import (
	"crypto/rand"

	"github.com/oklog/ulid/v2"
)

func New(prefix string) string {
	return prefix + "_" + ulid.MustNew(ulid.Now(), ulid.Monotonic(rand.Reader, 0)).String()
}
