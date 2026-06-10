package databases

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"database/sql"
	"errors"
	"reflect"
)

func mapRows[T any](rows *sql.Rows, singleExec bool) ([]T, error) {

	var itemList []T = make([]T, 0, 128)

	rowCount := 0
	for rows.Next() {
		rowCount++
		if singleExec {
			if rowCount > 1 {
				return nil, errors.New("multiple rows found")
			}
		}
		mapped, err := mapRow[T](rows, singleExec)
		if err == nil {
			itemList = append(itemList, mapped)
		}

	}
	return itemList, nil
}

func mapRow[T any](rows *sql.Rows, singleExec bool) (T, error) {
	var item T
	val := reflect.ValueOf(&item).Elem()
	typ := val.Type()
	columns, err := rows.Columns()
	if err != nil {
		return item, err
	}

	var columnIndexCache []int
	cacheBuilt := false

	if !cacheBuilt {
		columnIndexCache = make([]int, len(columns))

		for i, colName := range columns {
			foundIndex := -1
			for j := 0; j < typ.NumField(); j++ {
				if typ.Field(j).Tag.Get("zql") == colName {
					foundIndex = j
					break
				}
			}

			if foundIndex == -1 {
				if f, ok := typ.FieldByName(colName); ok {
					foundIndex = f.Index[0]
				}
			}
			columnIndexCache[i] = foundIndex
		}
		cacheBuilt = true
	}

	scanArgs := make([]any, len(columns))
	for i := range columns {
		fieldIdx := columnIndexCache[i]

		if fieldIdx != -1 {
			field := val.Field(fieldIdx)
			if field.CanSet() {
				scanArgs[i] = field.Addr().Interface()
				continue
			}
		}

		var ignore any
		scanArgs[i] = &ignore
	}

	if err := rows.Scan(scanArgs...); err != nil {
		return item, err
	}

	return item, nil
}
