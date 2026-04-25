package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Development string
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT is empty")
	}

	dev := os.Getenv("DEVELOPMENT")
	if dev == "" {
		dev = "local"
	}

	conf := &Config{
		Port:        port,
		Development: dev,
	}

	return conf, nil
}
