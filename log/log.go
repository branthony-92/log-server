package log

import "time"

type LogMessage struct{
	Severity  string
	Message   string
	Timestamp time.Time
	Tags      map[string]interface{}
}

