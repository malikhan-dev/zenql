# Zen-Q

**Expressive data querying for Go — Streaming Capabilities — Fast Collection Processing, Flexible Design.**

### GitHub Achievements

[![Starstruck](https://img.shields.io/badge/GitHub-Starstruck-yellow?style=for-the-badge&logo=github)](https://github.com/users/malikhan-dev/achievements/starstruck)
[![Pull Shark](https://img.shields.io/badge/GitHub-Pull%20Shark-blue?style=for-the-badge&logo=github)](https://github.com/users/malikhan-dev/achievements/pull-shark)
![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/github/license/malikhan-dev/zenq?style=for-the-badge)

# License (MIT)

This library was written and designed by Mohammadreza Malikhan. The source code is free to use with proper attribution. This project is licensed under the MIT License (see the `LICENSE` file for details).

# Intro

zenq is a DSL (Domain Specific Language) for Go that helps you filter, search, validate, process, and stream your data in a fluent and readable way. It is inspired by LINQ in C# and Streams in Java, while staying practical for Go developers. make sure you review the benchmarks section at the end of this document. 

At its core, zenq is a modular library. Currently, it has two modules: **Collections** and **Streams**. 

There are two ways of processing collections:
1. Using default APIs.
2. Using the advanced collection query engine known as **Thor**. 

Thor is designed and architected to provide the maximum performance possible. It uses the operation fusion pattern to provide maximum speed and run the entire query chain in a single execution unit. Streams, on the other hand, use famous Golang concepts such as channels to allow the user to stream data while respecting the cancellation concepts of Go. 

Here are some examples:

```go
// A streaming example
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

count := 0
buffer_size := 10

for v := range streams.Throttle(ctx, streams.FilterStream(ctx, buffer_size, streams.FromData(ctx, buffer_size, items), func(item ComplexObjectToSearch) bool {
    return item.Id > 2
}), 0) {
    fmt.Println(v)
    count++

    if count == 100000 {
        cancel()
        break
    }
}
```

``` go
// Grouping Collections using the Thor engine
collections.Collect(
    collections.Group[bool, ComplexObjectToSearch](
        collections.From(items).Where(func(search ComplexObjectToSearch) bool {
            return search.Age > 20
        }),
        func(item ComplexObjectToSearch) bool {
            return item.Flag
        },
    ),
)
// Took around 3.8 seconds to filter and group a slice of 50,000,000 items

```

``` go
// The default APIs for collections
collections.From(items).Where("Name", "John").Where("Flag", true).First().Collect()
```

zenq supports two querying styles:
- **Dynamic field-based querying** for flexible runtime searches.
- **Type-safe predicate-based querying** for safer and more explicit logic.

Whether you want convenience, readability, or performance, zenq gives you a clean way to work with data.

---

## Installation

``` bash
go get github.com/malikhan-dev/zenq@latest

go mod tidy
```

## Default Collections API




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

# zenq Stream API

When dealing with large datasets, it is not always recommended to collect everything into memory using the traditional `Queryable` execution model.

zenq provides a Stream API that allows data to be processed incrementally as it flows through a pipeline. Also, streams can be executed with a compiled mode mechanism which is 35% faster than regular streams.

There are 3 adapters available to initiate a stream:



import path
``` go

streams  "github.com/malikhan-dev/zenq/streams"

```


## FromQueryable

Creates a stream from a `Queryable`.

**Args:**
1. A context to manage cancellation.
2. A buffer size of type `int`.
3. A queryable.

It returns a channel of the generic type `T`.

## FromData

Creates a stream from in-memory data.

**Args:**
1. A context to manage cancellation.
2. A buffer size of type `int`.
3. A slice of objects.

It returns a channel of the generic type `T`.

## FromChannel

Creates a stream from an existing Go channel.

**Args:**
1. A context to manage cancellation.
2. A buffer size of type `int`.
3. A read channel of `T`.

It returns a channel of the generic type `T`.

---

# Stream Pipelines

Once a stream is created, it can be processed using different pipeline stages.

## FilterStream

Works similarly to `Where()` or `Filter()`, but operates on streamed data.

**Args:**
1. A context to manage cancellation.
2. A buffer size of type `int`.
3. A read channel of `T`.
4. A function to filter the stream of data (`predicate func(T) bool`).

It returns a channel of the generic type `T`.

## Throttle

Adds a delay between streamed items.

**Args:**
1. A context to manage cancellation.
2. A read channel of `T`.
3. A duration (waiting time in milliseconds).

It returns a channel of the generic type `T`.

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

# Example Of Streams

Streams respect `context.Context` cancellation to:
- Prevent goroutine leaks.
- Support early termination.
- Properly manage pipeline lifecycle.

``` go
// Process a Stream From Queryable

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

queryable := zenq.From(items)

mappedStream := streams.MapStream[ComplexObjectToSearch, SimplerType](ctx,
    streams.Throttle(ctx,
        streams.FilterStream(ctx, buffer_size,
            streams.FromQueryable[ComplexObjectToSearch](ctx, buffer_size, *queryable),
            func(item ComplexObjectToSearch) bool {
                return item.Id > 0
            }), 0), 
    func(search ComplexObjectToSearch) SimplerType {
        return SimplerType{
            Enabled: search.Flag,
            Id:      search.Id,
            Name:    search.Name,
        }
    })

for v := range mappedStream {
    // Process stream items
}
```

``` go
// Process a Stream From Data

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

mappedStream := streams.MapStream[ComplexObjectToSearch, SimplerType](ctx,
    streams.Throttle(ctx,
        streams.FilterStream(ctx, buffer_size,
            streams.FromData[ComplexObjectToSearch](ctx, buffer_size, items),
            func(item ComplexObjectToSearch) bool {
                return item.Id > 0
            }), 0), 
    func(search ComplexObjectToSearch) SimplerType {
        return SimplerType{
            Enabled: search.Flag,
            Id:      search.Id,
            Name:    search.Name,
        }
    })

for v := range mappedStream {
    // Process stream items
}
```

``` go
// Process a Stream From A Channel

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

mappedStream := streams.MapStream[ComplexObjectToSearch, SimplerType](ctx,
    streams.Throttle(ctx,
        streams.FilterStream(ctx, buffer_size,
            streams.FromChannel[ComplexObjectToSearch](ctx, buffer_size, myChannel),
            func(item ComplexObjectToSearch) bool {
                return item.Id > 0
            }), 0), 
    func(search ComplexObjectToSearch) SimplerType {
        return SimplerType{
            Enabled: search.Flag,
            Id:      search.Id,
            Name:    search.Name,
        }
    })

for v := range mappedStream {
    // Process stream items
}
```


# Thor Engine APIs For Collection Processing

A faster, more Go-idiomatic alternative to the default collections API is to use the **Thor** engine to query your data. The Thor engine uses the operator fusion pattern to ensure maximum speed and a single execution unit.


import path
``` go
collections "github.com/malikhan-dev/zenq/collections/Thor"
```

### Core Concepts:
1. **`CollectionCompiledQueryable[T]`**: After each chain of operation, we use this type as a contract (much like `Queryable` in the default collections API).
2. **`AssertCompiledQueryable[T]`**: In our query chains, if we want to assert the result like the `Any()` operator, this is the output type.
3. **`GroupCompiledQueryable[K, T]`**: After a grouping operation, the returning type is `GroupCompiledQueryable`.

All three types nest `CompiledQueryable[T]` inside them. `CompiledQueryable` represents the result of the operation in the `Items` property and the list of operators.

``` go
type CompiledQueryable[T any] struct {
    Operators []zenqOperator[T]
    Items     *[]T
}
```

Thor Engine APIs are as follows:

- **`From[T any]`**: Accepts a slice of `[]T` and returns a `*CollectionCompiledQueryable[T]` to initiate a query chain.
- **`Where[T any]`**: Accepts a function `func(T) bool` as an argument, filters the collection, and returns a `*CollectionCompiledQueryable[T]`.
- **`Collect()`**: Collects the result and returns the `CollectionCompiledQueryable[T]` which holds the data.

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

- **`Group` and `Collect`**: The `Group` function expects a `CompiledQueryable[T]` as an argument and a Key Selector function. For collecting the result of a group, we can use the `collections.Collect()` function.

A grouping example: filtering users whose age is greater than 20 and grouping them by their presence:

``` go
res := collections.Collect(
    collections.Group[bool, ComplexObjectToSearch](
        collections.From(items).Where(func(search ComplexObjectToSearch) bool {
            return search.Age > 20
        }),
        func(item ComplexObjectToSearch) bool {
            return item.Flag
        },
    ),
)

fmt.Println(res.Items[false][1])
fmt.Println(res.Items[true][1])

```

- **`Assert()`**: Asserts the collection on a given criteria.

``` go
result2 := collections.From(result).Any(func(search ComplexObjectToSearch) bool {
    return (search.Name != "Jane") || (search.Flag != false)
}).Assert()

```

## benchmark

in a slice of 50,000,000 users it took less than 2 seconds just to filter them and around 4 seconds to filter then group the items.


<img width="1138" height="893" alt="bench2" src="https://github.com/user-attachments/assets/644db764-425e-4a70-97b3-1b649ca9864f" />
<img width="1133" height="772" alt="bench1" src="https://github.com/user-attachments/assets/6dca9160-e0ed-4c04-bcb2-8ea519a7f27d" />


## Project Status

zenq is actively evolving, and more operators, examples, and documentation are on the way.

If you find it useful, feel free to star the repository (it motivates us) and follow future updates!
