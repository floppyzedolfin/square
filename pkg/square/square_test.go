package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestValidate(t *testing.T) {
	anInt := 2
	tt := map[string]struct {
		req    Request
		errMsg string
	}{
		"value = 2": {
			req: Request{Value: &anInt},
		},
		"unset value": {
			req:    Request{},
			errMsg: "invalid request, field Value musn't be empty",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := tc.req.Validate()
			if tc.errMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRequestString(t *testing.T) {
	anInt := 3
	tt := map[string]struct {
		req    Request
		resp string
	}{
		"value = 2": {
			req: Request{Value: &anInt},
			resp: "Value: 3",
		},
		"unset value": {
			req:    Request{},
			resp: "Value: nil",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			str := tc.req.String()
			assert.Equal(t, tc.resp, str)
		})
	}
}
