package streams

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"context"
	"sync"
	"time"

	"github.com/malikhan-dev/zenql/contracts/v2"
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

	stream := fromData[T](ctx, buffer_size, items)

	return Streamable[T]{
		Context:    ctx,
		Channel:    stream,
		BufferSize: 256,
		Err:        nil,
		Initiated:  true,
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

func (currStr Streamable[T]) StopIf(StopWhen func(item T) bool, cancelFunc context.CancelFunc) Streamable[T] {

	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    stopIf(currStr.Context, cancelFunc, currStr.Channel, StopWhen),
		BufferSize: currStr.BufferSize,
	}

}

func (currStr Streamable[T]) CallIf(CallWhen func(item T) bool, callback func(item T)) Streamable[T] {

	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    callIf(currStr.Context, currStr.Channel, CallWhen, callback),
		BufferSize: currStr.BufferSize,
	}

}

func FromSqlRows[T any](ctx context.Context, conn contracts.RDBMSFacade, query string, args ...any) Streamable[T] {
	stream, err := frmSqlRows[T](ctx, conn, query, args...)
	return Streamable[T]{
		Context:    ctx,
		Channel:    stream,
		BufferSize: 256,
		Err:        []error{err},
		Initiated:  err == nil,
	}
}

func (currStr Streamable[T]) Process() {
	for range currStr.Channel {
	}
}

func (currStr Streamable[T]) BackgroundProcess(waitGroup *sync.WaitGroup) {

	go func() {
		defer waitGroup.Done()

		for range currStr.Channel {

		}

	}()

}

func (currStr Streamable[T]) Pipe() <-chan T {
	return currStr.Channel
}
