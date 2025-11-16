package database

import (
	"errors"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/repository"
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

const (
	InstanceSqlite int = iota
)

func NewDatabaseSQLFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstanceSqlite:
		return NewSqliteHandler(newConfigSqlite())
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}
