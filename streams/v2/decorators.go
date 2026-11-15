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

// Initiates data stream from a csv file. needs configuration with an instance of CsvStreamConf[T]
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

// Initiates data stream from a JSON file. needs configuration with an instance of StreamConf[T]
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

// Initiates data stream from a generic slice.
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

// Initiates data stream from a generic channel.
func FromChannel[T any](ctx context.Context, items <-chan T) Streamable[T] {

	return Streamable[T]{
		Context: ctx,
		Channel: fromChannel[T](ctx, buffer_size, items),
	}
}

// Filters the streams of data with a specific criteria.
func (currStr Streamable[T]) FilterStream(Filter func(T) bool) Streamable[T] {
	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    filterStream[T](currStr.Context, currStr.BufferSize, currStr.Channel, Filter),
		BufferSize: currStr.BufferSize,
		Err:        currStr.Err,
	}
}

// Collects All the stream result.
func (currStr Streamable[T]) TakeAll() []T {
	return takeAll[T](currStr.Context, currStr.Channel)
}

// Creates a pause while streaming each item.
func (currStr Streamable[T]) Throttle(duration time.Duration) Streamable[T] {
	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    throttle[T](currStr.Context, currStr.Channel, duration),
		BufferSize: currStr.BufferSize,
	}
}

// Stops the stream of data on the given criteria.
func (currStr Streamable[T]) StopIf(StopWhen func(item T) bool, cancelFunc context.CancelFunc) Streamable[T] {

	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    stopIf(currStr.Context, cancelFunc, currStr.Channel, StopWhen),
		BufferSize: currStr.BufferSize,
	}

}

// if a given criteria is met during the stream a function call will be performed.
func (currStr Streamable[T]) CallIf(CallWhen func(item T) bool, callback func(item T)) Streamable[T] {

	return Streamable[T]{
		Context:    currStr.Context,
		Channel:    callIf(currStr.Context, currStr.Channel, CallWhen, callback),
		BufferSize: currStr.BufferSize,
	}

}

// Initiates data stream from Sql Rows. needs an established db connection. query statement and args.
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

// loops through the stream of data and process the items.
func (currStr Streamable[T]) Process(Process func(item T)) {
	for item := range currStr.Channel {
		Process(item)
	}
}

// loops through the stream of data and process the items in a seperate goroutine.
func (currStr Streamable[T]) BackgroundProcess(waitGroup *sync.WaitGroup, Process func(item T)) {

	go func() {
		defer waitGroup.Done()

		for item := range currStr.Channel {

			Process(item)
		}

	}()

}

// Pipes the stream result.
func (currStr Streamable[T]) Pipe() <-chan T {
	return currStr.Channel
}
