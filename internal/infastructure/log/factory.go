package log

import (
	"errors"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
)

const (
	InstanceSlogLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceSlogLogger:
		return NewSlogLogger()
	default:
		return nil, errInvalidLoggerInstance
	}
}
