package main

import (
	"fmt"
	"os"
	"restService/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Development: %s, Port: %s", cfg.Development, cfg.Port)
}
