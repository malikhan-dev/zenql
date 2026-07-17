<p>
<img width="20" height="20" src="https://github.com/user-attachments/assets/095647c1-b3dd-4d5a-95ea-bccb3e610585"/>
<img src="https://img.shields.io/badge/Go-1.25+-00ADD8"/>
<img src="https://img.shields.io/badge/tests-passing-brightgreen"/>
<img src="https://img.shields.io/badge/version-2.0.4-green"/>
<img src="https://visitor-badge.laobi.icu/badge?page_id=malikhan-dev.zenq"/>
<a href="https://pkg.go.dev/github.com/malikhan-dev/zenql"><img src="https://pkg.go.dev/badge/github.com/malikhan-dev/zenql.svg" alt="Go Reference"/></a>
<img src="https://img.shields.io/badge/license-MIT-blue"/>
<img src="https://img.shields.io/badge/clones-2.2k%2B-brightgreen?logo=github"/>
<img width="20" height="20" src="https://github.com/user-attachments/assets/095647c1-b3dd-4d5a-95ea-bccb3e610585"/>
</p>

###  ZenQL-V2


**Expressive, Loosely Coupled and Type-Safe Query Engine for Go. Inspired By LINQ.**


<div align="center">
	<img width="600" height="250" alt="Demo-2" src="https://github.com/user-attachments/assets/7407c8ee-511c-4738-ab28-91a5b5e0ce68" />
    <img width="600" height="250" alt="Screencastfrom2026-07-1423-54-55-ezgif com-video-to-gif-converter" src="https://github.com/user-attachments/assets/4f3f3272-f220-4c9a-82c9-a8f03f78b6a5" />

</div>


<div align="center">
 <b>Trusted By 2.2K+ Cloners!</b>

</div>



### Support Us
ZenQL is built and maintained with passion. If you find it useful, dropping a ⭐ on the repo is the way to show your support and it genuinely matters.




### Why ZenQL?

ZenQL brings the power of polymorphic querying to the Go ecosystem, adhering to idiomatic practices while solving performance bottlenecks. its your integrated language when dealing with data in different places and formats. wether its json,csv,mysql,postgres or you simply want to work with in memory data fluently and comfortably. its really fast... 

in some scenario and test cases (the 50,000,000 in-memory records filtering and validation) it can be faster than C# LINQ (Not Guaranteed). and its 100% type safe with zero reflections. for collection processing it performs compiled queries in a single execution unit (operator fusion pattern). 

ZenQL consumes memory very carefully, efficiently and in a controlled and predictable way (thanks to smart memory management mechanism). so its not just a query language thats fluent and readable, its an efficient query language. 


*   ⚡ Thor Engine & Sifu: Fused execution pattern for maximum performance, supporting Zen-QL's Expression Builder (Sifu).
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

warning: benchmarks depend on the environment and the results below are the best results collected from a series of repeatable tests.

| Metric | Result |
| :--- | :--- |
| **Ops** | 13.6 ms/op |
| **Allocations** | 0 allocs/op |
| **Memory** | 22.7 MB/op |



---



### License (MIT)

This library was written and designed by Mohammadreza Malikhan. The source code is free to use with proper attribution. This project is licensed under the MIT License (see the `LICENSE` file for details). also other contributors involved with the project. visit contributors section for more information

### Intro
ZenQL is an internal query language for Go, designed to help you filter, search, validate, process, and stream data with readability and ease. Inspired by LINQ in C# and Java Streams, ZenQL brings the power of polymorphic querying to the Go ecosystem while adhering to idiomatic Go practices.


ZenQL is built as a modular library, currently featuring following components:

Collections (Thor): Designed for high-performance in-memory data processing. It leverages the operation fusion pattern to run entire query chains within a single execution unit, minimizing overhead.

Expression Builder (Sifu): you can generate expressions that compiles to functions which is acceptable to Collections Processor Apis as an argument. 

Streams: Built for asynchronous data handling. By utilizing core Go concepts like channels and goroutines, Streams allow you to process data asynchronously while fully respecting Go’s context and cancellation patterns.

Databases: Enables seamless communication with async data sources, such as MySQL databases or postgres, using the same fluent syntax.

At the moment, ZenQL supports a wide range of data sources, including in-memory slices, channels, CSV/JSON files, and MySQL databases—with more connectors on the way.


###  Migration To V2

ZenQL V2 is a modular library. modules and its dependencies are reviewed and refactored. it contains four modules.

1 - contracts: contracts and abstractions of ZenQL

