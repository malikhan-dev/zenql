# Lingo

**Expressive data querying for Go — inspired by LINQ, designed for flexibility, and built with generics.**

Lingo is a data querying framework for Go that helps you filter, search, validate, and process data in a fluent and readable way. It is inspired by LINQ in C# and Streams in Java, while staying practical for Go developers.

Lingo supports two querying styles:

- **Dynamic field-based querying** for flexible runtime searches
- **Type-safe predicate-based querying** for safer and more explicit logic

Whether you want convenience, readability, or performance, Lingo gives you a clean way to work with data.

---

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
  Keep chaining while querying, then explicitly unwrap results when needed

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

<img width="480" height="270" alt="Screenshot from 2026-05-07 11-06-07" src="https://github.com/user-attachments/assets/5a8c20c2-9f87-46d5-a1c6-93703da09de8" />

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

<img width="466" height="42" alt="Screenshot from 2026-05-07 11-21-06" src="https://github.com/user-attachments/assets/573c4717-86ba-4bbf-b7d3-78872d3fa02c" />

**Important:** the field value must be exactly the same type as the struct field.  
For example, if the field type is `uint32`, you must pass `uint32(2)` instead of `2`.

<img width="377" height="40" alt="Screenshot from 2026-05-07 11-24-48" src="https://github.com/user-attachments/assets/224ac952-4f3f-4fda-a480-099c6e2ef3d9" />

---

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

After calling a collector, the result is no longer a pointer to `Queryable[T]`.

<img width="705" height="38" alt="Screenshot from 2026-05-07 11-50-30" src="https://github.com/user-attachments/assets/bcb085a7-634d-43bd-8b78-41358fd37b6e" />

---

## Nested Search Example

Imagine you have a slice of users, and each user has multiple addresses.  
Now suppose you want to find all users where a specific city exists in their addresses.

Lingo makes this kind of nested search much easier to express.

<img width="1357" height="480" alt="Lingo8" src="https://github.com/user-attachments/assets/5d24d6c4-b998-4d62-b09a-6a945d138b8c" />

By reading this example, you can get a good sense of how the core functions work together in real use cases.

---

## `Any()`

`Any()` accepts:

- a slice
- a predicate function that returns a boolean

It returns `true` if at least one item matches the condition, otherwise `false`.

This is especially useful for nested queries.

<img width="694" height="81" alt="Screenshot from 2026-05-07 12-00-16" src="https://github.com/user-attachments/assets/182516c6-dc0e-4a45-bc3d-c3d38058c4b2" />

---

## Notes

- Dynamic field-based queries rely on exact type matching
- `First()` and `All()` are strict variants and may panic on empty results
- `FirstOrDefault()` and `AllOrDefault()` are safer alternatives when you want error collection instead of panics

---

## Project Status

Lingo is actively evolving, and more operators, examples, and documentation are on the way.

If you find it useful, feel free to star the repository and follow future updates.
