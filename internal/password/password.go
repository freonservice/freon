package password

import (
	"github.com/MarcSky/freon/internal/app"

	"golang.org/x/crypto/bcrypt"
)

type (
	Password struct {
		cost int
	}
	Option func(*Password)
)

func New(options ...Option) app.Password {
	p := &Password{cost: bcrypt.DefaultCost}

	for i := range options {
		options[i](p)
	}

	return p
}

func (p *Password) Hashing(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), p.cost)
}

func (p *Password) Compare(hashedPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