2 - collections/Thor: the collection processor. (depends on contracts)

3 - streams: for streaming data. (depends on contracts)

4 - databases: our mini-orm. (depends on contracts, streams, external db drivers described at the end of documents)

we wanted you to have a choice to use any part of ZenQL you want. maybe all of it or some of it.

the migrations process is not really that hard:

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


## Changelog 

this release works with following modules

1 - zenql/collections/Thor/v2@v2.0.4

2 - zenql/contracts/v2@v2.0.3

3 - zenql/streams/v2@v2.0.4

4 - zenql/databases/v2@v2.0.5

5 - zenql/expressions/Sifu@v1.0.3

### v2.0.4

 
Introducing Sifu Expressions Builder. 

Warning: Sifu Expressions prior to v1.0.3 is unstable for production use. upgrade to v1.0.3.

Query your in-memory data effortlessly with Sifu Expressions and the Thor Collections API — zero runtime crashes, truly optimized performance.



``` go
		
    go get github.com/malikhan-dev/zenql/collections/Thor/v2@v2.0.4
    
    go get github.com/malikhan-dev/zenql/contracts/v2@v2.0.3

    go get github.com/malikhan-dev/zenql/expressions/Sifu/@v1.0.3
    
    go get github.com/malikhan-dev/zenql/streams/v2@v2.0.4
    
    go get github.com/malikhan-dev/zenql/databases/v2@v2.0.5
    
```





### Thor Collection Api

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


### Update():
updates the rows that passed through Where(). recommended to use before calling the Collect()
  

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

