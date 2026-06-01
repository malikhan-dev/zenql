[![Starstruck](https://img.shields.io/badge/GitHub-Starstruck-yellow?style=for-the-badge&logo=github)](https://github.com/users/malikhan-dev/achievements/starstruck)
![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)
![Tests](https://img.shields.io/badge/tests-passing-brightgreen?style=for-the-badge&logo=go)
![Coverage](https://img.shields.io/badge/coverage-75%25-brightgreen?style=for-the-badge)
![Maintained](https://img.shields.io/badge/maintained-yes-brightgreen?style=for-the-badge)
![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)
![Version](https://img.shields.io/badge/version-1.7.5-blue?style=for-the-badge)




# Zen-Q (zenq)

**Expressive, Polymorphic Queries with Streaming Capabilities and a User-Friendly API, inspired by LINQ.** 
 
```go
	if jsonStream:= FromJsonArr[User](ctx, jsonStreamConfig.StreamConf); jsonStream.Initiated {
		 JsonData := jsonStream.FilterStream(func(c User) bool {
			return c.ID > 0
		})
			
			
	for v := range JsonData.Channel {
		 time.Sleep(time.Millisecond * 10)
		 fmt.Println(" value: ", v)
		}
	}

	cursor := FromMySqlRows[UserModel](ctx, conn,"select * from Test.users where id>?", func(rows *sql.Rows) (UserModel, error){
		var id, age int
		var name string
		var err error
		err = rows.Scan(&id, &name, &age)
		model := UserModel {
			UserId:   id,
			Age:      age,
			UserName: name,
		}

		return model, err
	}, id)

	if cursor.Initiated {
		for v := range cursor.FilterStream(func(model UserModel) bool {
			return model.Age > 25
		}).Throttle(time.Millisecond * 1000).Channel {

				/// business logic

			}
	‌}

```




# License (MIT)

This library was written and designed by Mohammadreza Malikhan. The source code is free to use with proper attribution. This project is licensed under the MIT License (see the `LICENSE` file for details).

# Intro

zenq is a DSL (Domain Specific Language) for Go that helps you filter, search, validate, process, and stream your data in a fluent and readable way. It is inspired by LINQ in C# and Streams in Java, while staying practical for Go developers. make sure you review the benchmarks section at the end of this document. 

At its core, zenq is a modular library. Currently, it has two modules: **Collections** and **Streams**. streams used to initiate communications with async data-sources, such as a csv file or a MySql database. 

There are two ways of processing collections:
1. Using default APIs.
2. Using the advanced collection query engine known as **Thor**. 

Thor is designed and architected to provide the maximum performance possible. It uses the operation fusion pattern to provide maximum speed and run the entire query chain in a single execution unit. Streams, on the other hand, use famous Golang concepts such as channels and goroutines to allow the user to stream data while respecting the cancellation concepts of Go. at the moment zenq operations allowed on various data sources, in-memory slices, channels, csv or json files and MySql Database. more and more data-sources will be supported soon.


## Installation
you can install the package using the commands below.

``` bash
go get github.com/malikhan-dev/zenq@latest

go mod tidy
```

## Default Collections API
Default Collections APIs are the old ways of processing collections, like filtering them, grouping them and etc...

import path

``` go
collections  "github.com/malikhan-dev/zenq/collections"
```

### `Queryable[T]`

`Queryable[T]` is the core type passed between chained operations such as `Where`, `First`, `FirstOrDefault`, `All`, and `AllOrDefault`.

It wraps:
- A data slice: `[]T`
- An error slice: `[]error`

Collectors unwrap this type into concrete results.

```go
type Queryable[T any] struct {
    Items []T
    Err   []OpError
}
```

---

### `From([]T)`

`From([]T)` creates a `Queryable[T]` from a slice and is usually the starting point of a query chain. It accepts a slice of `[]T` and returns a pointer to `Queryable[T]`.

---

### `Where()`

`Where(fieldName, fieldValue)` filters a slice using a field name and value.
- `fieldName` must be a `string`.
- `fieldValue` can be any type, but it must exactly match the actual type of the target field.

This function modifies the current `Queryable[T]` and returns the same pointer for further chaining.

``` go
_, err2 := From(items).Where("Name", "John").Where("Flag", true).FirstOrDefault().Collect()
```

**Important:** The field value must be exactly the same type as the struct field.
For example, if the field type is `uint32`, you must pass `uint32(2)` instead of `2`.

``` go
_, err := From(Examples).Where("Id", uint32(2)).AllOrDefault().Collect()
```
---

### `First()` and `FirstOrDefault()`

These functions return the first item in the current query chain.
- `First()` panics if no item is found.
- `FirstOrDefault()` appends an error instead of panicking.

Both still return a pointer to `Queryable[T]`.

---

### `All()` and `AllOrDefault()`

These functions return all items in the current query chain.
- `All()` panics if no item is found.
- `AllOrDefault()` appends an error instead of panicking.

Both still return a pointer to `Queryable[T]`.

---

## Collectors

**Available since version `v1.3.2`**

After a chained operation such as:

``` go
zenq.From(data).Where(...).AllOrDefault()
```

You can use collectors to unwrap the `Queryable[T]` result into concrete values.

- `Collect()` returns the full result set and errors.
- `CollectRange(cnt)` returns a limited number of items based on the `cnt` argument, along with errors.
- `Pipe(buffersize)` (formerly `CollectChan(buffersize)`) collects data and errors using Go channels for large datasets. Available since version `v1.4.0`.

``` go
res, err := From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {
    return item.Id > 200000
}).AllOrDefault().CollectRange(500)
```
``` go
// Using Pipe
for item := range From(items).Where("Flag", true).AllOrDefault().Pipe(256) {
    if item.Err.Code != 0 {
        t.Error(item.Err)
    }
}
```

``` go
// Grouping and Piping
groupable := zenq.GroupBy[bool, student](zenq.From(students).AllOrDefault(), "Present")

for item := range groupable.Pipe(0) {
    for k, v := range item.Value {
        // process items
    }
}
// Changed to Pipe() since v1.4.1
```


`Pipe(size)` returns a new type named `CollectStream`.

``` go
type CollectStream[T any] struct {
    Value T
    Err   OpError
}

* If `Err.Code == 0`, it means there is no error.
* `Pipe()` returns data and errors in a single type, which is `CollectStream`.

```
---

## Nested Search Example

Imagine you have a slice of users, and each user has multiple addresses.
Now suppose you want to find all users where a specific city exists in their addresses. zenq makes this kind of nested search much easier to express.

``` go
results, errors := From(UserList).Filter(func(user Users) bool {
    return Any(user.Addr, func(address Address) bool {
        return address.City == "Karaj"
    })
}).AllOrDefault().Collect()

By reading this example, you can get a good sense of how the core functions work together in real use cases.
```
---

## `Any()`

`Any()` accepts:
- A slice.
- A predicate function that returns a boolean.

It returns `true` if at least one item matches the condition, otherwise `false`. This is especially useful for nested queries.

```go
result := Any(items, func(item ComplexObjectToSearch) bool {
    return item.Flag
})
```

---

## `GroupBy()`

`GroupBy()` accepts:
- A queryable.
- A string for the property name.

It groups the data based on the specific key.

``` go
result, err := GroupBy[bool, SysUser](From(users), "Flag").Collect()

result, err2 := GroupBy[uint32, SysUser](From(users).Filter(func(user SysUser) bool {
    return user.Id > 0
}), "AuthorityId").Collect()

```



# Thor Collection Api

A faster, more Go-idiomatic alternative to the default collections API is to use the **Thor** engine to query your data. The Thor engine uses the operator fusion pattern to ensure maximum speed and a single execution unit. like the default collections api, the thor collection api's can help you to filter, validate, group your collections.


import path
``` go
collections "github.com/malikhan-dev/zenq/collections/Thor"
```

### Core Concepts:

**`CollectionCompiledQueryable[T]`**: After each chain of operation, we use this type as a contract (much like `Queryable` in the default collections API).

   
**`AssertCompiledQueryable[T]`**: In our query chains, if we want to assert the result like the `Any()` operator, this is the output type.


**`GroupCompiledQueryable[K, T]`**: After a grouping operation, the returning type is `GroupCompiledQueryable`.


All three types nest `CompiledQueryable[T]` inside them. `CompiledQueryable` represents the result of the operation in the `Items` property and the list of operators.


``` go
type CompiledQueryable[T any] struct {
    Operators []zenqOperator[T]
    Items     *[]T
}
```

Thor Engine APIs are as follows:

**`From[T any]`**: Accepts a slice of `[]T` and returns a `*CollectionCompiledQueryable[T]` to initiate a query chain.

  
**`Where[T any]`**: Accepts a function `func(T) bool` as an argument, filters the collection, and returns a `*CollectionCompiledQueryable[T]`.

  
**`Collect()`**: Collects the result and returns the `CollectionCompiledQueryable[T]` which holds the data.

  

**Example:**

``` go
result := collections.From(items).Where(func(search ComplexObjectToSearch) bool {
    return search.Name == "Jane" && search.Flag == false
}).Collect()

result2 := collections.From(result).Any(func(search ComplexObjectToSearch) bool {
    return (search.Name != "Jane") || (search.Flag != false)
}).Assert()

if result2 {
    t.Error("result should be false")
}
```

**`Group` and `Collect`**: The `Group` function expects a `CompiledQueryable[T]` as an argument and a Key Selector function. For collecting the result of a group, we can use the `collections.Collect()` function.

A grouping example: filtering users whose age is greater than 20 and grouping them by their presence:

``` go
res := 
    collections.Group[bool, ComplexObjectToSearch](
        collections.From(items).Where(func(search ComplexObjectToSearch) bool {
            return search.Age > 20
        }),
        func(item ComplexObjectToSearch) bool {
            return item.Flag
        },
    ).Collect()


fmt.Println(res.Items[false][1])
fmt.Println(res.Items[true][1])

```


**`Assert()`**: Asserts the collection on a given criteria.

``` go
result2 := collections.From(result).Any(func(search ComplexObjectToSearch) bool {
    return (search.Name != "Jane") || (search.Flag != false)
}).Assert()

```



## Nested Search Example (Thor Api)

Imagine you have a slice of users, and each user has multiple addresses.
Now suppose you want to find all users where a specific city exists in their addresses. how can we make such query using Thor API?

``` go

	res :=
		collections.From(UserList).Where(func(user Users) bool {

			return collections.From(user.Addr).Any(func(address Address) bool {
				return address.City == "Karaj"
			}).Assert()

		}).Collect()

	fmt.Println(res)

```
---


# zenq Stream API

When dealing with large datasets, it is not always recommended to collect everything into memory using the traditional `Queryable` execution model.

zenq provides a Stream API that allows data to be processed incrementally as it flows through a pipeline. imagine you want a way to process a large csv file record by record... or a MySql Database. you need to open a cursor of your database and start processing the rows. as mentioned earlier, its not a good idea to collect all the data in memory. you can use zenq streams which is compatible with numerous data-sources to achieve your goal. filter the stream of your data, cause a delay to the streams and process your data with ease.


import path
``` go

streams  "github.com/malikhan-dev/zenq/streams"

```

Currently there are 5 adapters available to initiate a stream:



## FromData

Creates a stream from in-memory data.

**Args:**
1. A context to manage cancellation.
2. A slice of objects.
   

## FromChannel

Creates a stream from an existing Go channel.

**Args:**
1. A context to manage cancellation.
2. A read channel of `T`.


## FromCsv

Creates a stream from a specific csv file. can perform filters on the stream of data.

**Args:**

1. A context to manage cancellation
2. A contracts.CsvStreamConf[T] type that configures how the stream will initiate.

- contracts.CsvStreamConf[T] contains following properties:

``` go
		type StreamConf struct {
			FilePath string
			BufferSize int
			ParseErrorCallback func([]error, int)
			ItemCount int
		}
		type CsvStreamConf[T any] struct {
			Parser        func(row []string) (T, []error)
			StreamHeaders bool
			StreamConf
		}

```
  1- A parser thats responsible to map a csv row to a type. 
  
  2- A flag Represents that headers of csv should be streamed or not. 
  
  3- A FilePath of the csv file
  
  4 - A BufferSize. atleast 128 recommended.
  
  5 - A callback for when the parser cant parse the row and an error occures, other rows will be streamed though, unless user signals the cancelation through the 'context'.
  
  6 - An ItemCount for when we want to fetch a limited number of csv rows. use 0 to fetch them all.



## FromJsonArr

Creates a stream from a specific json file. can perform filters on the stream of data.

**Args:**

1. A context to manage cancellation
2. A contracts.StreamConf type that configures how the stream will initiate.

 ``` go
		type StreamConf struct {
			FilePath string		
			BufferSize int
			ParseErrorCallback func([]error, int)
			ItemCount int **unsupported at the moment**
		}
 ```
 
  1 - A FilePath of the json file
  
  2 - A BufferSize. atleast 128 recommended.
  
  3 - A callback for when the parser cant parse the row and an error occures, other rows will be streamed though, unless user signals the cancelation through the 'context'.
  
  4 - To be supported on the next releases.


## FromMySqlRows

creates a stream or better a cursor from the rows of a MySql database. first we need to prepare for connecting to the database. in the example below we created a new db-context and started the connection. the ZenqMySqlDb uses the pooling mechanism of the golang database package. so its compatible with concurrency and works with the standards of golang.

``` go

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	db := ZenqMySqlDb{}

	if conn, err := db.NewConnection(constr); err != nil {

		t.Fatal(err)

	}

```

after the connection is initiated, its time to use FromMySqlRows to start the stream. it needs the following arguments.

1 - a cancelation context

2 - a connection to the database. (which we created before)

3 - the query. in mysql we use ? to repressent an argument in the querystring. using parameters is very important and can prevent sql-injection attacks

4 - a mapper function to convert each rows of the cursor to the model defined at application. func(rows *sql.Rows) (UserModel, error)

5 - a variadic argument of any. as the query parameters.

here is how to initiate a stream.

``` go

	defer conn.Close()

		id := 0
		stream :=
			FromMySqlRows[UserModel](ctx, conn,
				"select * from Test.users where id>?", func(rows *sql.Rows) (UserModel, error) {
					var id, age int
					var name string
					var err error

					err = rows.Scan(&id, &name, &age)
					model := UserModel{
						UserId:   id,
						Age:      age,
						UserName: name,
					}
					return model, err
				}, id)
```

when the stream initiated. you can use all the pipelines available for other data-sources such as csvs, json, channels and etc... 


  
# Stream Pipelines

Once a stream is created, it can be processed using different pipeline stages.


## FilterStream

Works similarly to `Where()` or `Filter()`, but operates on streamed data.

**Args:**
1. A function to filter the stream of data (`predicate func(T) bool`).


## Throttle

Adds a delay between streamed items.

**Args:**
1. duration time.Duration.

**Important:**
- Use e.g., `100 * time.Millisecond`.
- Use `0` for no delay.

## MapStream

Transforms streamed data into another type.

**Args:**
1. A context to manage cancellation.
2. A read channel of `T`.
3. A mapping function that maps `T` to another type `M`.

It returns a channel of `M`.

---


Streams respect `context.Context` cancellation to:
- Prevent goroutine leaks.
- Support early termination.
- Properly manage pipeline lifecycle.


# Initiated Stream

it is strongly recommended that when initiating a stream from an asynch source, check that the stream is actually possible. a go idiomatic stream initiation can be something like:

``` go


	if stream := FromJsonArr[User](ctx, jsonStreamConfig.StreamConf); stream.Initiated {

		data := stream.FilterStream(func(c User) bool {
			return c.ID > 0
		})

		for v := range data.Channel {
			time.Sleep(time.Millisecond * 10)
			fmt.Println(" value: ", v)
		}
	}

```

# Example Of Streams

Process a Stream From Data

``` go

ctx, cancel := context.WithCancel(context.Background())
defer cancel()


for v := range FromData[ComplexObjectToSearch](ctx, items).FilterStream(func(search ComplexObjectToSearch) bool {
	return search.Id > 2
}).Throttle(0).TakeAll() {

}


```

process stream from a channel

``` go


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


```


## A Real‑World Example of Querying CSV Files

imagine we have a csv file with the following structure. the first 3 rows have wrong values for Index, cause it should be an int, like other rows. 

our goal is to read the entire csv file, and then have a groupped object based on the index field (a map[k][T]).

we also need to filter the rows that their index field is greater than 60. we want to collect streams then use the thor engine to group the objects. if any other errors occures besides those 3 first rows, the operation must be stopped.





``` csv
	Index,CustomerId,FirstName,LastName,Company,City,Country,Phone1,Phone2,Email,SubscriptionDate,Website
    C681dDd0cc422f7,C681dDd0cc422f7,Kelli,Hardy,Petty Ltd,Huangfort,Sao Tome and Principe,020.324.2191x2022,424-157-8216,kristopher62@oliver.com,2020-12-20,http://www.kidd.com/,
    C681dDd0cc422f7,C681dDd0cc422f7,Kelli,Hardy,Petty Ltd,Huangfort,Sao Tome and Principe,020.324.2191x2022,424-157-8216,kristopher62@oliver.com,2020-12-20,http://www.kidd.com/,
    C681dDd0cc422f7,a940cE42e035F28,Lynn,Pham,"Brennan, Camacho and Tapia",East Pennyshire,Portugal,846.468.6834x611,001-248-691-0006,mpham@rios-guzman.com,2020-08-21,https://www.murphy.com/,
    60,9Cf5E6AFE0aeBfd,Shelley,Harris,"Prince, Malone and Pugh",Port Jasminborough,Togo,423.098.0315x8373,+1-386-458-8944x15194,zachary96@mitchell-bryant.org,2020-12-10,https://www.ryan.com/,
    65,aEcbe5365BbC67D,Eddie,Jimenez,Caldwell Group,West Kristine,Ethiopia,+1-235-657-1073x6306,(026)401-7353x2417,kristiwhitney@bernard.com,2022-03-24,http://cherry.com/,
    65,FCBdfCEAe20A8Dc,Chloe,Hutchinson,Simon LLC,South Julia,Netherlands,981-544-9452,+1-288-552-4666x060,leah85@sutton-terrell.com,2022-05-15,https://mitchell.info/,
```

Here's How:


first we define the data type


``` go
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

```

then its time to configure the streaming options. simply create an instance of CsvStreamConf[Customer]. its important to pass along sufficient amount for the buffer size. since we expect the three first rows are corrupted then we configure our Error Callback in a way that if any other errors occures, we signal the stream to stop.


```

    ctx, cancel := context.WithCancel(context.Background())

    defer cancel()

    var CsvStreamConfig contracts.CsvStreamConf[customer]  //define the stream config

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "customers-100.csv"

	CsvStreamConfig.BufferSize = 256

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 4 {  // we expect that the first 3 rows have problems and if we have error on other records we want to cancel, please note that headernames is the 1 row
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


```


then its time to design our pipelines. we plugin the csv adapter (FromCsv) then we call the filter pipeline so that we can have all the customers that their index field is bigger than 60. then we call TakeAll(), the result is now compatible for the thor engine to start grouping by passing the data to group function. and finally we collect the grouping result. 


```

data := FromCsv(ctx, CsvStreamConfig).FilterStream(func(c customer) bool {
		return c.Index > 60
	}).TakeAll()


GroupCollection := TCollection.Group[int, customer](TCollection.From(data), func(t customer) int {
		return t.Index
	}).Collect()


for k, v := range GroupCollection.Items {
		fmt.Println(k)
		fmt.Println(v)
	}

```


all together this is the final statements


``` go

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

	var CsvStreamConfig contracts.CsvStreamConf[customer]  //define the stream config

	CsvStreamConfig.StreamHeaders = false

	CsvStreamConfig.FilePath = "customers-100.csv"

	CsvStreamConfig.BufferSize = 256

	CsvStreamConfig.ParseErrorCallback = func(err error, i int) {

		fmt.Println(err, " at", i)

		if i > 3 {    // we expect that the first 3 rows have problems and if we have error on other records we want to cancel
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


GroupCollection := TCollection.Group[int, customer](TCollection.From(data), func(t customer) int {
		return t.Index
	}).Collect()


for k, v := range GroupCollection.Items {
		fmt.Println(k)
		fmt.Println(v)
	}


```


## A Real‑World Example of Querying JSON Files
imagine we want to read a json file and stream its data in real-time. we dont want to wait for all the rows of our json array to be read. and its required that we skip any errors that might happens at the first row. 


``` go
[
  {
    "id": 1,
    "username": "user_1",
    "email": "user_1@example.com",
    "is_active": true,
    "created_at": "2026-05-23T13:15:06.534680"
  },
  {
    "id": 2,
    "username": "user_2",
    "email": "user_2@example.com",
    "is_active": false,
    "created_at": "2026-05-22T13:15:06.534750"
  },
  {
    "id": 3,
    "username": "user_3",
    "email": "user_3@example.com",
    "is_active": true,
    "created_at": "2026-05-21T13:15:06.534758"
  },
  {
    "id": 4,
    "username": "user_4",
    "email": "user_4@example.com",
    "is_active": true,
    "created_at": "2026-05-20T13:15:06.534763"
  }
]
```


``` go

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

		if i > 1 {
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

```


## A Real‑World Example of MySql Streams

imagine we have a large number of users. we want to start a stream and process this users one by one. or better we need a cursor that loops through all of these users. then we want to call a web service and determine wether the current user is a valid user or not. its obvious that we cant just fetch all the records inside the memory and process them then update the database. its not very performance-wise to just read users from the database one by one and for each-one we connect to the database then disconnect and connect for the next users. we can use zenq-streams-api to initiate a stream, using a single db-connection and process the rows one by one. this way we consumed alot less memory and avoided round-trip connections to the database.


``` go

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	db := ZenqMySqlDb{}

	if conn, err := db.NewConnection(constr); err != nil {
		t.Fatal(err)
	} else {

		defer conn.Close()

		id := 0
		stream :=
			FromMySqlRows[UserModel](ctx, conn,
				"select * from Test.users where id>?", func(rows *sql.Rows) (UserModel, error) {
					var id, age int
					var name string
					var err error
					var active bool

					err = rows.Scan(&id, &name, &age,&active)
					model := UserModel{
						UserId:   id,
						Age:      age,
						UserName: name,
						Active: active
					}
					return model, err
				}, id)

		if stream.Initiated {
			for v := range stream.FilterStream(func(model UserModel) bool {
				return model.Age > 25
			}).Throttle(time.Millisecond * 5000).Channel {

				/// business logic

				business_logic_satisfied := true

				if business_logic_satisfied {

					result := Exec(conn, "update Test.users set Active = ? where Id =?", 1, v.UserId)
					if result.Err != nil {
						t.Error(result.Err)
					} else {
						fmt.Println(v, " - updated. ", result.RowsAffected)
					}
				}

			}
		} else {
			fmt.Println("stream not initiated")
		}

```


# Compiled Streams
the main difference between streams and compiled streams is that the compiled streams starts the streaming from a single execution unit. while the streams pass around the data after each pipelines. in the following example we initiate a compilable stream using the method CompileFromQueryable, which accepts a slice, then we used filter pipeline to filter it, after that we called CompileStream which is our execution unit, we can remove the throttle pipeline if we dont need any delays. please note that this section is an experimental part of the project.

``` go

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for i := range Throttle(ctx, CompileStream(ctx, Filter(CompileFromQueryable(items), func(student ComplexObjectToSearch) bool {
		return !student.Flag
	})), time.Duration(250*time.Millisecond)) {
		fmt.Println(i)
	}

```




## benchmark

in a slice of 50,000,000 users it took less than 2 seconds just to filter them and around 4 seconds to filter then group the items. to achieve this result we used thor collections api, not the default api collections.






## Project Status

zenq is actively evolving, and more operators, examples, and documentation are on the way.

If you find it useful, feel free to star the repository (it motivates us) and follow future updates!


## Third-Party Software References:

Third‑Party Software Notice: This package includes/uses the third‑party MySQL driver github.com/go-sql-driver/mysql.
Copyright © The github.com/go-sql-driver/mysql authors.

Third‑Party Software Notice: This package includes/uses the third‑party Postgres driver github.com/lib/pq
Copyright © The github.com/lib/pq authors.

License applies as stated in those repository.
