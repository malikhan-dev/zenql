<p>
<img width="20" height="20" src="https://github.com/user-attachments/assets/095647c1-b3dd-4d5a-95ea-bccb3e610585"/>
<img src="https://img.shields.io/badge/Go-1.25+-00ADD8"/>
<img src="https://img.shields.io/badge/tests-passing-brightgreen"/>
<img src="https://img.shields.io/badge/version-2.0.0-green"/>
<img src="https://visitor-badge.laobi.icu/badge?page_id=malikhan-dev.zenq"/>
<a href="https://pkg.go.dev/github.com/malikhan-dev/zenql"><img src="https://pkg.go.dev/badge/github.com/malikhan-dev/zenql.svg" alt="Go Reference"/></a>
<img src="https://goreportcard.com/badge/github.com/malikhan-dev/zenql"/>
<img src="https://img.shields.io/badge/license-MIT-blue"/>
<img src="https://img.shields.io/badge/clones-2.2k%2B-brightgreen?logo=github"/>
<img width="20" height="20" src="https://github.com/user-attachments/assets/095647c1-b3dd-4d5a-95ea-bccb3e610585"/>
</p>

###  ZenQL-V2


**Expressive, Loosely Coupled and Type-Safe Query Engine for Go. Inspired By LINQ.**

<div align="center">
   <img width="600" height="250" alt="zql-demo-big" src="https://github.com/user-attachments/assets/64cbe792-2189-4b90-ab4d-a90a2d0feef6" />
   <img width="600" height="250" alt="demo-best" src="https://github.com/user-attachments/assets/359069c6-c98d-4ad8-a214-10ea11abeb58" />

</div>

<div align="center">
 <b>Trusted By 2.2K+ Cloners!</b>

</div>

</div>


### Support Us
ZenQL is built and maintained with passion. If you find it useful, dropping a ⭐ on the repo is the simplest way to show your support — and it genuinely matters.


### ⚡ Quick Start
See how ZenQL simplifies data querying:

```go

	if cursor := FromSqlRows[Users](ctx, conn, "select * from users where id > ?", id); cursor.Initiated {
		for v := range cursor.FilterStream(func(u Users) bool {
			return u.ID > 0
		}).Throttle(1 * time.Second).Channel {
			// Process business logic
		}
	}

```

---



### Why ZenQL?

ZenQL brings the power of polymorphic querying to the Go ecosystem, adhering to idiomatic practices while solving performance bottlenecks. its your integrated language when dealing with data in different places and formats. wether its json,csv,mysql,postgres or you simply want to work with in memory data fluently and comfortably. its really fast... 

in some scenario and test cases (the 50,000,000 in-memory records filtering and validation) it is significantly faster than C# LINQ. and its 100% type safe with zero reflections. and for collection processing it performs compiled queries in a single execution unit. 

ZenQL consumes memory very carefully, efficiently and in a controlled and predictable way (thanks to smart memory management mechanism). so its not just a query language thats fluent and readable, its an efficient query language. 


*   ⚡ Thor Engine (Collections Api): Fused execution pattern for maximum performance.
*   🌊 Async Streaming: Process large datasets incrementally without memory spikes.
*   🛡️ Smart Memory Management: GC-friendly allocations with automatic capacity tuning.
*   🔗 Unified Streaming API: Consistent syntax for Slices, Channels, CSV, JSON, and RDBMS (MySQL/Postgres).

---


<div align="center">

| Category | Feature | Status |
| :--- | :--- | :---: |
| **Data Sources** | In-Memory Slices | ✅ |
| | Channels | ✅ |
| | JSON Files | ✅ |
| | CSV Files | ✅ |
| | MySQL | ✅ |
| | PostgreSQL | ✅ |
| **Capabilities** | Fluent Querying/Filtering/Grouping/Sorting/Projection/Pagination | ✅ |
| | Async Streaming | ✅ |
| | Context Cancellation | ✅ |
| | Operation Fusion (Thor) | ✅ |

</div>


### 📊 Performance 
ZenQL is built with speed in mind. Our Thor engine minimizes overhead to keep your application blazing fast.

**Benchmark: Filtering 50,000,000 records via collections api**

| Metric | Result |
| :--- | :--- |
| **Ops** | 13.6 ms/op |
| **Allocations** | 0 allocs/op |
| **Memory** | 22.7 MB/op |

---



### License (MIT)

This library was written and designed by Mohammadreza Malikhan. The source code is free to use with proper attribution. This project is licensed under the MIT License (see the `LICENSE` file for details). also other contributors involved with the project. visit contributors section for more information

