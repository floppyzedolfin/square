package square

import (
	"github.com/floppyzedolfin/square/pkg/logger"
	squaredef "github.com/floppyzedolfin/square/pkg/square"
	"github.com/gofiber/fiber/v2"
)

// Square squares the value contained in the request
func Square(_ fiber.Ctx, req squaredef.Request) (squaredef.Response, *fiber.Error) {
	if req.Value == nil {
		return squaredef.Response{}, fiber.NewError(fiber.StatusBadRequest, "unset request")
	}
	v := *req.Value
	if v == 0 {
		logger.Log(logger.Warning, "You've entered Castle Anthrax!")
		return squaredef.Response{}, fiber.NewError(fiber.StatusNotAcceptable, "naught, naught, naught")
	}

	// implement logic here
	result := v * v

	return squaredef.Response{Value: result}, nil
}