result: = From( & personList).Where(func(person Person) bool {

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
You can use Take(count) and Skip(count) functions to skip some rows and take some other rows. these functionalities are specifically useful for pagination too. Take and Skip support early exit strategies and supports operator fusion pattern discussed earlier. these functions accept int32.

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


Thor Collection Api comes with a set of useful utilities to traverse trees! (since v2.0.1)


### FindParentNode
When dealing with a tree and wee need to find the parent node of a queried item, we can use this function. this function is compiled too.

args:

1 - a function to locate the starting node. func(T) bool


2 - a function that represents how the nodes link together. func(T,T) bool


this example returns the parent node of an item with the id of 9.

``` go
	targetNode := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Tehran"
		}).Assert()

	}).FindParentNode(func(user User) bool {

		return user.Id == 9

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id
	})

```

### FindRootNode
How about finding the root node of a tree? instead of finding the parent? we can use FindRootNode()

args:


1 - a function to locate the starting node. func(T) bool

2 - a function that represents how the nodes link together. func(T,T) bool

3 - a function that determines the lesser item. (a less function). func(T,T) bool

this example returns the root node of an item with the id of 9. it traverses the tree so there is no more parent left!

``` go
	targetNode1 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Karaj" || address.City == "Tehran"
		}).Assert()

	}).FindRootNode(func(user User) bool {

		return user.Id == 7

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	}, func(user User, user2 User) bool {

		return user.Id < user2.Id

	})


```


### TraverseRootNode
Works exactly like FindRootNode(). the only difference is this functions outputs the traversed path in a go channel so that you can see all the nodes that has been traversed. 


1 - a function to locate the starting node. func(T) bool

2 - a function that represents how the nodes link together. func(T,T) bool

3 - a function that determines the lesser item. (a less function). func(T,T) bool

4 - a context to prevent goroutine leaks and traverse stoppage

``` go

	ctx, cancel := context.WithCancel(context.Background())


	targetNode5 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Karaj" || address.City == "Tehran"
		}).Assert()

	}).TraverseRootNode(func(user User) bool {

		return user.Id == 7

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	}, func(user User, user2 User) bool {

		return user.Id < user2.Id

	}, ctx)

	for v := range targetNode5 {
	
	}
```


### CollectUpdated (Will Be Deprecated Soon)
just like the Collect() function it Collects all the item, but it updates all the items match the Where() criteria too. with no refrence attached you will get a new updated slice.

args:
1 - an update function. func(T) T

the following example with all the functions are compiled all together! 

``` go
    result := From(&CityList).Where(func(search city) bool {

		return search.Active

	}).Skip(1).Take(1).CollectUpdated(func(search city) city {

		search.Name += " is active"

		return search

	})
```


### Update
you can update the rows that matches the criteria by calling update method. use this method before calling collect() to be sure it rans on compiled mode. its a proper alternative to CollectUpdated() and runs in compiled mode.

args:

1 - an update function. func(T) T

``` go

	result := From(&CityList).Where(func(search city) bool {

		return !search.Active

	}).Skip(1).Take(1).Update(func(search city) city {
			search.Name += " Deactivated"
			return search
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




### Expression Builder (Sifu)

Warning: Sifu Expressions prior to v1.0.3 is unstable for production use. upgrade to v1.0.3.

Sifu is an expression builder for Zen-QL currently capable of generating native functions for all api's of the thor. the generated codes are heavily optimized and we conducted extensive tests to be sure it doesnt break runtime in case user made a mistake using them. for example setting invalid property name or assigning invalid data types. 



### Expr[T]

The expressions are based on a specific type. and we need to define an expression like below

``` go

	expr := Sifu.Expr[ComplexObjectToSearch]()

```


### Prop()

inorder to select a specific property we use Prop Method. if invalid property name given as argument, it will not cause any runtime break. simply the query returns nothing when its executed. the prop name located based on the expression type described above.

args:

1 - the property name as string

``` go

expr.Prop("Flag")

```

### True and False

after we select the prop its time to evaluate the expression. True or False analyze wether the selected property has the boolean value of true or false.

``` go

expr := Sifu.Expr[ComplexObjectToSearch]()

expr.Prop("Flag").True()

expr.Prop("Flag").False()

```

### StrEq, StrEqNot and StrIn

checks the string values are equal or not. StrIn checks wether a string can be found in an array of strings

``` go

expr := Sifu.Expr[ComplexObjectToSearch]()

expr.Prop("Name").StrEq("Jane")

expr.Prop("Name").StrEqNot("Jane")

expr.Prop("City").StrIn([]string{"Tehran", "Karaj"}),

```

### converting expressions to a generated function (Predicate)

expressions needs to be generated and compiled by calling Predicate() Method. it will generate the expected function.

for example the following code generates a func(T) bool, that can be passed as an argument for Thor Api

``` go

expr := Sifu.Expr[ComplexObjectToSearch]()

predFunc := expr.Prop("Name").StrEq("Jane").Predicate()


```


### And

Links two expressions of same type and performs And, for example

``` go


	expr := Sifu.Expr[ComplexObjectToSearch]()

	query1 := expr.Prop("Name").StrEq("Jane").And(expr.Prop("Flag").True())

    result := collections.From(&items).WhereEx(query1).Collect()

    result2 := collections.From(&items).Where(query1.Predicate()).Collect()


```


### Or

Links two expressions of same type and performs Or, for example

``` go

	expr := Sifu.Expr[ComplexObjectToSearch]()

	query1 := expr.Prop("Name").StrEq("Jane").Or(expr.Prop("Flag").True())

    result := collections.From(&items).WhereEx(query1).Collect()

    result2 := collections.From(&items).Where(query1.Predicate()).Collect()

```

### Any

this method accepts another expression as an argument and if a criteria mets returns true. can be used when searching for nested values. 

for example each user type has many addresses and we want to see if specific city can be found within users address.

``` go

userExpr := Sifu.Expr[User]()

addrExpr:= Sifu.Expr[Address]()

userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").StrIn([]string{"Karaj", "Tehran"}),
		).Predicate(),

```

this example generates the code of a func(T) bool



### NumBigger, NumSmaller, NumEq

after selecting a prop this functions can be used for comparing numbers. just be sure to check the argument and destination data type. for example if your data structure has uint32 type, pass your number as an uint argument.

args:

accepts an argument of any

``` go

expr.Prop("Identifier").NumBigger(uint32(200))

expr.Prop("Grade").NumSmaller(202.2)

```


### KeyAs

this method generates a key selector function. can be used while grouping data with thor api.

returns a KeySelectorExpression[T, K] and a func(T) K after calling Predicate()


args:

1 - a *PropExpression[T]

``` go

Sifu.KeyAs[Person, bool](expr.Prop("Active")).Predicate()

```



### Less

this method use to determine the lesser value between to types of same kind. returns a func(T,T) bool. for example if we want to generate something like bellow

``` go

func(item1 T, item2 T) bool {
   return item1.Identifier < item2.Identifier
}

```

we can use less() like this:

``` go

expr.Prop("Identifier").Less()

```

can be used as a sorting function for thor api.



### LinkEq

this function links two types of same kind and returns func(T,T) bool. can be used when traversing tree's and expressing the relationship between them.

``` go
   userExpr.Prop("ParentId").LinkEq("Id").Predicate()

```



### SetString, 

sets the string value in a expression.

the following example finds the id of 55 and sets the name property as Jack. all using Sifu Expressions.

``` go

	expr := Sifu.Expr[ComplexObjectToSearch]()

	updatedResult := collections.From(&items).Where(expr.Prop("Id").NumEq(55).Predicate()).Update(expr.Prop("Name").SetString("Jack").Predicate()).Collect()


```


### SetBool

sets the bool value in a expression


``` go

	expr := Sifu.Expr[ComplexObjectToSearch]()

	updatedResult := collections.From(&items).Where(expr.Prop("Id").NumEq(55).Predicate()).Update(expr.Prop("Flag").SetBool(True).Predicate()).Collect()

```


### SetInt, SetUint, SetFloat

we have to be a bit more catious when mutating numeric fields. thats why we have extensions for different numbers.

SetInt()

args: an int64 value. use int64(val) as the argument for maximum safety.

SetUint()

args: an uint64 value. use uint64(val) as the argument for maximum safety.


SetFloat()

args : a float64 value. use float64(val) as the argument for maximum safety.


mutating numeric fields is indeed sensitive and can be error prone. but as promissed, we prevented the panics as much as we can. and we will working more and more to improve this area.


### StrApp

Appends to a string value selected by expression

``` go

expr := Sifu.Expr[ComplexObjectToSearch]()

updatedResult := collections.From(&items).Where(expr.Prop("Id").NumEq(55).Predicate()).Update(expr.Prop("Name").AppStr(" FamilyName").Predicate()).Collect()

```


### AppStruct(), SetStruct()

Appends Or Sets A Struct inside an expression.

use this for slices of a type.

``` go

user := Sifu.Expr[User]()

	updated_result := collections.From(&Users).Where(user.Prop("Id").NumEq(10).Predicate()).Update(user.Prop("Addr").AppStruct(Address{
		Street: "La",
		City:   "La",
		State:  "La",
		Zip:    "La",
		No:     20,
	}).Predicate()).Collect()
	

```

if your property is just a struct and not a slice of struct then use SetStruct() like below:

``` go

	updatedResult := collections.From(&Users).Where(user.Prop("Id").NumEq(184).Predicate()).Update(user.Prop("Addr").SetStruct(ForeignAddress{
		Country: "USA",
	}).Predicate()).Collect()


```



### Sifu Expressions Sample

to have a better understanding of how to use Sifu. we gathered a set of examples below:


``` go

	expr := Sifu.Expr[ComplexObjectToSearch]()

	result := collections.From(&items).Where(
		expr.Prop("Flag").True().And(
			expr.Prop("Name").StrEq("Jane"),
		).Predicate(),
	).Take(1).Update(expr.Prop("Name").StrApp(" Updated").Predicate()).Collect()


```


``` go

expr := Sifu.Expr[ComplexObjectToSearch]()

	query1 := expr.Prop("Name").StrEq("Jane").And(expr.Prop("Flag").True())

	query2 := expr.Prop("Name").StrEqNot("Jane").Or(expr.Prop("Flag").False())

	for i := 0; i < b.N; i++ {

		result := collections.From(&items).WhereEx(query1).Collect()

		result2 := collections.From(&result).AnyEx(query2).Assert()

		if result2 {
			b.Error("result should be false")
		}

	}

```

``` go

expr := Sifu.Expr[Person]()

	groupped := collections.Group[bool, Person](

		collections.From(&personList).Where(expr.Prop("Identifier").NumBigger(0).Predicate()),

		Sifu.KeyAs[Person, bool](expr.Prop("Active")).Predicate(),

).Collect()

```

``` go

    expr := Sifu.Expr[Person]()

    addrExpr := Sifu.Expr[Addr]()

	assert1 := collections.From(&personList).Where(expr.Prop("Address").Any(
		addrExpr.Prop("City").StrEq("NYC")).Predicate(),
	).Where(
		expr.Prop("Name").StrEq("Mark").Predicate(),
	).Collect()

```


``` go

	userExpr := Sifu.Expr[User]()
	addrExpr := Sifu.Expr[Address]()

    targetNode := collections.From(&users).Where(
		userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").StrEq("Tehran"),
		).Predicate(),
	).FindParentNode(
		userExpr.Prop("Id").NumEq(9).Predicate(),
		userExpr.Prop("ParentId").LinkEq("Id").Predicate(),
	)

