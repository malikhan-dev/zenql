package contracts

type KeySelectPredicate[T, K comparable] interface {
	Predicate() func(T) K
}

type ComparePredicate[T any] interface {
	Predicate() func(T, T) bool
}

type ExpressionPredicate[T any] interface {
	Predicate() func(T) bool
}

type MutableExpressionPredicate[T any] interface {
	Predicate() func(T) T
}
