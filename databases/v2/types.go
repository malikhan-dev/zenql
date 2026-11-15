package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"database/sql"
)

// the db context. holds pointers of the pool and the active database transaction
type ZenqDbContext struct {
	Pool              *sql.DB
	ActiveTransaction *sql.Tx
}
