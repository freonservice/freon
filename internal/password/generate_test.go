package password_test

import (
	"testing"

	"github.com/freonservice/freon/internal/password"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	p := password.New()
	const passLength = 10
	pass := p.Generate(passLength)
	assert.Equal(t, passLength, len(pass))
}
