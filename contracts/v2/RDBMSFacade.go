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
	// Closes the db connection.
	Close() error

	// Pings the database connection.
	Ping() error

	GetPool() *sql.DB

	// Performs a Single Query. query params can be specified. if more than one row can be retrieved it will return an error. it will map the result automatically.
	Query(query string, args ...any) (*sql.Rows, error)

	// Commits the transaction. clears the transaction and returns true if succeeded. if it fails false will be returned.
	Commit() bool

	// Drops the transaction. changes will be rolled back.
	Rollback() bool

	// Begins the transaction.
	Begin() bool

	GetActiveTransaction() *sql.Tx

	// Executes a Db Command. if an active transaction exists, it will run via that transaction.
	Exec(query string, args ...any) (sql.Result, error)
}

type CommandResult struct {
	Err          error
	RowsAffected int64
	TimeStamp    time.Time
}
