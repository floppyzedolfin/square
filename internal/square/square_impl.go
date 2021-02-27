package square

import (
	squaredef "github.com/floppyzedolfin/square/pkg/square"
	"github.com/gofiber/fiber/v2"
)

// squareImpl squares the value contained in the request
func squareImpl(_ fiber.Ctx, req squaredef.Request) (squaredef.Response, *fiber.Error) {
	if req.Value == 0 {
		return squaredef.Response{}, fiber.NewError(fiber.StatusNotAcceptable, "naught, naught, naught")
	}

	// implement logic here
	result := req.Value * req.Value

	return squaredef.Response{Value: result}, nil
}