```


``` go

userExpr := Sifu.Expr[User]()

	addrExpr := Sifu.Expr[Address]()

	targetNode1 := collections.From(&Users).WhereEx(

		userExpr.Prop("Addr").Any(

			addrExpr.Prop("City").StrEq("Karaj").Or(addrExpr.Prop("City").StrEq("Tehran")),
		),
	).FindRootNode(

		userExpr.Prop("Id").NumEq(7).Predicate(),

		userExpr.Prop("ParentId").LinkEq("Id").Predicate(),

		userExpr.Prop("Id").Less().Predicate(),
	)

```


``` go


	userExpr := Sifu.Expr[User]()
	addrExpr := Sifu.Expr[Address]()

	query := collections.From(&UserList).Where(
		userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").StrEq("Karaj"),
		).Predicate(),
	).Collect()

```

### Points To Consider When Using Sifu

1 - Sifu takes a hybrid approach to type safety.

2 - We developed a set of safety tests to ensure Sifu won’t panic at runtime when an invalid property is accessed or an invalid value is assigned.

3 - Sifu generates heavily optimized functions with no major bottlenecks. In fact, we rewrote our 50,000,000-record query and validation benchmark using Sifu, it completed in 1.1 seconds. In practice, most applications don’t hold 50M records in memory at once, so Sifu may not match the raw speed of native functions written against the ZenQL core directly. That said, it offers a strong balance of generated-code convenience, solid performance, and a more readable codebase.

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


### zenql Stream API

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

creates a stream or better a cursor from the rows of a sql database (postgresql, mysql supported). first we need to prepare for connecting to the database. in the example below we created a new db-context and started the connection. the context uses the pooling mechanism of the golang database package. so its compatible with concurrency and works with the standards of golang.

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

}).Throttle(0).Pipe() {


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

	}).Throttle(time.Millisecond * 500).Pipe() {

    }
```


