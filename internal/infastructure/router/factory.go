package router

import (
	"errors"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/validator"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/repository"
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
	validator validator.Validator,
	db repository.SQL,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	case InstanceGin:
		return newGinServer(port, log, validator, db, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
