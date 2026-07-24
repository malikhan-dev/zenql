package collections

import (
	"context"
	"sort"
	"sync"

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
	SortCollection     = 10
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

func (op *CollectionCompiledQueryable[T]) Take(count int) *CollectionCompiledQueryable[T] {

	op.Page.Limit = count

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: TakeCollection,
	})
	return op
}
func (op *CollectionCompiledQueryable[T]) Skip(count int) *CollectionCompiledQueryable[T] {

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
		CollectionCompiler: op.Collect,
		PropLocator:        locator,
		Page:               op.Page,
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

func (op *CollectionCompiledQueryable[T]) WhereEx(expr contracts.ExpressionPredicate[T]) *CollectionCompiledQueryable[T] {
	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: WhereCollection,
		Filter:       contracts.Filterer[T]{Function: expr.Predicate()},
	})
	return op
}

func (op *CollectionCompiledQueryable[T]) AnyEx(expr contracts.ExpressionPredicate[T]) *AssertCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: AnyCollection,
		Filter: contracts.Filterer[T]{
			Function: expr.Predicate(),
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

func (op *CollectionCompiledQueryable[T]) UpdateEx(Updater contracts.MutableExpressionPredicate[T]) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: UpdateCollection,
		Update:       contracts.Updater[T]{Function: Updater.Predicate()},
	})

	return op

}

func (op *CollectionCompiledQueryable[T]) Sort(Less func(item1 T, item2 T) bool, Desc bool) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: SortCollection,
		Sort:         contracts.Sorter[T]{Function: Less, Desc: Desc},
	})

	return op

}

func (op *CollectionCompiledQueryable[T]) SortEx(Less contracts.ComparePredicate[T], Desc bool) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: SortCollection,
		Sort:         contracts.Sorter[T]{Function: Less.Predicate(), Desc: Desc},
	})

	return op

}

func (op *CollectionCompiledQueryable[T]) Collect() []T {

	var result []T

	result = contracts.AllocateSlice[T](len(*op.Items))

	skipLimit, takeLimit := op.Page.Skip, op.Page.Limit

	var skipCount int

	skipCount = 0

	var HasUpdate bool

	var SortDescending bool

	var UpdateFunc func(T) T

	var FilterFunc func(T) bool

	var HasSort bool

	var SortFunc func(T, T) bool

	var resultCounter int

	resultCounter = 0

	var wg sync.WaitGroup

	wg.Add(3)

	var slice []T

	slice = *op.Items

	go func() {

		HasSort, SortDescending, SortFunc = ExtractSortMeta(op.Operators)

		if HasSort {
			sort.Slice(slice, func(i, j int) bool {
				if SortDescending {
					return SortFunc(slice[j], slice[i])
				}
				return SortFunc(slice[i], slice[j])
			})

		}

		wg.Done()

	}()

	go func() {
		HasUpdate, UpdateFunc = ExtractUpdateMeta(op.Operators)
		wg.Done()
	}()

	go func() {
		FilterFunc = ExtractFilterMeta(op.Operators)
		wg.Done()
	}()

	wg.Wait()

	hasTake := takeLimit != -1

	for _, item := range slice {

		keep := true

		keep = FilterFunc(item)

		if keep {

			if skipCount < skipLimit {
				skipCount++
				continue
			}

			if hasTake {
				if resultCounter == takeLimit {
					return result
				}
			}
			if HasUpdate {
				item = UpdateFunc(item)
			}
			result = append(result, item)
			resultCounter++
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

	compiledResult := op.CollectionCompiler()

	result.Items = contracts.AllocateMap[K, T](len(compiledResult))

	var LocatedKey K

	for _, item := range compiledResult {

		LocatedKey = op.PropLocator(item)

		result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
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

func ExtractSortMeta[T any](op []contracts.ZenqlOperator[T]) (bool, bool, func(T, T) bool) {

	for _, op := range op {
		if op.OperatorType == SortCollection {
			return true, op.Sort.IsDescending(), op.Sort.Sort

		}
	}
	return false, false, func(item T, item2 T) bool {
		return true
	}
}
