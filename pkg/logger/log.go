package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
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

const defaultLevel = Debug

// Log prints a message on the standard output. The message has a json format and contains information about the caller
func Log(lvl Level, format string, a ...interface{}) {
	if lvl >= defaultLevel {
		msg := buildLogMessage(lvl, fmt.Sprintf(format, a), 2)
		marshalledLog, _ := json.Marshal(msg)
		fmt.Printf("%v", string(marshalledLog))
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
	Line    int    `json:"line"`
}

func buildLogMessage(lvl Level, msg string, skip int) logMessage {
	// need to skip twice: the buildLogMessage func and the Log func.
	_, file, no, _ := runtime.Caller(skip)
	return logMessage{
		Time:    time.Now().Format(time.RFC3339),
		Level:   levelNames[lvl],
		Message: msg,
		Caller:  file,
		Line:    no,
	}
}
