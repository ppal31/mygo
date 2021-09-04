package server

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/ppal31/mygo/internal/config"
)

func load() (*config.AppConfig, error) {
	config := new(config.AppConfig)
	// read the configuration from the environment and
	// populate the configuration structure.
	err := envconfig.Process("", config)
	return config, err
}
