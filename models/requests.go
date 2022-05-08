package models

import (
	"errors"
	"time"
)

type LogRetrievalRequestDate struct {
	Source     string
	RangeStart time.Time
	RangeEnd   time.Time
}

type LogRetrievalRequestID struct {
	Source string
	IDs    []string
}

type LogUploadRequest struct {
	Source   string
	LogLevel string
	Message  LogMessage
}

func (m LogUploadRequest) Validate() error {
	if m.LogLevel == "" {
		return errors.New("message invalid, log level empty")
	}
	if m.Source == "" {
		return errors.New("message invalid, log source empty")
	}
	if m.Message.Severity == "" {
		return errors.New("message invalid, log severity empty")
	}
	if m.Message.Text == "" {
		return errors.New("message invalid, log message empty")
	}

	if _, ok := logLevel[m.LogLevel]; !ok {
		return errors.New("message invalid, invalid log level")
	}
	if _, ok := logLevel[m.Message.Severity]; !ok {
		return errors.New("message invalid, invalid severity")
	}

	return nil
}

func (m LogUploadRequest) StoreLog() bool {
	return logLevel[m.LogLevel] >= logLevel[m.Message.Severity]
}
