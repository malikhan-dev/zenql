# Lingo

**Expressive data querying for Go — stream your data efficiently with Lingo Stream Api — inspired by LINQ, designed for flexibility, and built with generics.**

This library was written and designed by Mohammadreza Malikhan. The source code is free to use with proper attribution. This project is licensed under the MIT License (see the LICENSE file for details).


Lingo is a DSL (Domain Specific Language) for Go that helps you filter, search, validate, process and lately stream your data in a fluent and readable way. It is inspired by LINQ in C# and Streams in Java, while staying practical for Go developers.

Lingo supports two querying styles:

- **Dynamic field-based querying** for flexible runtime searches
- **Type-safe predicate-based querying** for safer and more explicit logic

Whether you want convenience, readability, or performance, Lingo gives you a clean way to work with data.

---


## Introducing lingo stream api's (v1.4.3)

Useful pipelines available for your needs to stream your data.

<img width="721" height="676" alt="Screenshot from 2026-05-11 21-08-37" src="https://github.com/user-attachments/assets/ecc015ab-f905-4891-839f-bededf93c5e5" />


## Installation

```bash
go get github.com/malikhan-dev/lingo@latest
go mod tidy
```

---

## Why Lingo?

- **Fluent query chaining**  
  Write data operations in a clean, readable flow

- **Two query styles**  
  Use dynamic field-based queries when flexibility matters, or type-safe predicates when you want stronger compile-time guarantees

- **Works with nested data**  
  Useful for searching inside slices of structs and nested collections

- **Generic core type**  
  Built around `Queryable[T]` using modern Go generics

- **Collectors for result unwrapping**  
  Keep chaining while querying, then explicitly unwrap results when needed or stream them.

---

## Quick Tour

### Simple API

<img width="1273" height="359" alt="Lingo5" src="https://github.com/user-attachments/assets/6b230dc5-01f7-47ad-b358-a1949f75c6b3" />

### Fast on large datasets

Lingo can query and validate large datasets efficiently.

**50,000,000 records queried and validated in under 5 seconds**  
(benchmark available in the test files)

<img width="1178" height="308" alt="Lingo6" src="https://github.com/user-attachments/assets/eed784b1-ecd0-4177-91a0-71108734ff15" />

### Expressive syntax

<img width="1287" height="465" alt="Screenshot from 2026-05-06 21-03-58" src="https://github.com/user-attachments/assets/d6ebf9cb-6a20-4a91-bb2d-aa6e9f019e47" />

**Focus on the problem you want to solve.**

---

## Core Concepts

### `Queryable[T]`

`Queryable[T]` is the core type passed between chained operations such as `Where`, `First`, `FirstOrDefault`, `All`, and `AllOrDefault`.

It wraps:

- a data slice: `[]T`
- an error slice: `[]error`

Collectors unwrap this type into concrete results.

```Go
type Queryable[T any] struct {
	Items []T
	Err   []OpError
}
```

---

### `From([]T)`

`From([]T)` creates a `Queryable[T]` from a slice and is usually the starting point of a query chain.

It accepts a slice of `[]T` and returns a pointer to `Queryable[T]`.

---

### `Where()`

`Where(fieldName, fieldValue)` filters a slice using a field name and value.

- `fieldName` must be a `string`
- `fieldValue` can be any type, but it must exactly match the actual type of the target field

This function modifies the current `Queryable[T]` and returns the same pointer for further chaining.

```Go
	_, err2 := From(items).Where("Name", 12).Where("Flag", true).FirstOrDefault().Collect()
```

**Important:** the field value must be exactly the same type as the struct field.  
For example, if the field type is `uint32`, you must pass `uint32(2)` instead of `2`.

```Go
	_, err := From(Examples).Where("Id", uint32(2)).AllOrDefault().Collect()
```

### `First()` and `FirstOrDefault()`

These functions return the first item in the current query chain.

- `First()` panics if no item is found
- `FirstOrDefault()` appends an error instead of panicking

Both still return a pointer to `Queryable[T]`.

---

### `All()` and `AllOrDefault()`

These functions return all items in the current query chain.

- `All()` panics if no item is found
- `AllOrDefault()` appends an error instead of panicking

Both still return a pointer to `Queryable[T]`.

---

## Collectors

**Available since version `v1.3.2`**

After a chained operation such as:

```go
lingo.From(data).Where(...).AllOrDefault()
```

you can use collectors to unwrap the `Queryable[T]` result into concrete values. 

- `Collect()` returns the full result set and errors
- `CollectRange(cnt)` returns a limited number of items based on the `cnt` argument, along with errors
-  `PipeStream(buffersize) formerly( CollectChan(buffersize) )` collect data and errors using go chan for your large data . available since version v1.4.0

```go
	res, err := From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {

		return item.Id > 200000

	}).AllOrDefault().CollectRange(500)

```


``` go

for item := range From(items).Where("Flag", true).AllOrDefault().PipeStream(256) {

		if item.Err.Code != 0 {
			t.Error(item.Err)
		}
	}




	groupable := lingo.GroupBy[bool, student](lingo.From(students).AllOrDefault(), "Present")

	for item := range groupable.PipeStream(0) {

		for k, v := range item.Value {

        }
    }

changed to PipeStream Since v1.4.1
```
---

