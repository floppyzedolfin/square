package main

import (
	"github.com/floppyzedolfin/square/internal"
)

func main() {
	s := internal.NewServer()

	s.Listen(3000)
}