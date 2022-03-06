package config

import "errors"

// Database presents db connection.
type Database struct {
	Dialect  string `env:"DB_DIALECT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	SSLMode  string `env:"DB_SSL_MODE"`
}

// Validate checks the correctness of the db configurations.
func (db Database) Validate() error {
	if db.Dialect == "" {
		return errors.New("empty DB_DIALECT variable")
	}
	if db.User == "" {
		return errors.New("empty DB_USER variable")
	}
	if db.Password == "" {
		return errors.New("empty DB_PASSWORD variable")
	}
	if db.Name == "" {
		return errors.New("empty DB_NAME variable")
	}
	if db.Host == "" {
		return errors.New("empty DB_HOST variable")
	}
	if db.Port == 0 {
		return errors.New("empty DB_PORT variable")
	}
	if db.SSLMode == "" {
		return errors.New("empty DB_SSL_MODE variable")
	}
	return nil
}
