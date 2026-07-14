package Sifu

type Expression[T any] interface {
	OfType()
}

type TypeExpression[T any] struct {
	op []PropExpression[T]
}

type PropExpression[T any] struct {
	FieldName string
}

type ExpressionEvaluation[T any] struct {
	result func(item T) bool
}
