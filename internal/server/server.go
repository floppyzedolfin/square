package server

import (
	"github.com/gofiber/fiber/v2"
)

// Server is the microservice application
type Server struct {
	app *fiber.App
}

// NewServer returns a fully operational server, ready to Listen()
func NewServer() *Server {
	s := Server{app: fiber.New()}
	s.registerEndpoints()
	return &s
}

// Listen starts the server on the associated port
func (s *Server) Listen(port string) {
	s.app.Listen(port)
}

// registerEndpoints adds all necessary endpoints to the server
// add each endpoint here
func (s* Server) registerEndpoints() {
	s.registerSquare()
}
