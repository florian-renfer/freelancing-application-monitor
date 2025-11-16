package main

import (
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/log"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/router"
)

func main() {
	// TODO: add database configuration

	var app = infastructure.NewConfig().
		Timeout(10 * time.Second).
		Logger(log.InstanceSlogLogger).
		WebServerPort("4000").
		WebServer(router.InstanceGin)

	app.Start()
}
