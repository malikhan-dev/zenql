package streams

import "context"

type Streamable[T any] struct {
	Channel <-chan T
	Context context.Context
}
