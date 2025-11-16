package router

import (
	"errors"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid web server instance")
)

const (
	InstanceGin int = iota
)

func NewWebServerFactory(
	instance int,
	log logger.Logger,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	case InstanceGin:
		return newGinServer(port, log, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
