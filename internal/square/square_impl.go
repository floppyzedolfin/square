package square

import (
	"github.com/floppyzedolfin/square/pkg"
	"github.com/gofiber/fiber/v2"
)

func squareImpl(_ fiber.Ctx, req pkg.SquareRequest) (pkg.SquareResponse, *fiber.Error) {
	if req.Value == 0 {
		return pkg.SquareResponse{}, fiber.NewError(fiber.StatusNotAcceptable, "naught, naught, naught")
	}

	// implement logic here
	result := req.Value * req.Value

	return pkg.SquareResponse{Value: result}, nil
}
