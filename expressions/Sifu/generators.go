package Sifu

func (curr KeySelectorExpression[T, K]) Gen() func(T) K {
	return curr.Result
}
func (curr CompareOperation[T]) Gen() func(T, T) bool {
	return curr.Result
}

func (curr ExpressionEvaluation[T]) Gen() func(T) bool {
	return curr.Result
}

func (op MutableExpression[T]) Gen() func(item T) T {
	return op.Result
}
