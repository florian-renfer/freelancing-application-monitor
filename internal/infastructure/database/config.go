package database

import (
	"os"
)

type config struct {
	file   string
	driver string
}

func newConfigSqlite() *config {
	return &config{
		file:   os.Getenv("SQLITE_FILE"),
		driver: os.Getenv("SQLITE_DRIVER"),
	}
}
