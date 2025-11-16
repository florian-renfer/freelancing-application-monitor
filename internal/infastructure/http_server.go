package infastructure

import (
	"os"
	"strconv"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/log"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/router"
)

type config struct {
	logger        logger.Logger
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Logger(instance int) *config {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}

	c.logger = log
	c.logger.Infof("Successfully configured log")

	return c
}

func (c *config) Timeout(ctxTimeout time.Duration) *config {
	c.ctxTimeout = ctxTimeout
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Errorf("%s", err)
		os.Exit(1)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.logger,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Errorf("%s", err)
		os.Exit(1)
	}

	c.logger.Infof("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}

// TODO: add database configuration
