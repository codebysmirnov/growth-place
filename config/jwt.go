package config

import "errors"

// JWT presents JWT configuration.
type JWT struct {
	Secret string `env:"JWT_SECRET"`
}

// Validate checks the correctness of the JWT configurations.
func (j JWT) Validate() error {
	if j.Secret == "" {
		return errors.New("empty JWT_SECRET variable")
	}
	return nil
}
