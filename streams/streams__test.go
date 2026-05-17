package streams

import (
	"github.com/malikhan-dev/lingo/collections"
	"context"
	"fmt"
	"testing"
)

type ComplexObjectToSearch struct {
	Name string
	Age  int
	Id   int
	Flag bool
}

var items []ComplexObjectToSearch

const Heavy_Load = true

func init() {

	items = []ComplexObjectToSearch{
		ComplexObjectToSearch{
			Name: "John",
			Age:  20,
			Id:   1,
			Flag: true,
		},
		ComplexObjectToSearch{
			Name: "Jane",
			Age:  20,
			Id:   2,
			Flag: false,
		},
		ComplexObjectToSearch{
			Name: "Jane",
			Age:  20,
			Id:   3,
			Flag: true,
		},
		ComplexObjectToSearch{
			Name: "jack",
			Age:  20,
			Id:   4,
			Flag: true,
		},
	}

	if Heavy_Load {
		LoadLargeData()
	}

}

func LoadLargeData() {
	randFlag := false
	for i := 0; i < 200; i++ {

		items = append(items, ComplexObjectToSearch{
			Name: "Jane",
			Flag: randFlag,
			Id:   i,
			Age:  i,
		})
		randFlag = !randFlag
	}

}

func TestStreamsFromData(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	count := 0

	var buffer_size int

	buffer_size = 10

	for v := range Throttle(ctx, FilterStream(ctx, buffer_size, FromData(ctx, buffer_size, items), func(item ComplexObjectToSearch) bool {
		return item.Id > 2
	}), 0) {

		fmt.Println(v)

		count++

		if count == 100000 {
			cancel()
			break
		}
	}

}

func TestStreamsFromQueryable(t *testing.T) {

	// check for errors before streaming

	type SimplerType struct {
		Enabled bool
		Id      int
		Name    string
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	count := 0

	var buffer_size int

	buffer_size = 10

	queryable := collections.From(items)

	mappedStream := MapStream[ComplexObjectToSearch, SimplerType](ctx,
		Throttle(ctx,
			FilterStream(ctx, buffer_size,
				collections.FromQueryable(ctx, buffer_size, *queryable),
				func(item ComplexObjectToSearch) bool {

					return item.Id > 0

				}), 0), func(search ComplexObjectToSearch) SimplerType {

			return SimplerType{
				Enabled: search.Flag,
				Id:      search.Id,
				Name:    search.Name,
			}
		})

	for v := range mappedStream {

		fmt.Println(v)

		count++

	}

}

func TestStreamsFromThorQueryable(t *testing.T) {

	// check for errors before streaming

	type SimplerType struct {
		Enabled bool
		Id      int
		Name    string
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	count := 0

	var buffer_size int

	buffer_size = 10

	queryable := TCollection.From(items)

	mappedStream := MapStream[ComplexObjectToSearch, SimplerType](ctx,
		Throttle(ctx,
			FilterStream(ctx, buffer_size,
				TCollection.FromQueryable(ctx, buffer_size, *queryable),
				func(item ComplexObjectToSearch) bool {

					return item.Id > 0

				}), 0), func(search ComplexObjectToSearch) SimplerType {

			return SimplerType{
				Enabled: search.Flag,
				Id:      search.Id,
				Name:    search.Name,
			}
		})

	for v := range mappedStream {

		fmt.Println(v)

		count++

	}

}

func TestStreamsFromChannel(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	count := 0

	var buffer_size int

	buffer_size = 10

	channel := make(chan ComplexObjectToSearch, buffer_size)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- ComplexObjectToSearch{
				Name: "Jack",
				Flag: true,
				Id:   i,
				Age:  i,
			}
		}
		close(channel)
	}()

	for v := range Throttle(ctx,
		FilterStream(ctx, buffer_size, FromChannel(ctx, buffer_size, channel),
			func(item ComplexObjectToSearch) bool {
				return item.Id > 2
			}), 0) {

		fmt.Println(v)

		count++

		if count == 10 {
			cancel()
			break
		}
	}

}
