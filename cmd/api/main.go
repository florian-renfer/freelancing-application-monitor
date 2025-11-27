package main

import (
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/database"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/log"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/router"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/validation"
)

func main() {
	var app = infastructure.NewConfig().
		Timeout(10 * time.Second).
		Logger(log.InstanceSlogLogger).
		Persistence(database.InstanceSqlite).
		Validator(validation.InstanceGoPlayground).
		WebServerPort("4000").
		WebServer(router.InstanceGin)

	app.Start()
}
