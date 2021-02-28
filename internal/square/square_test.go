package square

import (
	"testing"

	"github.com/floppyzedolfin/square/pkg/square"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSquare(t *testing.T) {
	tt := map[string]struct {
		inputValue   *int
		returnCode   int
		errMsg       string
		squaredValue int
	}{
		"2 squared": {
			inputValue:   intPtr(2),
			squaredValue: 4,
		},
		"0 squared - error case": {
			inputValue:   intPtr(0),
			returnCode:   fiber.StatusNotAcceptable,
			errMsg: "naught, naught, naught",
		},
		"invalid request - error case": {
			inputValue:   nil,
			returnCode:   fiber.StatusBadRequest,
			errMsg: "unset request",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			req := square.Request{Value: tc.inputValue}
			res, err := Square(fiber.Ctx{}, req)
			if tc.errMsg != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.returnCode, err.Code)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.Nil(t, err)
				assert.Equal(t, tc.squaredValue, res.Value)
			}
		})
	}
}

func intPtr(v int) *int {
	return &v
}