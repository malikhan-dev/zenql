package lingo

import (
	"fmt"
)

func (query *Queryable[T]) Collect() ([]T, []OpError) {
	return query.Items, query.Err
}

func (query *Queryable[T]) CollectRange(cnt int) ([]T, []OpError) {

	if len(query.Items) >= cnt {
		return query.Items[0:cnt], query.Err
	} else {
		query.Err = append(query.Err, ErrFactory(5, fmt.Sprintf("CollectRange(%d)", cnt)))
	}
	return nil, query.Err

}

func (query *Queryable[T]) CollectChan(BufferSize int) <-chan CollectStream[T] {

	ch := make(chan CollectStream[T], BufferSize)

	go func() {

		defer close(ch)

		for _, v := range query.Items {
			ch <- CollectStream[T]{Value: v}
		}
		for _, v := range query.Err {
			ch <- CollectStream[T]{Err: v}
		}

	}()
	return ch
}

func (query *GroupedQueryable[K, T]) Collect() (map[K][]T, []OpError) {
	return query.Items, query.Err
}
