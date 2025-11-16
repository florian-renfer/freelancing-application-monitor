package database

import (
	"os"
	"time"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

func newConfigSqlite() *config {
	return &config{
		host:     os.Getenv("SQLITE_HOST"),
		database: os.Getenv("SQLITE_DATABASE"),
		port:     os.Getenv("SQLITE_PORT"),
		driver:   os.Getenv("SQLITE_DRIVER"),
		user:     os.Getenv("SQLITE_USER"),
		password: os.Getenv("SQLITE_PASSWORD"),
	}
}
