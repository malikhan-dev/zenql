package streams

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"

	"github.com/malikhan-dev/zenq/contracts"
)

func CompileFromQueryable[T any](items []T) *contracts.CompiledQueryable[T] {

	var result contracts.CompiledQueryable[T]

	result.Operators = make([]contracts.LingoOperator[T], 0)

	result.Items = &items

	var operator contracts.LingoOperator[T]

	operator.OperatorType = 1

	operator.MetaData = contracts.OpData[T]{
		MetaData: "FromQueryable",
		Function: func(item T) bool {
			return true
		},
	}
	return &result
}

func fromData[T any](ctx context.Context, BufferSize int, items []T) <-chan T {
	out := make(chan T, BufferSize)

	go func() {
		defer close(out)

		for _, v := range items {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()

	return out
}

func fromChannel[T any](ctx context.Context, BufferSize int, items <-chan T) <-chan T {
	out := make(chan T, BufferSize)

	go func() {
		defer close(out)

		for val := range items {
			select {
			case <-ctx.Done():
				return
			case out <- val:
			}
		}
	}()

	return out
}

func fromCsv[T any](ctx context.Context, conf contracts.CsvStreamConf[T]) (<-chan T, error) {
	out := make(chan T, conf.BufferSize)

	f, err := os.Open(conf.FilePath)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(out)

		defer f.Close()

		reader := csv.NewReader(f)
		rowCounter := 0

		if !conf.StreamHeaders {
			if _, err := reader.Read(); err != nil {
				return
			}
		}

		for {

			if conf.ItemCount > 0 && rowCounter >= conf.ItemCount {
				break
			}

			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {

				break
			}

			rowCounter++

			v, perr := conf.Parser(row)
			if perr != nil {
				if conf.ParseErrorCallback != nil {
					conf.ParseErrorCallback(perr, rowCounter)
				}
				continue
			}

			select {
			case <-ctx.Done():
				return
			case out <- v:

			}
		}
	}()

	return out, nil
}

func fromJsonArr[T any](ctx context.Context, conf contracts.StreamConf) (<-chan T, error) {
	out := make(chan T, conf.BufferSize)

	file, err := os.Open(conf.FilePath)

	if err != nil {
		return nil, err
	}

	go func() {

		defer close(out)

		defer file.Close()

		dec := json.NewDecoder(file)

		_, err = dec.Token()

		if err != nil {
			return
		}

		rowCounter := 0

		for dec.More() {
			select {
			case <-ctx.Done():
				return
			default:
				var item T
				err := dec.Decode(&item)
				if err != nil {
					rowCounter++
					conf.ParseErrorCallback([]error{err}, rowCounter)
				}
				out <- item
			}
		}

	}()
	return out, nil
}
