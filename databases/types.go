package databases

import (
	"database/sql"
)

type ZenqDbContext struct {
	ZenqDB
	ActiveTransaction *sql.Tx
}
