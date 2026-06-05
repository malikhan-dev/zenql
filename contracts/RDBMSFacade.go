package contracts

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"database/sql"
	"time"
)

type RDBMSFacade interface {
	Close() error
	Ping() error
	GetPool() *sql.DB
	Query(query string, args ...any) (*sql.Rows, error)
	Commit() bool
	Rollback() bool
	Begin() bool
	GetActiveTransaction() *sql.Tx
	Exec(query string, args ...any) (sql.Result, error)
}

type Commandesult struct {
	Err          error
	RowsAffected int64
	TimeStamp    time.Time
}
