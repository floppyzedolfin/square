package square

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func Square(c *fiber.Ctx) {
	req := new(SquareRequest)
	if err := c.BodyParser(req); err != nil {
		c.Status(503).Send(err)
		return
	}
	fmt.Printf("%v\n",req)
	resp := SquareResponse{Value: req.Value*req.Value}
	c.JSON(resp)
}

type SquareRequest struct {
	gorm.Model
	Value int `json:"value"`
}

type SquareResponse struct {
	Value int `json:"value"`
}