### CallIf
Registers a callback if a criteria met.

args

1 - The If function or when to call

2 - The callback or what to call when the criteria met

``` go
ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for v := range FromData(ctx, items).Throttle(time.Millisecond*100).CallIf(func(search ComplexObjectToSearch) bool {

		return search.Id >= 15

	}, func(item ComplexObjectToSearch) {
	
	// log stream item
	
	}).pipe(){
	
	}
```

### StopIf

If the Criteria is valid the streaming stops immediately.

args:

1 - the criteria function. func(T) bool

2 - the cancel function

``` go
    ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for v := range FromData(ctx, items).Throttle(time.Millisecond*100).CallIf(func(search ComplexObjectToSearch) bool {

		return search.Id >= 15

	}, func(item ComplexObjectToSearch) {

	}).StopIf(func(search ComplexObjectToSearch) bool {

		return search.Id > 40

	}, cancel).Pipe() {

		fmt.Println(v)

	}
```

### Pipe

Instead of using the channel at the end of the pipeline chain to loop through you can use Pipe() function

``` go

	for v := range FromData(ctx, items).Throttle(time.Millisecond*100).Pipe() {

		fmt.Println(v)

	}
```


### Process

You can register a callback to process the stream of data one by one

args

1 - a func(T)

``` go

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	FromData(ctx, items).Throttle(time.Millisecond*500).StopIf(func(item ComplexObjectToSearch) bool {

		return item.Id >= 20

	}, cancel).Process(func(item ComplexObjectToSearch) {
	
		fmt.Println(item)
	
	})
	
```

### BackgroundProcess

just like the process but it runs on a background goroutine. you can process multiple chains concurrently!


args:

1 - a func(T)

2 - a pointer to a sync.waitgroup

``` go


	var wg sync.WaitGroup

	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())

	FromData(ctx, items).FilterStream(func(search ComplexObjectToSearch) bool {

		return search.Id >= 25

	}).Throttle(time.Millisecond*100).BackgroundProcess(&wg, func(item ComplexObjectToSearch) {
		fmt.Println(item)
	})

	FromData(ctx, items).FilterStream(func(search ComplexObjectToSearch) bool {

		return search.Id < 25

	}).Throttle(time.Millisecond*150).CallIf(func(item ComplexObjectToSearch) bool {

		return !item.Flag

	}, func(item ComplexObjectToSearch) {

		fmt.Println(item)

	}).BackgroundProcess(&wg, func(item ComplexObjectToSearch) {
		fmt.Println(item)
	})

	wg.Wait()

	defer cancel()



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



### The Database Module
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

Third‑Party Software Notice: This package includes/uses the third‑party MySQL driver github.com/go-sql-driver/mysql.
Copyright © The github.com/go-sql-driver/mysql authors.

Third‑Party Software Notice: This package includes/uses the third‑party Postgres driver github.com/lib/pq
Copyright © The github.com/lib/pq authors.

License applies as stated in those repository.
