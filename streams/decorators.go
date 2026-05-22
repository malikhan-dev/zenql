package streams

import (
	"context"
	"time"

	"github.com/malikhan-dev/zenq/contracts"
)

const buffer_size = 256

func FromCsv[T any](ctx context.Context, Conf contracts.CsvStreamConf[T]) Streamable[T] {
	return Streamable[T]{
		Context: ctx,
		Channel: fromCsv[T](ctx, Conf),
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
		Context: currStr.Context,
		Channel: filterStream[T](currStr.Context, buffer_size, currStr.Channel, Filter),
	}
}

func (currStr Streamable[T]) TakeAll() []T {
	return takeAll[T](currStr.Context, currStr.Channel)
}

func (currStr Streamable[T]) Throttle(duration time.Duration) Streamable[T] {
	return Streamable[T]{
		Context: currStr.Context,
		Channel: throttle[T](currStr.Context, currStr.Channel, duration),
	}
}
