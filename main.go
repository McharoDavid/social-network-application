package main

import (
	"github/social-network/handlers"
	"os"

	"github.com/google/logger"
)

func main() {

	// create a file to store logs
	logFile, err := os.Create("logs.txt")
	if err != nil {
		panic(err)
	}
	// Initialize logger
	logger.Init("social-network", false, true, logFile)
	// Initialize the routes
	handlers.Init()

}

// fg - is the commnand used to bring back the jobs or servers running in the background
