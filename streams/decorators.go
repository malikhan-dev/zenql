package streams

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"context"
	"time"

	"github.com/malikhan-dev/zenql/contracts"
)

const buffer_size = 256

func FromCsv[T any](ctx context.Context, Conf contracts.CsvStreamConf[T]) Streamable[T] {
	stream, err := fromCsv[T](ctx, Conf)
	return Streamable[T]{
		Context:    ctx,
		Channel:    stream,
		BufferSize: Conf.BufferSize,
		Err:        []error{err},
		Initiated:  err == nil,
	}
}

func FromJsonArr[T any](ctx context.Context, Conf contracts.StreamConf) Streamable[T] {
	stream, err := fromJsonArr[T](ctx, Conf)
	return Streamable[T]{
		Context:    ctx,
		Channel:    stream,
		BufferSize: Conf.BufferSize,
		Err:        []error{err},
		Initiated:  err == nil,
	}
}

func FromData[T any](ctx context.Context, items []T) Streamable[T] {

	return Streamable[T]{
		Context: ctx,
		Channel: fromData[T](ctx, buffer_size, items),
	}
}

func FromChannel[T any](ctx context.Context, items <-chan T) Streamable[T] {

	return Streamable[T]{
		Context: ctx,
		Channel: fromChannel[T](ctx, buffer_size, items),
	}
}

func (currStr Streamable[T]) FilterStream(Filter func(T) bool) Streamable[T] {
	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    filterStream[T](currStr.Context, currStr.BufferSize, currStr.Channel, Filter),
		BufferSize: currStr.BufferSize,
		Err:        currStr.Err,
	}
}

func (currStr Streamable[T]) TakeAll() []T {
	return takeAll[T](currStr.Context, currStr.Channel)
}

func (currStr Streamable[T]) Throttle(duration time.Duration) Streamable[T] {
	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    throttle[T](currStr.Context, currStr.Channel, duration),
		BufferSize: currStr.BufferSize,
	}
}
