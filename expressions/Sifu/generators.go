package Sifu

func (curr KeySelectorExpression[T, K]) Gen() func(T) K {
	return curr.result
}

func (curr CompareOperation[T]) Gen() func(T, T) bool {
	return curr.result
}

func (curr ExpressionEvaluation[T]) Gen() func(T) bool {
	return curr.result
}

func (op MutableExpression[T]) Gen() func(item T) T {
	return op.Result
}