### Intro
ZenQL is a fluent Domain-Specific Language (DSL) for Go, designed to help you filter, search, validate, process, and stream data with readability and ease. Inspired by LINQ in C# and Java Streams, ZenQL brings the power of polymorphic querying to the Go ecosystem while adhering to idiomatic Go practices.


ZenQL is built as a modular library, currently featuring three primary components:

Collections (Thor): Designed for high-performance in-memory data processing. It leverages the operation fusion pattern to run entire query chains within a single execution unit, minimizing overhead.


Streams: Built for asynchronous data handling. By utilizing core Go concepts like channels and goroutines, Streams allow you to process data asynchronously while fully respecting Go’s context and cancellation patterns.

Databases: Enables seamless communication with async data sources, such as MySQL databases or postgres, using the same fluent syntax.

At the moment, ZenQL supports a wide range of data sources, including in-memory slices, channels, CSV/JSON files, and MySQL databases—with more connectors on the way.


### Installation and Migration

ZenQL V2 is a modular library. modules and its dependencies are reviewed and refactored. it contains four modules.

1 - contracts: contracts and abstractions of ZenQL

2 - collections/Thor: the collections processor. (depends on contracts)

3 - streams: for streaming data. (depends on contracts)

4 - databases: our mini-orm. (depends on contracts, streams, external db drivers described at the end of documents)

we wanted you to have a choice to use any part of ZenQL you want. maybe all of it or some of it.

the migrations process isnt really that hard:

1 - go clean -modcache.

2 - remove the dependencies manually from go.mod (if thats needed).

3 - start getting the packages.

``` go

		go get github.com/malikhan-dev/zenql/collections/Thor/v2@v2.0.0
		
		go get github.com/malikhan-dev/zenql/contracts/v2@v2.0.0
		
		go get github.com/malikhan-dev/zenql/streams/v2@v2.0.2
		
		go get github.com/malikhan-dev/zenql/databases/v2@v2.0.3

```

4 - changing the import paths.



### Changelog 

### v2.0.0

ZenQL is Modular now. and each modules installs in seperate.


``` go

		go get github.com/malikhan-dev/zenql/collections/Thor/v2@v2.0.0
		
		go get github.com/malikhan-dev/zenql/contracts/v2@v2.0.0
		
		go get github.com/malikhan-dev/zenql/streams/v2@v2.0.2
		
		go get github.com/malikhan-dev/zenql/databases/v2@v2.0.3

```

this release contains the following modules:

1 - zenql/collections/Thor/v2@v2.0.0 .

2 - zenql/contracts/v2@v2.0.0 .

3 - zenql/streams/v2@v2.0.2 .

4 - zenql/databases/v2@v2.0.3 .


## Module1: Thor Collection Api

earlier we developed a new module to process the collections named as default collections api (which is deprecated and removed in v1.8.0). later on a new collections query engine developed named Thor. A faster, more Go-idiomatic alternative to the default collections API. The Thor engine uses the operator fusion pattern to ensure maximum speed and a single execution unit. like the default collections api, the thor collection api's can help you to filter, validate and group your collections.


import path
``` go
collections "github.com/malikhan-dev/zenql/collections/Thor/v2"
```

### Core Concepts:

**`CollectionCompiledQueryable[T]`**: After each chain of operation, we use this type as a contract (much like `Queryable` in the default collections API).

   
**`AssertCompiledQueryable[T]`**: In our query chains, if we want to assert the result like the `Any()` operator, this is the output type.


**`GroupCompiledQueryable[K, T]`**: After a grouping operation, the returning type is `GroupCompiledQueryable`.


All three types nest `CompiledQueryable[T]` inside them. `CompiledQueryable` represents the result of the operation in the `Items` property and the list of operators.


``` go
type CompiledQueryable[T any] struct {
    Operators []zenqlOperator[T]
    Items     *[]T
}
```



Thor Engine APIs are as follows:

### From[T any]:
Accepts a pointer to an slice of `[]T` and returns a `*CollectionCompiledQueryable[T]` to initiate a query chain.

  
### Where[T any]: 
Accepts a function `func(T) bool` as an argument, filters the collection, and returns a `*CollectionCompiledQueryable[T]`.

  
### Collect(): 
Collects the result and returns the `CollectionCompiledQueryable[T]` which holds the data.

  

**Example:**

