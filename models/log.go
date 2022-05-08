package models

import (
	"time"
)

type LogSeverity string

const (
	Info    = "INFO"
	Warning = "WARN"
	Error   = "ERROR"
	Fatal   = "FATAL"
	Debug   = "DEBUG"
)

var logLevel map[string]int

func init() {
	logLevel = make(map[string]int, 0)

	logLevel[Info] = 0
	logLevel[Warning] = 1
	logLevel[Error] = 2
	logLevel[Fatal] = 3
	logLevel[Debug] = 4
}

type LogMessage struct {
	Severity  string
	Text      string
	Timestamp time.Time
	Tags      map[string]any
}
