package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"context"
	"database/sql"

	"github.com/malikhan-dev/zenql/contracts"
	"github.com/malikhan-dev/zenql/streams"
)

func FromSqlRows[T any](ctx context.Context, conn contracts.RDBMSFacade, query string, Mapper func(rows *sql.Rows) (T, error), args ...any) streams.Streamable[T] {
	stream, err := frmSqlRows[T](ctx, conn, query, Mapper, args...)
	return streams.Streamable[T]{
		Context:    ctx,
		Channel:    stream,
		BufferSize: 256,
		Err:        []error{err},
		Initiated:  err == nil,
	}
}