``` go
result: = collections.From( & items).Where(func(search ComplexObjectToSearch) bool {
    return search.Name == "Jane" && search.Flag == false
}).Collect()

result2: = collections.From( & result).Any(func(search ComplexObjectToSearch) bool {
    return (search.Name != "Jane") || (search.Flag != false)
}).Assert()

if result2 {
    t.Error("result should be false")
}
```

### Group(): 
The `Group` function expects a `CompiledQueryable[T]` as an argument and a Key Selector function. For collecting the result of a group, we can use the `collections.Collect()` function.

A grouping example: filtering users whose age is greater than 20 and grouping them by their presence:

``` go
res: =
    collections.Group[bool, ComplexObjectToSearch](

        collections.From( & items).Where(func(search ComplexObjectToSearch) bool {
            return search.Age > 20
        }),

        func(item ComplexObjectToSearch) bool {
            return item.Flag
        },

    ).Collect()


fmt.Println(res.Items[false][1])
fmt.Println(res.Items[true][1])
```


### Assert():
Asserts the collection on a given criteria.

``` go
result2 := collections.From(&result).Any(func(search ComplexObjectToSearch) bool {
    return (search.Name != "Jane") || (search.Flag != false)
}).Assert()

```


### CollectSorted():

Sorts the thor engine collections. arguments are:

- a less function or (less func(T, T) bool): to determine the way of comparing two items of the same kind
- desc bool: determinse the sort direction, ascending or descending

``` go

result: = From(&personList).Where(func(person Person) bool {

    return person.Active == true

}).CollectSorted(func(person Person, person2 Person) bool {

    return person.Identifier < person2.Identifier

}, true)

```

### Project():

Projects (maps) the operation result to another type using a user defined function. use it instead of Collect() to Collect and Project the result at the same time! (it will be compiled).

this is a generic function that requires types of source and destinations. arguments are: 

1 - a CollectionCompiledQueryable[T]  

2 - a user defined mapper



``` go

MapPersonToSysUser: = func(person Person) SysUser {

    user: = SysUser {
        FName: person.Name,
        LName: person.LastName,
        Id: person.Identifier,
        Email: person.Mail,
        Enabled: person.Active,
    }

        if len(person.Address) > 0 {

        user.Address = fmt.Sprintf("%s mapped", person.Address[0].City)
    }

    return user
}

newUsers: = Project[Person, SysUser](

    From(&personList).Where(func(person Person) bool {
        return person.Identifier > 0
    }),

    MapPersonToSysUser,
)

```


### Take And Skip
You can use Take(count) and Skip(count) functions to skip some rows and take some other rows. these functionalities are specifically useful for pagination too. Take and Skip support early exit strategies and supports operator fusion pattern discussed earlier.

``` go

result := Project[Employee, InternalEmp](
	From(&employees).Skip(2).Take(1), 
	  func(e Employee) InternalEmp {
              return InternalEmp{
                  FullName: e.Name,
                  Dep:      e.Department,
              }
	   },
	)
```

```go
result := From(&employees).Where(func(employee Employee) bool {
	return employee.Department == "IT"
        }).Take(1).Skip(1).Collect()
```

``` go
result := From(&personList).Skip(0).Take(1).Where(func(person Person) bool {
     return person.Active == false
 }).CollectSorted(func(person Person, person2 Person) bool {
     return person.Identifier < person2.Identifier
}, true)
```

``` go
GroupResult := Group[bool, Student](From(&students).Skip(2).Take(2).Where(func(student Student) bool {
	return student.Age > 0
}), func(student Student) bool {
	return student.Pressent
}).Collect()
```
### Nested Search Example (Thor Api)

Imagine you have a slice of users, and each user has multiple addresses.
Now suppose you want to find all users where a specific city exists in their addresses. how can we make such query using Thor API?

``` go
res: = collections.From( & UserList).Where(func(user Users) bool {

    return collections.From(user.Addr).Any(func(address Address) bool {

        return address.City == "Karaj"

    }).Assert()

}).Collect()

fmt.Println(res)

```
---


### Smart Memory Management

We are happy to announce that ZenQL now includes a new smart memory management mechanism for its internal operations.

This feature is designed to help ZenQL remain performant, stable, and production-ready by providing a more predictable and controlled allocation strategy.

Internal allocations are now handled based on several factors, including:

1. Available heap memory at runtime
2. Estimated memory usage before allocation
3. A user-defined maximum allocation guard

This new approach helps reduce GC pressure, lowers unnecessary memory consumption, and provides safer behavior in production environments, especially when working with large datasets.

