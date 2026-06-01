package databases

import (
	"database/sql"
)

type ZenqPgSqlDb struct {
	ZenqDB
}

type ZenqDB struct {
	Pool *sql.DB
}

