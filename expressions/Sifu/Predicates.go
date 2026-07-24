package Sifu

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

func (curr KeySelectorExpression[T, K]) Predicate() func(T) K {
	return curr.result
}

func (curr CompareOperation[T]) Predicate() func(T, T) bool {
	return curr.result
}

func (op ExpressionEvaluation[T]) Predicate() func(T) bool {
	return op.result
}

func (op MutableExpression[T]) Predicate() func(item T) T {
	return op.result
}
