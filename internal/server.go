package internal

import (
	"github.com/floppyzedolfin/square/internal/square"
	"github.com/gofiber/fiber"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	s := Server{app: fiber.New()}
	s.createEndpoints()
	return &s
}

func (s *Server) Listen(port int) {
	s.app.Listen(port)
}

func (s* Server) createEndpoints() {
	s.createSquareEndpoint()
}

func (s* Server) createSquareEndpoint() {
	s.app.Post("/square", square.Square)
}