package streams

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

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

	for v := range FromData[ComplexObjectToSearch](ctx, items).FilterStream(func(search ComplexObjectToSearch) bool {

		return search.Id > 2
	}).Throttle(0).TakeAll() {

		fmt.Println(v)

		count++

		if count == 100000 {
			cancel()
			break
		}
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

	for v := range FromChannel[ComplexObjectToSearch](ctx, channel).FilterStream(func(complex ComplexObjectToSearch) bool {
		return complex.Id > 2
	}).Throttle(time.Millisecond * 500).TakeAll() {

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

	for i := range FromCsv(ctx, CsvStreamConfig).FilterStream(func(c customer) bool {
		return c.Index > 40
	}).Throttle(250 * time.Millisecond).TakeAll() {
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
		takeAll[customer](ctx,
			filterStream(ctx, 5,
				fromCsv(ctx, CsvStreamConfig), func(customer customer) bool {
					//this is filter on data
					return customer.Index > 0
				}))

	for _, v := range result2 {
		fmt.Println(v.Index)
		fmt.Println(v.FirstName)
	}

}

func TestStreamFromCsv4(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var CsvStreamConfig contracts.CsvStreamConf[customer]

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "../customers-100.csv"

	CsvStreamConfig.BufferSize = 8

	CsvStreamConfig.ItemCount = 8

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 3 {
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
		takeAll[customer](ctx,
			filterStream(ctx, CsvStreamConfig.BufferSize,
				fromCsv(ctx, CsvStreamConfig), func(customer customer) bool {
					return customer.Index > 0
				}))

	for _, v := range res {
		fmt.Println(v)

	}

}

func TestStreamFromCsv5(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var CsvStreamConfig contracts.CsvStreamConf[customer]

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "../customers-100.csv"

	CsvStreamConfig.BufferSize = 8

	CsvStreamConfig.ItemCount = 8

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 4 {
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

	res := FromCsv(ctx, CsvStreamConfig).FilterStream(func(c customer) bool {
		return c.Index > 0
	}).TakeAll()

	for _, v := range res {
		fmt.Println(v)

	}

}

func TestStreamFromCsv6(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var CsvStreamConfig contracts.CsvStreamConf[customer]

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "../customers-100.csv"

	CsvStreamConfig.BufferSize = 8

	CsvStreamConfig.ItemCount = 8

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 4 {
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

	data := FromCsv(ctx, CsvStreamConfig).FilterStream(func(c customer) bool {
		return c.Index > 0
	}).TakeAll()

	fmt.Println(data)

	collected := TCollection.Group[int, customer](TCollection.From(data), func(t customer) int {
		return t.Index
	}).Collect()

	for k, v := range collected.Items {
		fmt.Println(k)
		fmt.Println(v)
	}
}
