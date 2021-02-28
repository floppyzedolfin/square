package main

import (
	"github.com/floppyzedolfin/square/internal/server"
)

func main() {
	// Create the server that will have all the necessary endpoints
	s := server.NewServer()

	s.Listen(":3000")
}
