package main

import (
	"app/internal/logger"
	"app/internal/server"
	"app/internal/tools"
	"os"
)

// main entrypoint
func main() {
	port := tools.EnvPortOr("3000")

	logger.Stdout.Info("starting server on port " + port[1:])

	// start listening on port
	if err := server.StartServer(port); err != nil {
		logger.Stderr.Error(err.Error())
		os.Exit(1)
	}
}
