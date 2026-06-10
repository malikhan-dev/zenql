package collections

import "context"

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

func FromQueryable[T any](ctx context.Context, BufferSize int, items CollectionCompiledQueryable[T]) <-chan T {
	out := make(chan T, BufferSize)

	go func() {
		defer close(out)

		for _, v := range *items.Items {

			select {
			case <-ctx.Done():
				return
			case out <- v:
			}

		}

	}()
	return out
}
