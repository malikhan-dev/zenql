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
	//Close Closes the db connection.
	Close() error

	//Ping Pings the database connection.
	Ping() error

	GetPool() *sql.DB

	// Query Performs a Single Query. query params can be specified. if more than one row can be retrieved it will return an error. it will map the result automatically.
	Query(query string, args ...any) (*sql.Rows, error)

	// Commit Commits the transaction. clears the transaction and returns true if succeeded. if it fails false will be returned.
	Commit() bool

	// Rollback Drops the transaction. changes will be rolled back.
	Rollback() bool

	// Begin Begins the transaction.
	Begin() bool

	GetActiveTransaction() *sql.Tx

	// Exec Executes a Db Command. if an active transaction exists, it will run via that transaction.
	Exec(query string, args ...any) (sql.Result, error)
}

type CommandResult struct {
	Err          error
	RowsAffected int64
	TimeStamp    time.Time
}
