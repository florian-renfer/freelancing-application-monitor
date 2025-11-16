package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func NewSqliteHandler(c *config) (*sqliteHandler, error) {
	// TODO: use correct DSN provided via github.com/mattn/go-sqlite3
	var ds = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.database,
		c.password,
	)

	fmt.Println(ds)
	db, err := sql.Open(c.driver, ds)
	if err != nil {
		return &sqliteHandler{}, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &sqliteHandler{db: db}, nil
}

// TODO: Implement adapter.SQL and adapter.Tx interfaces
