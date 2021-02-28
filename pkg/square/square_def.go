package square

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Request holds the contents of the request, along with its json fields' names
type Request struct {
	gorm.Model // needed to parse this body
	Value *int `json:"value` // making it a pointer makes it mandatory, if we implement the checks properly
}

func (r *Request) Validate() error {
	if r.Value == nil {
		return fmt.Errorf("invalid request, field Value musn't be empty")
	}
	return nil
}

// String implements the stringer interface, and hides the gorm thingy
func (r *Request) String() string {
	return fmt.Sprintf("Value: %d", r.Value)
}

// Response holds the contents of the response, along with its json fields's names
type Response struct {
	Value int `json:"value"`
}
