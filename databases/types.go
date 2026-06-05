package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"database/sql"
)

type ZenqDbContext struct {
	Pool              *sql.DB
	ActiveTransaction *sql.Tx
}