You can configure the maximum allocation guard through the `contracts` module:
```go
contracts.SetMaxAllocGuard(200000)
```

This allows ZenQL to use a maximum initial allocation capacity of 200,000 for the underlying array created by make().

By default, the max allocation guard is set to 5,000,000.

Use this feature with caution and tune it based on the memory limits and workload characteristics of your environment.


## Module2: Streams Api

When dealing with large datasets, it is not always recommended to collect everything into memory using the traditional `Queryable` execution model.

zenql provides a Stream API that allows data to be processed incrementally as it flows through a pipeline. imagine you want a way to process a large csv file record by record... or a MySql Database. you need to open a cursor of your database and start processing the rows. as mentioned earlier, its not a good idea to collect all the data in memory. you can use zenql streams which is compatible with numerous data-sources to achieve your goal. filter the stream of your data, cause a delay to the streams and process your data with ease.


import path
``` go

streams  "github.com/malikhan-dev/zenql/streams/v2"

```

Currently there are 5 adapters available to initiate a stream:



### FromData

Creates a stream from in-memory data.

**Args:**
1. A context to manage cancellation.
2. A slice of objects.
   

### FromChannel

Creates a stream from an existing Go channel.

**Args:**
1. A context to manage cancellation.
2. A read channel of `T`.


### FromCsv

Creates a stream from a specific csv file. can perform filters on the stream of data.

**Args:**

1. A context to manage cancellation
2. A contracts.CsvStreamConf[T] type that configures how the stream will initiate.

contracts.CsvStreamConf[T] contains following properties:

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



### FromJsonArr

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


### FromSqlRows

creates a stream or better a cursor from the rows of a sql database (postgresql, mysql supported). first we need to prepare for connecting to the database. in the example below we created a new db-context and started the connection. the context uses the pooling mechanism of the golang database package. so its compatible with concurrency and works with the standards of golang. if you dont want to use the database module, you have to implement the RDBMSFacade interface (in contracts package) to establish a connection to your database and provide it to this function.

import path

``` go
collections  "github.com/malikhan-dev/zenql/databases/v2"
```


``` go

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"


	if conn, err := db.Connect("mysql",constr); err != nil {

		t.Fatal(err)

	}

```

after the connection is initiated, its time to use FromSqlRows to start the stream. it needs the following arguments.

1 - a cancelation context

2 - a connection to the database. (which we created already)

3 - the query. in mysql we use ? to repressent an argument in the querystring. using parameters is very important and can prevent sql-injection attacks

4 - a variadic argument of any. as the query parameters.

model to map must tagged with 'zql'

here is how to initiate a stream.

``` go
defer conn.Close()
id: = 0
stream: = FromSqlRows[UserModel](ctx, conn, "select * from Test.users where id>?", id)

```

when the stream initiated. you can use all the pipelines available for other data-sources such as csvs, json, channels and etc... 


  
### Stream Pipelines

Once a stream is created, it can be processed using different pipeline stages.


### FilterStream

Works similarly to `Where()` or `Filter()`, but operates on streamed data.

**Args:**
1. A function to filter the stream of data (`predicate func(T) bool`).


### Throttle

Adds a delay between streamed items.

**Args:**
1. duration time.Duration.

**Important:**
- Use e.g., `100 * time.Millisecond`.
- Use `0` for no delay.

### MapStream

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


### Initiated Stream

it is strongly recommended that when initiating a stream from an asynch source, check that the stream is actually possible. a go idiomatic stream initiation can be something like:

``` go
if stream: = FromJsonArr[User](ctx, jsonStreamConfig.StreamConf);stream.Initiated {

    data: = stream.FilterStream(func(c User) bool {

        return c.ID > 0

    })

    for v: = range data.Channel {
        time.Sleep(time.Millisecond * 10)
        fmt.Println(" value: ", v)
    }
}

```

### Example Of Streams

Process a Stream From Data

``` go
ctx, cancel: = context.WithCancel(context.Background())

defer cancel()


for v: = range FromData[ComplexObjectToSearch](ctx, items).FilterStream(func(search ComplexObjectToSearch) bool {

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

    }
```


### A Real‑World Example of Querying CSV Files

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


### A Real‑World Example of Querying JSON Files
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


### A Real‑World Example of MySql Streams

Imagine a scenario with a large user base where you need to process users individually, such as validating each one against an external web service. Loading all records into memory is neither efficient nor scalable. Conversely, repeatedly opening and closing a database connection for every single row creates a significant performance bottleneck.
With the new zenql Streams API, you can initiate a stream using a single database connection to process rows iteratively, just like a cursor. This approach significantly reduces memory consumption and optimizes performance by eliminating unnecessary database round-trips.

