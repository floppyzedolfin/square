package main

import (
	"github.com/floppyzedolfin/square/internal"
)

func main() {
	// Create the server that will have all the necessary endpoints
	s := internal.NewServer()

	s.Listen(":3000")
}