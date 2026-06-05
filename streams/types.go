package streams

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import "context"

type Streamable[T any] struct {
	Channel    <-chan T
	Context    context.Context
	BufferSize int
	Err        []error
	Initiated  bool
}
