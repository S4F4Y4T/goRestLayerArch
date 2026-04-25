package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Development string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
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
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
	}

	return conf, nil
}
