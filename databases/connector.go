package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/malikhan-dev/zenq/contracts"
)

func Connect(dbName string, constr string) (contracts.RDBMSFacade, error) {

	var db *sql.DB
	var err error
	switch dbName {

	case "mysql":
		db, err = sql.Open(dbName, constr)

	case "postgres":
		db, err = sql.Open(dbName, constr)

	default:
		panic("Unsupported database " + dbName)
	}

	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()

	if pingErr != nil {
		defer db.Close()
		return nil, pingErr
	}

	DbContext := ZenqDB{db}

	return &ZenqDbContext{
		DbContext,
		nil,
	}, nil
}
