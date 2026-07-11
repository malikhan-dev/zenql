package collections

import (
	"context"
	"sort"

	"github.com/malikhan-dev/zenql/contracts/v2"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */
type OperatorType int

const (
	FromItems          = 1
	WhereCollection    = 2
	AnyCollection      = 4
	GroupCollection    = 5
	DistinctCollection = 6
	TakeCollection     = 7
	SkipCollection     = 8
	UpdateCollection   = 9
)

func CoreFilter[T any](Operator contracts.ZenqlOperator[T], item T) bool {

	ShouldKeep := true

	switch Operator.OperatorType {

	case WhereCollection:
		if !Operator.Filter.Filter(item) {
			ShouldKeep = false
			break
		}

	}
	return ShouldKeep
}

func (op *CollectionCompiledQueryable[T]) Take(count int32) *CollectionCompiledQueryable[T] {

	op.Page.Limit = count

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: TakeCollection,
	})
	return op
}
func (op *CollectionCompiledQueryable[T]) Skip(count int32) *CollectionCompiledQueryable[T] {

	op.Page.Skip = count

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: SkipCollection,
	})
	return op
}
func Group[K comparable, T any](op *CollectionCompiledQueryable[T], locator func(T) K) *GroupCompiledQueryable[K, T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: GroupCollection,
	})
	return &GroupCompiledQueryable[K, T]{
		CompiledQueryable: op.CompiledQueryable,
		PropLocator:       locator,
		Page:              op.Page,
	}
}
func (op *AssertCompiledQueryable[T]) Assert() bool {

	var fnc func(T) bool
	for _, op := range op.Operators {

		switch op.OperatorType {
		case AnyCollection:
			fnc = op.Filter.Filter
		}
	}

	for _, item := range *op.Items {

		if fnc(item) {
			return true
		}
	}
	return false
}
func From[T any](items *[]T) *CollectionCompiledQueryable[T] {

	initiateOperator := make([]contracts.ZenqlOperator[T], 0)
	initiateOperator = append(initiateOperator, contracts.ZenqlOperator[T]{
		OperatorType: FromItems,
	})
	queryData := contracts.CompiledQueryable[T]{
		Items:     items,
		Operators: initiateOperator,
	}

	return &CollectionCompiledQueryable[T]{
		queryData,
		contracts.PageOption{
			Limit: -1,
			Skip:  -1,
		},
	}
}

func (op *CollectionCompiledQueryable[T]) WhereEx(expr contracts.ExpressionGenerator[T]) *CollectionCompiledQueryable[T] {
	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: WhereCollection,
		Filter:       contracts.Filterer[T]{Function: expr.Gen()},
	})
	return op
}

func (op *CollectionCompiledQueryable[T]) AnyEx(expr contracts.ExpressionGenerator[T]) *AssertCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: AnyCollection,
		Filter: contracts.Filterer[T]{
			Function: expr.Gen(),
		},
	})
	return &AssertCompiledQueryable[T]{
		op.CompiledQueryable,
	}
}

func (op *CollectionCompiledQueryable[T]) Where(function func(T) bool) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: WhereCollection,
		Filter:       contracts.Filterer[T]{Function: function},
	})
	return op
}
func (op *CollectionCompiledQueryable[T]) Any(function func(T) bool) *AssertCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: AnyCollection,
		Filter: contracts.Filterer[T]{
			Function: function,
		},
	})
	return &AssertCompiledQueryable[T]{
		op.CompiledQueryable,
	}
}

func (op *CollectionCompiledQueryable[T]) Update(Updater func(T) T) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: UpdateCollection,
		Update:       contracts.Updater[T]{Function: Updater},
	})

	return op

}

