package main

import (
	"github.com/Koakovski/gopportunities/config"
	"github.com/Koakovski/gopportunities/router"
)

var (
	logger *config.Logger = config.GetLogger()
)

func main() {
	// Initialize configurations
	if err := config.Init(); err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
