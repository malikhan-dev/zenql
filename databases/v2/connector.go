package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/malikhan-dev/zenql/contracts/v2"
)

var UnsupportedDb = errors.New("Unsupported DB")

// To establish a connection with a database. mysql and postgres supported at the moment.
func Connect(dbName string, constr string) (contracts.RDBMSFacade, error) {

	var db *sql.DB
	var err error
	switch dbName {

	case "mysql":
		db, err = sql.Open(dbName, constr)

	case "postgres":
		db, err = sql.Open(dbName, constr)

	default:
		err = UnsupportedDb
	}

	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()

	if pingErr != nil {
		defer db.Close()
		return nil, pingErr
	}

	return &ZenqDbContext{
		db,
		nil,
	}, nil
}
