package repo

import "fmt"

type Config struct {
	Host          string
	Port          int
	Name          string
	User          string
	Pass          string
	MaxIdleConns  int
	MaxOpenConns  int
	MigrationPath string
}

func (c *Config) FormatDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Pass, c.Name,
	)
}
