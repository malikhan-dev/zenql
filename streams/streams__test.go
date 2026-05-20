package streams

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/malikhan-dev/zenq/collections"
	"github.com/malikhan-dev/zenq/contracts"

	TCollection "github.com/malikhan-dev/zenq/collections/Thor"
)

type ComplexObjectToSearch struct {
	Name string
	Age  int
	Id   int
	Flag bool
}

type customer struct {
	Index            int
	CustomerId       string
	FirstName        string
	LastName         string
	Company          string
	City             string
	Country          string
	Phone1           string
	Phone2           string
	Email            string
	SubscriptionDate string
	Website          string
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
	for i := 0; i < 50; i++ {

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

func TestStreamFromCsv(t *testing.T) {

	type customer struct {
		Index            int
		CustomerId       string
		FirstName        string
		LastName         string
		Company          string
		City             string
		Country          string
		Phone1           string
		Phone2           string
		Email            string
		SubscriptionDate string
		Website          string
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var CsvStreamConfig contracts.CsvStreamConf[customer]

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "../customers-100.csv"

	CsvStreamConfig.BufferSize = 256

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 0 {
			cancel()
		}
	}

	CsvStreamConfig.Parser = func(row []string) (customer, error) {
		index, err := strconv.Atoi(row[0])

		return customer{
			CustomerId:       row[1],
			Index:            index,
			FirstName:        row[2],
			LastName:         row[3],
			Company:          row[4],
			City:             row[5],
			Country:          row[6],
			Phone1:           row[7],
			Phone2:           row[8],
			Email:            row[9],
			SubscriptionDate: row[10],
			Website:          row[11],
		}, err
	}

	for i := range Throttle(ctx, FilterStream(ctx, 5,
		FromCsv(ctx, CsvStreamConfig), func(customer customer) bool {
			//this is filter on data
			return customer.Index > 40
		}), 250*time.Millisecond) {
		fmt.Println(i)
	}

}

func TestStreamFromCsv2(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var CsvStreamConfig contracts.CsvStreamConf[customer]

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "../customers-100.csv"

	CsvStreamConfig.BufferSize = 256

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 0 {
			cancel()
		}
	}

	CsvStreamConfig.Parser = func(row []string) (customer, error) {
		index, err := strconv.Atoi(row[0])
		return customer{
			CustomerId:       row[1],
			Index:            index,
			FirstName:        row[2],
			LastName:         row[3],
			Company:          row[4],
			City:             row[5],
			Country:          row[6],
			Phone1:           row[7],
			Phone2:           row[8],
			Email:            row[9],
			SubscriptionDate: row[10],
			Website:          row[11],
		}, err
	}

	result2 :=
		TakeAll[customer](ctx,
			FilterStream(ctx, 5,
				FromCsv(ctx, CsvStreamConfig), func(customer customer) bool {
					//this is filter on data
					return customer.Index > 0
				}))

	for _, v := range result2 {
		fmt.Println(v.Index)
		fmt.Println(v.FirstName)
	}

}

func TestStreamFromCsv3(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var CsvStreamConfig contracts.CsvStreamConf[customer]

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "../customers-100.csv"

	CsvStreamConfig.BufferSize = 256

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 0 {
			cancel()
		}
	}

	CsvStreamConfig.Parser = func(row []string) (customer, error) {
		index, err := strconv.Atoi(row[0])
		return customer{
			CustomerId:       row[1],
			Index:            index,
			FirstName:        row[2],
			LastName:         row[3],
			Company:          row[4],
			City:             row[5],
			Country:          row[6],
			Phone1:           row[7],
			Phone2:           row[8],
			Email:            row[9],
			SubscriptionDate: row[10],
			Website:          row[11],
		}, err
	}

	res :=
		TCollection.Collect(TCollection.Group[int, customer](
			TCollection.From(
				TakeAll[customer](ctx,
					FilterStream(ctx, CsvStreamConfig.BufferSize,
						FromCsv(ctx, CsvStreamConfig), func(customer customer) bool {
							return customer.Index > 60
						}))),

			func(c customer) int {
				return c.Index
			},
		))

	for k, v := range res.Items {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("=================")
	}

}