``` go

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := db.connect("mysql",constr); err != nil {
		t.Fatal(err)
	} else {

		defer conn.Close()

		id := 0
		stream :=
			db.FromSqlRows[UserModel](ctx, conn,
				"select * from Test.users where id>?", id)

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



## Module3: Databases 

Zen-Q supports popular relational database management systems (RDBMS) such as MySql and Postgres. and we use appropriate drivers for these databases mentioned in copyright notice section of the document. a database facade interface created to interact with relational databases.


import path
``` go
collections  "github.com/malikhan-dev/zenql/databases/v2"
```

``` go
type RDBMSFacade interface {
	Close() error
	Ping() error
	GetPool() *sql.DB
	Query(query string, args ...any) (*sql.Rows, error)
	Commit() bool
	Rollback() bool
	Begin() bool
	GetActiveTransaction() *sql.Tx
}
```

we already implemented this interface for mysql and postgres databases. you can use this interface to develop your own drivers for databases or just simply make your facade available to zenql operations.

### commands and queries 

#### Connecting
the module currently supports mysql and postgresql. 

The connect() method accepts a string as database identifier such as 'postgresql' or 'mysql'. returns a dbContext that implements RDBMSFacade.

```
if conn, err := db.Connect("postgres", constrPgsql); err != nil {

		t.Fatal(err)

	}
```

#### Executing a command

the Exec function accepts an RDBMSFacade type, query string and variadic arguments

``` go
postgres sample
res := db.Exec(conn, "DELETE FROM users WHERE Id =$1", 2)
```

``` go
mysql sample
res := db.Exec(conn, "DELETE FROM users WHERE Id =?", 2)
```

returns an object of CommandResult Type

``` go
type Commandesult struct {
	Err          error
	RowsAffected int64
	TimeStamp    time.Time
}
```


#### Performing a query

the Query function accepts an RDBMSFacade type, a query string and variadic arguments. its a generic function that requires a model to map. you can use the 'zql' tag to define the mapping property name.

``` go
    type UserModel struct {
        UserId   int    `zql:"Id"`
        UserName string `zql:"Name"`
        Age      int    `zql:"Age"`
    }
```

``` go

    constr := "root:1245Sa@tcp(127.0.0.1:30306)/Test?parseTime=true&charset=utf8mb4"

	if conn, err := db.Connect("mysql", constr); err != nil {

		t.Fatal(err)

	} else {

		limit := 4

		result, err := db.Query[UserModel](conn, "select * from Test.users  limit ?", limit)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}

```


### transaction control
the RDBMSFacade contains a method named GetActiveTransaction(). when performing a command, the execution pipeline will first try to run the command on the active transaction *Sql.Tx. if there are no transactions initiated the command will be run normally.


- initiate a transaction
  the connection (RDBMSFacade) contains a method named Begin(). this method starts a transaction

- Commiting a transaction
  the connection (RDBMSFacade) contains a method named Commit(). this method Commits the active transaction.

- Rolling back a transaction
  the connection (RDBMSFacade) contains a method named Rollback(). this method reverts all the changes made so far.

here is an example of concepts:

```go
 if conn, err := db.Connect("postgres", constrPgsql); err != nil {
	 
	 t.Fatal(err)
	 
 } else {
	 
	 conn.Begin()
	 
	 res := db.Exec(conn, "DELETE FROM users WHERE Id =$1", 2)
	 
	 if res.Err == nil {
	
		 name := "mohammadreza"
	
		 age := 65
	
		 cmd2 := db.Exec(conn, "INSERT INTO users (name,age) values($1,$2)", name, age)
	
		 if cmd2.Err != nil {
			 conn.Rollback()
		 } else {
			 conn.Commit()
		 }
		 
	 } else {
            conn.Rollback()	 
    }   
	
}
```


### Project Status

zenql is actively evolving, and more operators, examples, and documentation are on the way.

If you find it useful, feel free to star the repository (it motivates us) and follow future updates!


### Third-Party Software References:

Third‑Party Software Notice: database module includes/uses the third‑party MySQL driver github.com/go-sql-driver/mysql.
Copyright © The github.com/go-sql-driver/mysql authors.

Third‑Party Software Notice: database module includes/uses the third‑party Postgres driver github.com/lib/pq
Copyright © The github.com/lib/pq authors.

License applies as stated in those repository.
