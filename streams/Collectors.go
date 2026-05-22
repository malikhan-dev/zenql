package streams

import (
	contracts "github.com/malikhan-dev/zenq/contracts"
	"context"
)

func CompileStreamWithMapping[T any, V any](ctx context.Context, currentOps *contracts.CompiledQueryable[T], mapper func(items T) V) <-chan V {

	var channel = make(chan V)

	go func() {
		defer close(channel)

		for _, v := range *currentOps.Items {
			Keep := true
			select {
			case <-ctx.Done():
				return
			default:
			}
			for _, op := range currentOps.Operators {

				switch op.OperatorType {

				case BuildQueryable:
					if !op.MetaData.Function(v) {
						Keep = false
						break
					}

				case FilterQueryable:

					if !op.MetaData.Function(v) {
						Keep = false
						break
					}

				case ThrottleQueryable:

					if op.MetaData.MetaData != "0" {
						if op.MetaData.Function(v) {

						}
					}

				}

			}
			if Keep {
				channel <- mapper(v)
			}

		}

	}()

	return channel
}

func CompileStream[T any](ctx context.Context, currentOps *contracts.CompiledQueryable[T]) <-chan T {

	var channel = make(chan T)

	go func() {
		defer close(channel)

		for _, v := range *currentOps.Items {
			Keep := true
			select {
			case <-ctx.Done():
				return
			default:

			}
			for _, op := range currentOps.Operators {

				switch op.OperatorType {

				case BuildQueryable:
					if !op.MetaData.Function(v) {
						Keep = false
						break
					}

				case FilterQueryable:

					if !op.MetaData.Function(v) {
						Keep = false
						break
					}

				case ThrottleQueryable:

					if op.MetaData.MetaData != "0" {
						if op.MetaData.Function(v) {

						}
					}

				}

			}
			if Keep {
				channel <- v
			}

		}

	}()

	return channel
}

