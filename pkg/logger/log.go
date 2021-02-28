package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type Level int

const (
	Debug Level = iota
	Info
	Warning
	Error
	noLog // set the defaultLevel to noLog to discard all logs
)

var (
	// default parameters - they aren't exposed and can't be changed by the rest of the world, but, in order to properly test this, I can't have hardwritten consts
	defaultLevel                  = Debug     // change this to change the logging threshold
	output       io.Writer        = os.Stdout // change this to change the default output - check the tests
	clock        func() time.Time = time.Now  // used to get time
)

// Log prints a message on the standard output. The message has a json format and contains information about the caller
func Log(lvl Level, format string, a ...interface{}) {
	if lvl >= defaultLevel {
		msg := buildLogMessage(lvl, fmt.Sprintf(format, a...), clock, 2)
		marshalledLog, _ := json.Marshal(msg)
		fmt.Fprintf(output, "%v\n", string(marshalledLog))
	}
}

var levelNames = map[Level]string{
	Debug:   "Debug",
	Info:    "Info",
	Warning: "Warning",
	Error:   "Error",
}

type logMessage struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Caller  string `json:"caller"`
}

func buildLogMessage(lvl Level, msg string, t func() time.Time, skip int) logMessage {
	// get info about the caller
	callerFunc := ""
	fpcs := make([]uintptr, 1)
	// need to go back one step in the stack (to the level of this func's caller + whatever that skips)
	n := runtime.Callers(skip+1, fpcs)
	if n != 0 {
		caller := runtime.FuncForPC(fpcs[0] - 1)
		if caller != nil {
			callerFunc = caller.Name()
			path := strings.Split(callerFunc, string(os.PathSeparator))
			callerFunc = path[len(path)-1]
		}
	}

	return logMessage{
		Time:    t().Format(time.RFC3339Nano),
		Level:   levelNames[lvl],
		Message: msg,
		Caller:  callerFunc,
	}
}
