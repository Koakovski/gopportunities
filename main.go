package main

import (
	"fmt"

	"github.com/Koakovski/gopportunities/config"
	"github.com/Koakovski/gopportunities/router"
)

func main() {
	// Initialize configurations
	if err := config.Init(); err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
