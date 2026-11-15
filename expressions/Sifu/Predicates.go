package Sifu

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

// Predicate compiles the expression
func (curr KeySelectorExpression[T, K]) Predicate() func(T) K {
	return curr.result
}

// Predicate compiles the expression
func (curr CompareOperation[T]) Predicate() func(T, T) bool {
	return curr.result
}

// Predicate compiles the expression
func (op ExpressionEvaluation[T]) Predicate() func(T) bool {
	return op.result
}

// Predicate compiles the expression
func (op MutableExpression[T]) Predicate() func(item T) T {
	return op.result
}
