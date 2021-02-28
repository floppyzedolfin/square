package main

import (
	"github.com/floppyzedolfin/square/internal/server"
	"github.com/floppyzedolfin/square/pkg/logger"
)

func main() {
	// Create the server that will have all the necessary endpoints
	s := server.NewServer()

	const port = ":3000"
	logger.Log(logger.Info, "starting server on port %s", port)
	s.Listen(port)
}
