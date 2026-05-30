package MySql

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/malikhan-dev/zenq/contracts"
	"github.com/malikhan-dev/zenq/databases"
	"github.com/malikhan-dev/zenq/streams"
)

func (conn *ZenqMySqlDb) NewConnection(constr string) (contracts.RDBMSFacade, error) {

	db, err := sql.Open("mysql", constr)

	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()

	if pingErr != nil {
		db.Close()
		return nil, pingErr
	}

	DbContext := databases.ZenqDB{db}

	return &ZenqMySqlDb{
		DbContext,
	}, nil
}

func (conn *ZenqMySqlDb) Ping() error {
	return conn.Pool.Ping()
}

func (conn *ZenqMySqlDb) Close() error {
	return conn.Pool.Close()
}
func (conn *ZenqMySqlDb) GetPool() *sql.DB {
	return conn.Pool
}

func (conn *ZenqMySqlDb) Query(query string, args ...any) (*sql.Rows, error) {

	return conn.Pool.Query(query, args...)
}

func Query[T any](conn contracts.RDBMSFacade, query string, args ...any) ([]T, error) {

	rows, err := conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return mapRows[T](rows, false)
}

func Exec(conn contracts.RDBMSFacade, query string, args ...any) contracts.Commandesult {

	result, err := conn.GetPool().Exec(query, args...)

	var affected int64
	var error error

	if result == nil {
		affected = 0
		error = err
	} else {
		if aff, er := result.RowsAffected(); err == nil {
			affected = aff
			error = er
		} else {
			affected = 0
			error = err
		}
	}

	return contracts.Commandesult{
		Err:          error,
		RowsAffected: affected,
		TimeStamp:    time.Now(),
	}

}

func SingleQuery[T any](conn contracts.RDBMSFacade, query string, args ...any) ([]T, error) {

	rows, err := conn.Query(query, args...)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	return mapRows[T](rows, true)
}

func mapRows[T any](rows *sql.Rows, singleExec bool) ([]T, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var itemList []T = make([]T, 0)

	var columnIndexCache []int
	cacheBuilt := false

	rowCount := 0
	for rows.Next() {
		rowCount++
		if singleExec {
			if rowCount > 1 {
				return nil, errors.New("multiple rows found")
			}
		}
		var item T
		val := reflect.ValueOf(&item).Elem()
		typ := val.Type()

		if !cacheBuilt {
			columnIndexCache = make([]int, len(columns))

			for i, colName := range columns {
				foundIndex := -1
				for j := 0; j < typ.NumField(); j++ {
					if typ.Field(j).Tag.Get("zdb") == colName {
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
			return nil, err
		}

		itemList = append(itemList, item)

	}
	return itemList, nil
}

func FromMySqlRows[T any](ctx context.Context, conn contracts.RDBMSFacade, query string, Mapper func(rows *sql.Rows) (T, error), args ...any) streams.Streamable[T] {
	stream, err := frmMsqlRows[T](ctx, conn, query, Mapper, args...)
	return streams.Streamable[T]{
		Context:    ctx,
		Channel:    stream,
		BufferSize: 256,
		Err:        []error{err},
		Initiated:  err == nil,
	}
}

func frmMsqlRows[T any](ctx context.Context, conn contracts.RDBMSFacade, query string, Mapper func(rows *sql.Rows) (T, error), args ...any) (<-chan T, error) {

	var rows *sql.Rows
	var err error
	rows, err = conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	channel := make(chan T, 256)

	go func() {

		defer rows.Close()

		defer close(channel)

		for rows.Next() {
			select {
			case <-ctx.Done():
				return
			default:
			}

			item, err := Mapper(rows)
			if err != nil {
				return
			}

			select {
			case <-ctx.Done():
				return
			case channel <- item:
			}
		}

	}()

	return channel, nil
}
