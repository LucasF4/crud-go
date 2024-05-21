package main

import (
	"api/config"
	router "api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.ErrF("Config Initialization error: %s", err)
		return
	}

	router.Initializer()
}
