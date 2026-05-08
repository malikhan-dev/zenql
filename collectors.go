package lingo

import (
	"fmt"
)

func (query *Queryable[T]) Collect() ([]T, []OpError) {
	return query.Items, query.err
}

func (query *Queryable[T]) CollectRange(cnt int) ([]T, []OpError) {

	if len(query.Items) >= cnt {
		return query.Items[0:cnt], query.err
	} else {
		query.err = append(query.err, ErrFactory(5, fmt.Sprintf("CollectRange(%d)", cnt)))
	}
	return nil, query.err

}
