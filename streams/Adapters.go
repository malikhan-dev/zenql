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

func fromCsv[T any](ctx context.Context, conf contracts.CsvStreamConf[T]) <-chan T {
	out := make(chan T, conf.BufferSize)

	go func() {
		defer close(out)

		f, err := os.Open(conf.FilePath)
		if err != nil {
			// optionally call an error callback here
			return
		}
		defer f.Close()

		reader := csv.NewReader(f)
		rowCounter := 0

		// skip header if needed
		if !conf.StreamHeaders {
			if _, err := reader.Read(); err != nil {
				return
			}
		}

		for {
			// check item limit
			if conf.ItemCount > 0 && rowCounter >= conf.ItemCount {
				break
			}

			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				// read error, not parse error
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

			// ✅ اینجا اگه channel پر باشه، صبر میکنه
			// و اگه ctx cancel بشه، خارج میشه
			select {
			case <-ctx.Done():
				return // ← return نه break، از goroutine کامل خارج میشه
			case out <- v:
				// sent successfully, or waited until space was available
			}
		}
	}()

	return out // ← فوری برمیگرده، goroutine در پس‌زمینه کار میکنه
}
