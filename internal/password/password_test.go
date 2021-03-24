package password_test

import (
	"testing"

	"github.com/MarcSky/freon/internal/password"

	"github.com/stretchr/testify/assert"
)

var (
	pass = "pass"
)

func TestPassword(t *testing.T) {
	t.Parallel()

	passwords := password.New()
	hashPass, err := passwords.Hashing(pass)
	assert.NoError(t, err)
	compare := passwords.Compare(hashPass, []byte(pass))
	assert.Equal(t, true, compare)
}
