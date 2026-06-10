package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"context"

	"github.com/malikhan-dev/zenql/contracts"
	"github.com/malikhan-dev/zenql/streams"
)

func FromSqlRows[T any](ctx context.Context, conn contracts.RDBMSFacade, query string, args ...any) streams.Streamable[T] {
	stream, err := frmSqlRows[T](ctx, conn, query, args...)
	return streams.Streamable[T]{
		Context:    ctx,
		Channel:    stream,
		BufferSize: 256,
		Err:        []error{err},
		Initiated:  err == nil,
	}
}
