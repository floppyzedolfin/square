package internal

import (
	"github.com/floppyzedolfin/square/internal/square"
	"github.com/gofiber/fiber/v2"
)

// Server is the microservice application
type Server struct {
	app *fiber.App
}

// NewServer returns a fully operation server, ready to listen
func NewServer() *Server {
	s := Server{app: fiber.New()}
	s.attachEndpoints()
	return &s
}

// Listen starts the server on the associated port
func (s *Server) Listen(port string) {
	s.app.Listen(port)
}

// attachEndpoints adds all necessary endpoints to the server
// add each endpoint here
func (s* Server) attachEndpoints() {
	s.attachSquareEndpoint()
}

// attachSquareEndpoint exposes the Square endpoint on the server
func (s* Server) attachSquareEndpoint() {
	s.app.Post("/square", square.Square)
}