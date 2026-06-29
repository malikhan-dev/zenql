package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/malikhan-dev/zenql/contracts/v2"
)

func (conn *ZenqDbContext) Ping() error {
	return conn.Pool.Ping()
}

func (conn *ZenqDbContext) Close() error {
	return conn.Pool.Close()
}
func (conn *ZenqDbContext) GetPool() *sql.DB {
	return conn.Pool
}

func (conn *ZenqDbContext) Begin() bool {

	tr, err := conn.GetPool().Begin()

	if err == nil {
		conn.ActiveTransaction = tr
		return true
	}
	return false

}

func (conn *ZenqDbContext) GetActiveTransaction() *sql.Tx {
	return conn.ActiveTransaction
}
func (conn *ZenqDbContext) Commit() bool {

	if conn.ActiveTransaction != nil {
		err := conn.ActiveTransaction.Commit()
		if err == nil {
			conn.ActiveTransaction = nil
			return true
		}
		return false
	}
	return false
}
func (conn *ZenqDbContext) Rollback() bool {
	if conn.ActiveTransaction != nil {
		err := conn.ActiveTransaction.Rollback()

		if err == nil {
			conn.ActiveTransaction = nil
			return true
		}
		return false
	}
	return false
}

func (conn *ZenqDbContext) Query(query string, args ...any) (*sql.Rows, error) {

	return conn.Pool.Query(query, args...)
}

func (conn *ZenqDbContext) Exec(query string, args ...any) (sql.Result, error) {

	return conn.Pool.Exec(query, args...)
}

func Query[T any](conn contracts.RDBMSFacade, query string, args ...any) ([]T, error) {

	rows, err := conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return contracts.MapRows[T](rows, false)
}

func Exec(conn contracts.RDBMSFacade, cmd string, args ...any) contracts.CommandResult {

	var err error
	var result sql.Result

	if conn.GetActiveTransaction() != nil {
		result, err = conn.GetActiveTransaction().Exec(cmd, args...)
	} else {
		result, err = conn.Exec(cmd, args...)

	}

	var affected int64
	var error error

	if result == nil {
		affected = 0
		error = err
	} else {
		if aff, er := result.RowsAffected(); err == nil {
			affected = aff
			error = er
		} else {
			affected = 0
			error = err
		}
	}

	return contracts.CommandResult{
		Err:          error,
		RowsAffected: affected,
		TimeStamp:    time.Now(),
	}

}

func SingleQuery[T any](conn contracts.RDBMSFacade, query string, args ...any) ([]T, error) {

	rows, err := conn.Query(query, args...)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	return contracts.MapRows[T](rows, true)
}