PipeStream(size) returns a new type named CollectStream.

``` go

type CollectStream[T any] struct {
	Value T
	Err   OpError
}


* if Err.Code = 0 that means there is no error.
* The CollectChan() returns datas and errors in a Single type, which is CollectStream.
```

## Nested Search Example

Imagine you have a slice of users, and each user has multiple addresses.  
Now suppose you want to find all users where a specific city exists in their addresses.

Lingo makes this kind of nested search much easier to express.



```go

results, errors := From(UserList).Filter(func(user Users) bool {

		return Any(user.Addr, func(address Address) bool {
			return address.City == "Karaj"
		})

	}).AllOrDefault().Collect()

```

By reading this example, you can get a good sense of how the core functions work together in real use cases.

---

## `Any()`

`Any()` accepts:

- a slice
- a predicate function that returns a boolean

It returns `true` if at least one item matches the condition, otherwise `false`.

This is especially useful for nested queries.

```go

	result := Any(items, func(item ComplexObjectToSearch) bool {
		return item.Flag
	})

```



## `GroupBy()`

`GroupBy()` accepts:

- a queryable
- a string for property name

groups the data based on specific key.

```go

	result, err := GroupBy[bool, SysUser](From(users), "Flag").Collect()

	result, err2 := GroupBy[uint32, SysUser](From(users).Filter(func(user SysUser) bool {

		return user.Id > 0

	}), "AuthorityId").Collect()


```

# Lingo Stream API

When dealing with large datasets, it is not always recommended to collect everything into memory using the traditional `Queryable` execution model.

Lingo provides a Stream API that allows data to be processed incrementally as it flows through a pipeline. Also streams can be executed with a compiled mode mechanism which is 35% faster than regular streams.

There are 4 adapters available to initiate a stream:

---

## FromQueryable

Creates a stream from a `Queryable`.

args:

1- a context to manage cancelation.

2 - a buffer size of type int

3- a queryable.

it returns a chan of the generic type T



---

## FromData

Creates a stream from in-memory data.

args:

1- a context to manage cancelation.

2 - a buffer size of type int

3- a slice of objects.

it returns a chan of the generic type T


---

## FromChannel

Creates a stream from an existing Go channel.

args:


1- a context to manage cancelation.

2 - a buffer size of type int

3- a read chan of T.

it returns a chan of the generic type T




---

# Stream Pipelines

Once a stream is created, it can be processed using different pipeline stages.


## FilterStream

Works similarly to `Where()` or `Filter()`, but operates on streamed data.

args:

1 - a context to manage cancelation.

2 - a buffer size of type int.

3- a read chan of T.

4 - a func to filter the stream of data.  predicate func(T) bool


it returns a chan of the generic type T


---

## Throttle

Adds a delay between streamed items.

args:

1 - a context to manage cancelation.

2 - a read chan of T.

3 - a duration. waiting time in miliseconds.


it returns a chan of the generic type T




important:
- `100 * time.Millisecond`
- `0` for no delay

---

## MapStream

Transforms streamed data into another type.



1 - a context to manage cancelation.

2 - a read chan of T.

3 - a mapping function that maps the T to another type [M]

it returns a chan of M



---

# Example Of Streams

Streams respect `context.Context` cancellation to:
- prevent goroutine leaks
- support early termination
- properly manage pipeline lifecycle

Example:
```go


---
Process a Stream From Queryable
---

ctx, cancel := context.WithCancel(context.Background())
defer cancel()


	queryable := lingo.From(items)

	mappedStream := streams.MapStream[ComplexObjectToSearch, SimplerType](ctx,
		streams.Throttle(ctx,
			streams.FilterStream(ctx, buffer_size,
				streams.FromQueryable[ComplexObjectToSearch](ctx, buffer_size, *queryable),
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
    }


```



```go


---
Process a Stream From Data
---

ctx, cancel := context.WithCancel(context.Background())
defer cancel()


	queryable := lingo.From(items)

	mappedStream := streams.MapStream[ComplexObjectToSearch, SimplerType](ctx,
		streams.Throttle(ctx,
			streams.FilterStream(ctx, buffer_size,
				streams.FromData[ComplexObjectToSearch](ctx, buffer_size, *queryable),
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
    }


```




```go

---
Process a Stream From A Channel
---

ctx, cancel := context.WithCancel(context.Background())
defer cancel()


	queryable := lingo.From(items)

	mappedStream := streams.MapStream[ComplexObjectToSearch, SimplerType](ctx,
		streams.Throttle(ctx,
			streams.FilterStream(ctx, buffer_size,
				streams.FromChannel[ComplexObjectToSearch](ctx, buffer_size, *queryable),
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
    }


```





## Notes

- Dynamic field-based queries rely on exact type matching
- `First()` and `All()` are strict variants and may panic on empty results
- `FirstOrDefault()` and `AllOrDefault()` are safer alternatives when you want error collection instead of panics

---

## Project Status

Lingo is actively evolving, and more operators, examples, and documentation are on the way.

If you find it useful, feel free to star the repository and follow future updates.
