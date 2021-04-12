package password_test

import (
	"testing"

	"github.com/freonservice/freon/internal/password"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	var (
		pass = "pass"
	)

	passwords := password.New()
	hashPass, err := passwords.Hashing(pass)
	assert.NoError(t, err)
	compare := passwords.Compare(hashPass, []byte(pass))
	assert.Equal(t, true, compare)
}
