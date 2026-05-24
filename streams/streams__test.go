package streams

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/malikhan-dev/zenq/contracts"
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

	type UserDTO struct {
		ID       int    `json:"id" csv:"id"`
		Name     string `json:"name" csv:"name"`
		Age      int    `json:"age" csv:"age"`
		Email    string `json:"email" csv:"email"`
		IsActive bool   `json:"is_active" csv:"is_active"`
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var CsvStreamConfig contracts.CsvStreamConf[UserDTO]

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "users_data2.csv"

	CsvStreamConfig.BufferSize = 256

	CsvStreamConfig.ParseErrorCallback = func(err []error, i int) {

		fmt.Println(err, " at", i)

		if i > 3 {

			cancel()
		}

	}

	CsvStreamConfig.Parser = func(row []string) (UserDTO, []error) {

		var errorList []error

		index, err := strconv.Atoi(row[0])

		if err != nil {
			errorList = append(errorList, err)
		}
		age, err2 := strconv.Atoi(row[2])

		if err2 != nil {
			errorList = append(errorList, err2)
		}

		active, err3 := strconv.ParseBool(row[4])

		if err3 != nil {
			errorList = append(errorList, err3)
		}
		return UserDTO{
			ID:       index,
			Name:     row[1],
			Age:      age,
			Email:    row[3],
			IsActive: active,
		}, errorList
	}

	data := FromCsv(ctx, CsvStreamConfig).FilterStream(func(c UserDTO) bool {
		return c.ID > 0
	}).Channel

	for v := range data {
		fmt.Println(" value: ", v)
	}

}

func TestStreamFromJson1(t *testing.T) {

	type User struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		IsActive  bool   `json:"is_active"`
		CreatedAt string `json:"created_at"`
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var jsonStreamConfig contracts.JsonStreamConf

	jsonStreamConfig.FilePath = "users_data.json"

	jsonStreamConfig.BufferSize = 256

	jsonStreamConfig.ParseErrorCallback = func(err []error, i int) {

		fmt.Println(err, " at", i)

		if i > 3 {
			cancel()
		}
	}

	data := FromJsonArr[User](ctx, jsonStreamConfig.StreamConf).FilterStream(func(c User) bool {
		return c.ID > 0
	})

	for v := range data.Channel {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(" value: ", v)
	}

}