func (op *CollectionCompiledQueryable[T]) Collect() []T {
	var result []T
	result = contracts.AllocateSlice[T](len(*op.Items))

	skipLimit, takeLimit := op.Page.Skip, op.Page.Limit

	var skipCount, count int32
	skipCount = 0
	count = 0

	HasUpdate, UpdateFunc := ExtractUpdateMeta(op.Operators)

	var FilterFunc func(T) bool

	FilterFunc = ExtractFilterMeta(op.Operators)

	hasTake := takeLimit != -1
	hasSkip := skipLimit != -1

	for _, item := range *op.Items {

		keep := true

		keep = FilterFunc(item)

		if keep {
			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}

			if hasTake {
				if len(result) == int(takeLimit) {
					return result
				}
				if HasUpdate {
					item = UpdateFunc(item)
				}
				result = append(result, item)
				count++

			} else {
				if HasUpdate {
					item = UpdateFunc(item)
				}
				result = append(result, item)
				count++
			}
		}
	}
	return result
}

func (op *CollectionCompiledQueryable[T]) FindParentNode(NodeLocator func(T) bool, Criteria func(child T, parent T) bool) T {

	var result []T

	var TargetNode T

	result = contracts.AllocateSlice[T](len(*op.Items))

	for _, item := range *op.Items {

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		if keep {

			if NodeLocator(item) {
				TargetNode = item
			}
			result = append(result, item)

		}
	}

	var value T
	for _, val := range result {
		if Criteria(TargetNode, val) {
			value = val
			break
		}
	}

	return value
}

func (op *CollectionCompiledQueryable[T]) FindRootNode(Start func(T) bool, Link func(child T, parent T) bool, Less func(T, T) bool) T {

	var result []T

	var TargetNode T

	result = contracts.AllocateSlice[T](len(*op.Items))

	for _, item := range *op.Items {

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		if keep {

			if Start(item) {
				TargetNode = item
			}
			result = append(result, item)

		}
	}

	sort.Slice(result, func(i, j int) bool {
		return Less(result[j], result[i])
	})

	for _, val := range result {

		if Link(TargetNode, val) {
			TargetNode = val
		}

	}

	return TargetNode
}

func (op *CollectionCompiledQueryable[T]) TraverseRootNode(Start func(T) bool, Link func(child T, parent T) bool, Less func(T, T) bool, ctx context.Context) <-chan T {

	var result []T

	var TargetNode T

	result = contracts.AllocateSlice[T](len(*op.Items))

	for _, item := range *op.Items {

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		if keep {

			if Start(item) {
				TargetNode = item
			}
			result = append(result, item)

		}
	}

	sort.Slice(result, func(i, j int) bool {
		return Less(result[j], result[i])
	})

	out := make(chan T, 1)

	go func() {

		for _, val := range result {

			if Link(TargetNode, val) {
				select {
				case <-ctx.Done():
					break
				case out <- val:
					TargetNode = val
				}
			}
		}
		defer close(out)

	}()

	return out
}

func (op *GroupCompiledQueryable[K, T]) Collect() *GroupedQueryable[K, T] {

	var result GroupedQueryable[K, T]

	result.Items = contracts.AllocateMap[K, T](len(*op.Items))

	skipLimit, takeLimit := op.Page.Skip, op.Page.Limit

	var FilterFunc func(T) bool

	FilterFunc = ExtractFilterMeta(op.Operators)

	var LocatedKey K

	var skipCount, count int32

	for _, item := range *op.Items {

		LocatedKey = op.PropLocator(item)

		keep := FilterFunc(item)

		if keep {
			hasTake := takeLimit != -1
			hasSkip := skipLimit != -1

			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}
			if hasTake {
				if len(result.Items) == int(takeLimit) {
					return &result
				}
				result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
				count++

			} else {
				result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
				count++
			}

		}

	}

	return &result
}

func ExtractUpdateMeta[T any](op []contracts.ZenqlOperator[T]) (bool, func(T) T) {

	var HasUpdate bool
	var UpdateFunc func(T) T

	for _, op := range op {

		if op.OperatorType == UpdateCollection {
			HasUpdate = true
			UpdateFunc = op.Update.Update
			break
		}
	}
	return HasUpdate, UpdateFunc

}

func ExtractFilterMeta[T any](op []contracts.ZenqlOperator[T]) func(T) bool {

	for _, op := range op {
		if op.OperatorType == WhereCollection {
			return op.Filter.Filter

		}
	}
	return func(item T) bool {
		return true
	}
}
