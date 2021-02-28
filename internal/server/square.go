package server

import (
	"fmt"

	"github.com/floppyzedolfin/square/internal/square"
	"github.com/floppyzedolfin/square/pkg/logger"
	squaredef "github.com/floppyzedolfin/square/pkg/square"
	"github.com/gofiber/fiber/v2"
)

// registerSquare exposes the Square endpoint on the server
func (s* Server) registerSquare() {
	s.app.Post("/square", squareWrapper)
}

// squareWrapper parses the request, calls the intelligent computation function, and returns the result
func squareWrapper(c *fiber.Ctx) error {
	req := new(squaredef.Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("unable to parse body as request: %s", err.Error())})
	}
	logger.Log(logger.Info, "received request for endpoint square: %s", req)

	// Here lies the endpoint's smartness
	res, err := square.Square(*c, *req)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{"error": err.Message})
	}

	return c.JSON(res)
}
