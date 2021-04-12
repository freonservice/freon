package password

import (
	"math/rand"
)

// Generate using for generation password
func (p *Password) Generate(length int) string {
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" + digits + specials
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]     //nolint:gosec
	buf[1] = specials[rand.Intn(len(specials))] //nolint:gosec
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))] //nolint:gosec
	}
	for i := len(buf) - 1; i > 0; i-- {
		j := rand.Intn(i + 1) //nolint:gosec
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
