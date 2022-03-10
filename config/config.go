package config

import "github.com/ilyakaznacheev/cleanenv"

// Config presents main system configurations.
// All configurations takes from environment variables.
type Config struct {
	Database Database
	Server   Server
	JWT      JWT
}

// Validate checks the correctness of the configurations.
func (c Config) Validate() error {
	if err := c.Server.Validate(); err != nil {
		return err
	}
	if err := c.Database.Validate(); err != nil {
		return err
	}
	if err := c.JWT.Validate(); err != nil {
		return err
	}

	return nil
}

// LoadConfig returns Config
func LoadConfig() (Config, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
