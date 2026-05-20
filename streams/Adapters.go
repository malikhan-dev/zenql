package streams

import (
	"context"
	"encoding/csv"
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

func FromData[T any](ctx context.Context, BufferSize int, items []T) <-chan T {
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

func FromChannel[T any](ctx context.Context, BufferSize int, items <-chan T) <-chan T {
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

func FromCsv[T any](ctx context.Context, Conf contracts.CsvStreamConf[T]) <-chan T {
	out := make(chan T, Conf.BufferSize)

	defer close(out)

	f, err := os.Open(Conf.FilePath)

	if err == nil {

		defer f.Close()

		reader := csv.NewReader(f)

		var rowCounter int

		rowCounter = 0

		for {

			if rowCounter == 0 {

				if !Conf.StreamHeaders {

					reader.Read()
					rowCounter++
					continue
				}

			} else {
				rowCounter++

				row, err := reader.Read()

				if err == io.EOF {
					break
				}

				v, perr := Conf.Parser(row)

				if perr == nil {
					out <- v
				} else {
					Conf.ParseErrorCallback(err, rowCounter)
				}
			}

			if Conf.ItemCount > 0 && rowCounter == Conf.ItemCount+1 {
				break

			}

		}
	}

	return out

}
