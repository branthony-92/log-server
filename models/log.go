package models

import "time"

type LogMessage struct {
	Severity  string
	Message   string
	Timestamp time.Time
	Tags      map[string]any
	Source    string
}
