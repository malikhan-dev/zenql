package collections

import (
	"fmt"
	"testing"

	"github.com/malikhan-dev/zenql/expressions/Sifu"
)

func TestSifuTrueFalseAnd(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	result := From(&items).Where(

		expr.Prop("Flag").True().And(
			expr.Prop("Name").EqualToString("Jane"),
		).Eval(),
	).Take(1).Update(expr.Prop("Name").AppendString(" Updated").Exec()).Collect()

	if len(result) > 20 {
		t.Errorf("Expected 20, got %d", len(result))
	}

	test := From(&result).Any(func(search ComplexObjectToSearch) bool {
		return !search.Flag
	}).Assert()

	if test {
		t.Errorf("Expected false, got true")
	}

	result2 := From(&items).Where(

		expr.Prop("Flag").True().And(
			expr.Prop("Name").EqualToString("Jane").And(
				expr.Prop("Name").EqualToString("Jack"),
			)).Eval(),
	).Take(20).Update(expr.Prop("Name").AppendString(" Updated").Exec()).Collect()

	if len(result2) > 0 {
		t.Errorf("Expected 0, got %d", len(result2))
	}
}

func BenchmarkQueryEngineWithSifu(b *testing.B) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	result := From(&items).Where(expr.Prop("Name").EqualToString("Jane").And(expr.Prop("Flag").True()).Eval()).Collect()

	result2 := From(&result).Any(expr.Prop("Name").NotEqualToString("Jane").Or(expr.Prop("Flag").False()).Eval()).Assert()

	if result2 {
		b.Error("result should be false")
	}

}

func TestGroupByNewWithSifu(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()
	res :=

		Group[bool, ComplexObjectToSearch](
			From(&items).Where(expr.Prop("Age").BiggerThanInt(20).Eval()),
			func(item ComplexObjectToSearch) bool {
				return item.Flag
			}).Collect()

	fmt.Println(res.Items[false][1])
	fmt.Println(res.Items[true][1])

}

func TestValidFilterWithSifu(t *testing.T) {

	type Student struct {
		Name     string
		Age      int
		Id       int
		Pressent bool
	}

	var students []Student

	students = append(students, Student{
		Name:     "Jane",
		Age:      20,
		Id:       1,
		Pressent: true,
	})

	students = append(students, Student{
		Name:     "John",
		Age:      22,
		Id:       2,
		Pressent: false,
	})

	expr := Sifu.Expr[Student]()

	results := From(&students).Where(expr.Prop("Name").EqualToString("Jane").And(expr.Prop("Pressent").False()).Eval()).Collect()

	if len(results) > 0 {
		t.Error("result should be empty")
	}

	result2 := From(&students).Any(expr.Prop("Name").EqualToString("Jane").And(expr.Prop("Pressent").True()).Eval()).Assert()

	if !result2 {
		t.Error("student should exists")
	}

	GroupResult := Group[bool, Student](From(&students).Where(expr.Prop("Age").BiggerThanInt(0).Eval()), func(student Student) bool {
		return student.Pressent
	}).Collect()

	if len(GroupResult.Items) != 2 {
		t.Error("Group Failed")
	}

	if len(GroupResult.Items[true]) != 1 {
		t.Error("Group Failed")
	}

	if len(GroupResult.Items[false]) != 1 {
		t.Error("Group Failed")
	}

	var jane = GroupResult.Items[true][0]

	if jane.Name != "Jane" {
		t.Error("Group Failed")
	}

	var john = GroupResult.Items[false][0]

	if john.Name != "John" {
		t.Error("Group Failed")
	}
}

func TestNestedSearch_Thor_WithSifu(t *testing.T) {

	var UserList []Users

	UserList = append(UserList, Users{
		Username: "jane",
		Id:       1,
		Addr: []Address{
			{
				City: "London",
				Id:   1,
				Flag: true,
			},
			{
				City: "Paris",
				Id:   2,
				Flag: false,
			},
			{
				City: "NYC",
				Id:   3,
				Flag: true,
			},
		},
	})

	UserList = append(UserList, Users{
		Username: "max",
		Id:       4,
		Addr: []Address{
			{
				City: "London",
				Id:   1,
				Flag: true,
			},
			{
				City: "Karaj",
				Id:   2,
				Flag: false,
			},
			{
				City: "NYC",
				Id:   3,
				Flag: true,
			},
		},
	})

	UserList = append(UserList, Users{
		Username: "marty",
		Id:       1,
		Addr: []Address{
			{
				City: "Los Angeles",
				Id:   5,
				Flag: true,
			},
			{
				City: "Karaj",
				Id:   7,
				Flag: false,
			},
		},
	})

	userExpr := Sifu.Expr[Users]()

	addrExpr := Sifu.Expr[Address]()

	res := From(&UserList).Where(

		userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").EqualToString("Karaj"),
		).Eval(),
	).Collect()

	fmt.Println(res)

}

func TestWhereAnySifu(t *testing.T) {
	var UserList []Users

	UserList = append(UserList, Users{
		Username: "jane",
		Id:       1,
		Addr: []Address{
			{
				City: "London",
				Id:   1,
				Flag: true,
			},
			{
				City: "Paris",
				Id:   2,
				Flag: false,
			},
			{
				City: "NYC",
				Id:   3,
				Flag: true,
			},
		},
	})

	UserList = append(UserList, Users{
		Username: "max",
		Id:       2,
		Addr: []Address{
			{
				City: "London",
				Id:   1,
				Flag: true,
			},
			{
				City: "Karaj",
				Id:   2,
				Flag: false,
			},
			{
				City: "NYC",
				Id:   3,
				Flag: true,
			},
		},
	})

	UserList = append(UserList, Users{
		Username: "mat",
		Id:       3,
		Addr: []Address{
			{
				City: "Los Angeles",
				Id:   5,
				Flag: true,
			},
			{
				City: "Karaj",
				Id:   7,
				Flag: false,
			},
		},
	})

	UserList = append(UserList, Users{
		Username: "Wade",
		Id:       4,
		Addr: []Address{
			{
				City: "Los Angeles",
				Id:   5,
				Flag: true,
			},
			{
				City: "Karaj",
				Id:   7,
				Flag: false,
			},
		},
	})

	UserList = append(UserList, Users{
		Username: "Wade",
		Id:       5,
		Addr: []Address{
			{
				City: "Los Angeles",
				Id:   5,
				Flag: true,
			},
			{
				City: "Karaj",
				Id:   7,
				Flag: false,
			},
		},
	})

	expr := Sifu.Expr[Users]()

	mat := From(&UserList).Where(expr.Prop("Id").SmallerThanInt(5).Eval()).Where(
		expr.Prop("Username").EqualToString("mat").Eval(),
	).Collect()

	if len(mat) <= 0 {
		t.Error("Find Failed")
	} else {
		fmt.Println(mat)
	}

	assertion2 := From(&UserList).Where(expr.Prop("Id").SmallerThanInt(5).Eval()).Any(expr.Prop("Username").EqualToString("Wade").Eval()).Assert()

	if !assertion2 {
		t.Error("Wade should exists")
	}
}
