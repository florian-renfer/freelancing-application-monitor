package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/repository"

	_ "github.com/mattn/go-sqlite3"
)

type (
	sqliteHandler struct {
		db *sql.DB
	}

	sqliteTx struct {
		tx *sql.Tx
	}

	sqliteRow struct {
		row *sql.Row
	}

	sqliteRows struct {
		rows *sql.Rows
	}
)

func NewSqliteHandler(c *config) (*sqliteHandler, error) {
	var ds = fmt.Sprintf(
		"file:%s?cache=shared&mode=memory",
		c.file,
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

func newSqliteTx(tx *sql.Tx) sqliteTx {
	return sqliteTx{
		tx: tx,
	}
}

func newSqliteRow(row *sql.Row) sqliteRow {
	return sqliteRow{
		row: row,
	}
}

func newSqliteRows(rows *sql.Rows) sqliteRows {
	return sqliteRows{
		rows: rows,
	}
}

func (s sqliteRow) Scan(dest ...any) error {
	if err := s.row.Scan(dest...); err != nil {
		return err
	}
	return nil
}

func (s sqliteRows) Scan(dest ...interface{}) error {
	if err := s.rows.Scan(dest...); err != nil {
		return err
	}
	return nil
}

func (s sqliteRows) Next() bool {
	return s.rows.Next()
}

func (s sqliteRows) Err() error {
	return s.rows.Err()
}

func (s sqliteRows) Close() error {
	return s.rows.Close()
}

func (s sqliteHandler) ExecuteContext(ctx context.Context, query string, args ...any) error {
	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s sqliteHandler) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newSqliteRows(rows)

	return row, nil
}

func (s sqliteHandler) QueryRowContext(ctx context.Context, query string, args ...any) repository.Row {
	row := s.db.QueryRowContext(ctx, query, args...)
	return newSqliteRow(row)
}

func (s sqliteHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return sqliteTx{}, err
	}
	return newSqliteTx(tx), nil
}

func (s sqliteTx) ExecuteContext(ctx context.Context, query string, args ...any) error {
	_, err := s.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s sqliteTx) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	rows, err := s.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newSqliteRows(rows)
	return row, nil
}

func (s sqliteTx) QueryRowContext(ctx context.Context, query string, args ...any) repository.Row {
	row := s.tx.QueryRowContext(ctx, query, args...)
	return newSqliteRow(row)
}

func (s sqliteTx) Commit() error {
	return s.tx.Commit()
}

func (s sqliteTx) Rollback() error {
	return s.tx.Rollback()
}
