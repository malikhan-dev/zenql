package streams

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var data []ComplexObjectToSearch

const StressTest = true

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

	if StressTest {
		LoadLargeData()
	}

}

func TestCompiledQuery(t *testing.T) {

	type student struct {
		Id   int
		Name string
		Age  int
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for i := range throttle(ctx, CompileStream(ctx, Filter(CompileFromQueryable(items), func(student ComplexObjectToSearch) bool {
		return !student.Flag
	})), time.Duration(250*time.Millisecond)) {
		fmt.Println(i)
	}

}

func TestCompiledQueryWithMapping(t *testing.T) {
	type student struct {
		Id   int
		Name string
		Age  int
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for i := range throttle(ctx, CompileStreamWithMapping(ctx, Filter(CompileFromQueryable(items), func(student ComplexObjectToSearch) bool {
		return !student.Flag
	}), func(items ComplexObjectToSearch) student {
		return student{
			Id:   items.Id,
			Name: items.Name + " student",
			Age:  items.Age,
		}
	}), time.Duration(250*time.Millisecond)) {
		fmt.Println(i)
	}

}
