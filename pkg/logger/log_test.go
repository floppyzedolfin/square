package logger

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	tt := map[string]struct {
		defaultLevel  Level
		output        io.ReadWriter
		clock         func() time.Time
		messageLevel  Level
		messageFormat string
		arguments     []interface{}
		loggedMessage string
	}{
		"logging info with debug threshold": {
			defaultLevel:  Debug,
			output:        &bytes.Buffer{},
			clock:         fakeClock,
			messageLevel:  Info,
			messageFormat: "hello %s",
			arguments:     []interface{}{"world!"},
			loggedMessage: "{\"time\":\"1984-05-14T01:23:45.000000067Z\",\"level\":\"Info\",\"message\":\"hello world!\",\"caller\":\"logger.TestLog\"}\n",
		},
		"logging debug with info threshold": {
			defaultLevel:  Info,
			output:        &bytes.Buffer{},
			clock:         fakeClock,
			messageLevel:  Debug,
			messageFormat: "hello %s",
			arguments:     []interface{}{"world!"},
			loggedMessage: "",
		},
	}

	for name, tc := range tt {
		// set the local parameters for this execution
		defaultLevel = tc.defaultLevel
		output = tc.output
		clock = tc.clock

		Log(tc.messageLevel, tc.messageFormat, tc.arguments...)
		res, err := ioutil.ReadAll(tc.output)
		require.NoError(t, err, name)
		assert.Equal(t, tc.loggedMessage, string(res), name)
	}
}

// fakeClock returns an absolutely random date.
func fakeClock() time.Time {
	return time.Date(1984, time.May, 14, 01, 23, 45, 67, time.UTC)
}
