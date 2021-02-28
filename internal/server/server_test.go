package server

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Server(t *testing.T) {
	tt := map[string]struct {
		verb        string
		path        string
		body        string
		contentType string
		errMsg      string
		returnCode  int
		response    string
	}{
		"nominal, on existing endpoint": {
			verb:        http.MethodPost,
			path:        "/square",
			body:        "{\"value\":2}",
			contentType: "application/json",
			returnCode:  fiber.StatusOK,
			response:    "{\"value\":4}",
		},
		"endpoint is ok, but wrong body": {
			verb:        http.MethodPost,
			path:        "/square",
			body:        "{\"value\":\"a\"}",
			contentType: "application/json",
			returnCode:  fiber.StatusBadRequest,
			errMsg:      "{\"error\":\"unable to parse body as request", // and stuff that are not part of this code
		},
		"error in endpoint": {
			verb:        http.MethodPost,
			path:        "/square",
			body:        "{\"value\":0}",
			contentType: "application/json",
			returnCode:  fiber.StatusNotAcceptable, // I'm not sure I should be covering this - it's hidden within another package
			errMsg:      "{\"error\":\"",           // and stuff that are not part of this code
		},
		"wrong verb": {
			verb:       http.MethodGet,
			path:       "/square",
			returnCode: fiber.StatusMethodNotAllowed,
			errMsg:     "Method Not Allowed",
		},
		"non-existing endpoint": {
			verb:       http.MethodGet,
			path:       "/cube",
			returnCode: fiber.StatusNotFound,
			errMsg:     "Cannot GET /cube",
		},
	}

	s := NewServer()

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(tc.verb, tc.path, strings.NewReader(tc.body))
			require.NoError(t, err)
			req.Header.Add("content-type", tc.contentType)

			res, err := s.app.Test(req)
			assert.Equal(t, tc.returnCode, res.StatusCode)
			assert.NoError(t, err) // errors are handled inside the endpoint
			if tc.errMsg != "" {
				body, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				// we don't want to rely on what the inner libs do - a mere check that the error contains what is written in the endpoint is enough
				assert.Contains(t, string(body), tc.errMsg)
			} else {
				body, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				assert.Equal(t, tc.response, string(body))
			}
		})
	}
}
