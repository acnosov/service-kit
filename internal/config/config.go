package config

import (
	"github.com/kelseyhightower/envconfig"
	"os"
)

func init() {
	if os.Getenv("ENV") == "dev" {
		MustLoadEnv()
	}
}

// Load returns the settings from the environment.
func Load() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	return cfg, err
}
