package main

import (
	"github.com/Enthreeka/L0/internal/app/server"
	"github.com/Enthreeka/L0/internal/config"
	"github.com/Enthreeka/L0/pkg/logger"
)

func main() {

	configPath := "configs/config.json"

	log := logger.New()

	config, err := config.New(configPath)
	if err != nil {
		log.Fatal("Failed to load config: %s", err)
	}

	log.Fatal("%v", server.Run(log, config))

}
