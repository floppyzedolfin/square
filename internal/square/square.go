package square

import (
	"fmt"

	"github.com/floppyzedolfin/square/pkg/logger"
	squaredef "github.com/floppyzedolfin/square/pkg/square"
	"github.com/gofiber/fiber/v2"
)

// Square parses the request, performs the computation, and returns the result
func Square(c *fiber.Ctx) error {
	req := new(squaredef.Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("unable to parse body as request: %s", err.Error()))
	}
	logger.Log(logger.Info, "received request for endpoint square: %s", req)

	res, err := squareImpl(*c, *req)
	if err != nil{
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.JSON(res)
}
