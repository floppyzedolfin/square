package pkg

import (
	"github.com/jinzhu/gorm"
)

type SquareRequest struct {
	gorm.Model // needed to parse this body
	Value int `json:"value"`
}

type SquareResponse struct {
	Value int `json:"value"`
}
