package lingo

import "errors"

func (query *Queryable[T]) Collect() ([]T, []error) {
	return query.Items, query.err
}

func (query *Queryable[T]) CollectRange(cnt int) ([]T, []error) {

	if len(query.Items) >= cnt {
		return query.Items[0:cnt], query.err
	} else {
		query.err = append(query.err, errors.New("Index Out Of Range. CollectRange()."))
	}
	return nil, query.err

}
