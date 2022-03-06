package config

import "errors"

// Server presents server settings.
type Server struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`
}

// Validate checks the correctness of the server configurations.
func (s Server) Validate() error {
	if s.Host == "" {
		return errors.New("empty SERVER_HOST variable")
	}
	if s.Port == "" {
		return errors.New("empty SERVER_PORT variable")
	}
	return nil
}
