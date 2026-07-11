package contracts

type KeySelectGenerator[T, K comparable] interface {
	Gen() func(T) K
}

type CompareGenerator[T comparable] interface {
	Gen() func(T, T) bool
}

type ExpressionGenerator[T any] interface {
	Gen() func(T) bool
}

type MutableExpressionGenerator[T any] interface {
	Gen() func(T) T
}
