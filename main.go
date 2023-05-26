package main

import (
	"github.com/Koakovski/gopportunities/config"
	"github.com/Koakovski/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger()

	// Initialize configurations
	if err := config.Init(); err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
