package collections

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/malikhan-dev/zenql/expressions/Sifu"
)

func TestSifuTrueFalseAnd(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	result := From(&items).Where(
		expr.Prop("Flag").True().And(
			expr.Prop("Name").StrEq("Jane"),
		).Predicate(),
	).Take(1).Update(expr.Prop("Name").StrApp(" Updated").Predicate()).Collect()

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
			expr.Prop("Name").StrEq("Jane").And(
				expr.Prop("Name").StrEq("Jack"),
			)).Predicate(),
	).Take(20).Update(expr.Prop("Name").StrApp(" Updated").Predicate()).Collect()

	if len(result2) > 0 {
		t.Errorf("Expected 0, got %d", len(result2))
	}
}

func BenchmarkQueryEngineWithSifu(b *testing.B) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	query1 := expr.Prop("Name").StrEq("Jane").And(expr.Prop("Flag").True())

	query2 := expr.Prop("Name").StrEqNot("Jane").Or(expr.Prop("Flag").False())

	for i := 0; i < b.N; i++ {

		result := From(&items).WhereEx(query1).Collect()

		result2 := From(&result).AnyEx(query2).Assert()

		if result2 {
			b.Error("result should be false")
		}

	}

}

func TestGroupByNewWithSifu(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	res :=
		Group[bool, ComplexObjectToSearch](
			From(&items).Where(expr.Prop("Age").NumBigger(20).Predicate()),
			func(item ComplexObjectToSearch) bool {
				return item.Flag
			}).Collect()

	if res.Items[false][1].Id != 24 {
		t.Error("Expected 24,")
	}
	if res.Items[true][1].Id != 23 {
		t.Error("Expected 23,")
	}

}

func TestValidFilterWithSifu(t *testing.T) {

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

	results := From(&students).Where(expr.Prop("Name").StrEq("Jane").And(expr.Prop("Pressent").False()).Predicate()).Collect()

	if len(results) > 0 {
		t.Error("result should be empty")
	}

	result2 := From(&students).Any(expr.Prop("Name").StrEq("Jane").And(expr.Prop("Pressent").True()).Predicate()).Assert()

	if !result2 {
		t.Error("student should exists")
	}

	GroupResult := Group[bool, Student](From(&students).Where(expr.Prop("Age").NumBigger(0).Predicate()), func(student Student) bool {
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

	var UserList []User

	UserList = append(UserList, User{
		Name: "jane",
		Id:   1,
		Addr: []Address{
			{
				City:  "London",
				State: "London",
			},
			{
				City: "Paris",
			},
			{
				City: "NYC",
			},
		},
	})

	UserList = append(UserList, User{
		Name: "max",
		Id:   4,
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

	UserList = append(UserList, User{
		Name: "marty",
		Id:   1,
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

	userExpr := Sifu.Expr[User]()

	addrExpr := Sifu.Expr[Address]()

	res := From(&UserList).Where(
		userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").StrEq("Karaj"),
		).Predicate(),
	).Collect()

	fmt.Println(res)

}

func TestWhereAnySifu(t *testing.T) {
	var UserList []User

	UserList = append(UserList, User{
		Name: "jane",
		Id:   1,
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

	UserList = append(UserList, User{
		Name: "max",
		Id:   2,
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

	UserList = append(UserList, User{
		Name: "mat",
		Id:   3,
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

	UserList = append(UserList, User{
		Name: "Wade",
		Id:   4,
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

	UserList = append(UserList, User{
		Name: "Wade",
		Id:   5,
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

	expr := Sifu.Expr[User]()

	mat := From(&UserList).Where(
		expr.Prop("Id").NumSmaller(5).Predicate(),
	).Where(
		expr.Prop("Username").StrEq("mat").Predicate(),
	).Collect()

	if len(mat) <= 0 {
		t.Error("Find Failed")
	} else {
		fmt.Println(mat)
	}

	assertion2 := From(&UserList).Where(expr.Prop("Id").NumSmaller(5).Predicate()).Any(expr.Prop("Name").StrEq("Wade").Predicate()).Assert()

	if !assertion2 {
		t.Error("Wade should exists")
	}
}

func TestHeapInitializerWithSifu(t *testing.T) {

	type Person struct {
		Name       string
		LastName   string
		Identifier int
		Mail       string
		Active     bool
	}

	var personList []Person
	personList = append(personList, Person{
		Name:       "Jane",
		LastName:   "Jane",
		Identifier: 5,
		Mail:       "Jane@gmail.com",
		Active:     true,
	})

	personList = append(personList, Person{
		Name:       "Jack",
		LastName:   "Jack",
		Identifier: 3,
		Mail:       "Jack@gmail.com",
		Active:     true,
	})

	personList = append(personList, Person{
		Name:       "Jack",
		LastName:   "Jack",
		Identifier: 1,
		Mail:       "Jack@gmail.com",
		Active:     true,
	})

	personList = append(personList, Person{
		Name:       "Martin",
		LastName:   "Martin",
		Identifier: 18,
		Mail:       "Jack@gmail.com",
		Active:     false,
	})

	personList = append(personList, Person{
		Name:       "Marcus",
		LastName:   "Marcus",
		Identifier: 2,
		Mail:       "Jack@gmail.com",
		Active:     true,
	})

	expr := Sifu.Expr[Person]()

	result := From(&personList).Where(expr.Prop("Active").True().Predicate()).CollectSorted(expr.Prop("Identifier").Less().Predicate(), true)

	fmt.Println(result)
}

func TestOpFusionWithSifu(t *testing.T) {

	type Person struct {
		Name       string
		LastName   string
		Identifier int
		Mail       string
		Active     bool
	}
	var personList []Person

	active := false
	for i := 0; i <= 20; i++ {
		personList = append(personList, Person{
			Name:       "Jane",
			LastName:   "Jane",
			Identifier: 5,
			Active:     active,
		})
		active = !active
	}

	fmt.Println(personList)

	expr := Sifu.Expr[Person]()

	groupped := Group[bool, Person](
		From(&personList).Where(expr.Prop("Identifier").NumBigger(0).Predicate()), Sifu.KeyAs[Person, bool](expr.Prop("Active")).Predicate(),
	).Collect()

	fmt.Println(groupped.Items)
}

func TestFuseAnyWithSifu(t *testing.T) {

	type Addr struct {
		City string
	}
	type Person struct {
		Name       string
		LastName   string
		Identifier int
		Mail       string
		Active     bool
		Address    []Addr
	}
	var personList []Person

	personList = append(personList, Person{
		Name:       "Jane",
		LastName:   "Doe",
		Identifier: 1,
		Address: []Addr{
			{
				City: "Los Angeles",
			},
			{
				City: "Washington",
			},
		},
		Active: true,
	})

	personList = append(personList, Person{
		Name:       "Mark",
		LastName:   "Shepard",
		Identifier: 2,
		Address: []Addr{
			{
				City: "NYC",
			},
			{
				City: "LA",
			},
		},
		Active: true,
	})

	expr := Sifu.Expr[Person]()
	addrExpr := Sifu.Expr[Addr]()

	assert1 := From(&personList).Where(expr.Prop("Address").Any(
		addrExpr.Prop("City").StrEq("NYC")).Predicate(),
	).Where(
		expr.Prop("Name").StrEq("Mark").Predicate(),
	).Collect()

	assert2 := From(&personList).Where(expr.Prop("Address").Any(
		addrExpr.Prop("City").StrEq("NYC")).Predicate(),
	).Any(
		expr.Prop("Name").StrEq("Mark").Predicate(),
	).Assert()

	if len(assert1) <= 0 {
		t.Error("Find Failed")
	}
	fmt.Println(assert1)

	if !assert2 {
		t.Error("Find Failed")
	}
	fmt.Println(assert2)
}

func TestProject1WithSifu(t *testing.T) {

	type Addr struct {
		City string
	}
	type Person struct {
		Name       string
		LastName   string
		Identifier int
		Mail       string
		Active     bool
		Address    []Addr
	}
	var personList []Person

	type SysUser struct {
		FName   string
		LName   string
		Id      int
		Email   string
		Enabled bool
		Address string
	}

	personList = append(personList, Person{
		Name:       "Jane",
		LastName:   "Doe",
		Identifier: 1,
		Address: []Addr{
			{
				City: "Los Angeles",
			},
			{
				City: "Washington",
			},
		},
		Active: true,
	})

	personList = append(personList, Person{
		Name:       "Mark",
		LastName:   "Shepard",
		Identifier: 2,
		Address: []Addr{
			{
				City: "NYC",
			},
			{
				City: "LA",
			},
		},
		Active: true,
	})

	var newUsers []SysUser

	MapPersonToSysUser := func(person Person) SysUser {

		user := SysUser{
			FName:   person.Name,
			LName:   person.LastName,
			Id:      person.Identifier,
			Email:   person.Mail,
			Enabled: person.Active,
		}
		if len(person.Address) > 0 {
			user.Address = fmt.Sprintf("%s mapped", person.Address[0].City)
		}

		return user

	}

	expr := Sifu.Expr[Person]()
	newUsers = Project[Person, SysUser](
		From(&personList).Where(expr.Prop("Identifier").NumBigger(0).Predicate()),
		MapPersonToSysUser,
	)

	fmt.Println(newUsers)
}

func TestTakeOperatorWithSifu(t *testing.T) {
	var numbers []int
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	result := From(&numbers).Take(5).Collect()
	if len(result) != 5 {
		t.Errorf("Expected 5 items, got %d", len(result))
	}
	if result[4] != 5 {
		t.Errorf("Expected last item to be 5, got %d", result[4])
	}

	result2 := From(&numbers).Where(func(n int) bool {
		return n%2 == 0
	}).Take(2).Collect()

	if len(result2) != 2 {
		t.Errorf("Expected 2 even items, got %d", len(result2))
	}
	if result2[0] != 2 || result2[1] != 4 {
		t.Errorf("Expected [2, 4], got %v", result2)
	}
}

func TestTakeEdgeCasesWithSifu(t *testing.T) {
	var numbers []int
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	resultZero := From(&numbers).Take(0).Collect()
	if len(resultZero) != 0 {
		t.Errorf("Expected 0 items for Take(0), got %d", len(resultZero))
	}

	resultOverflow := From(&numbers).Take(100).Collect()
	if len(resultOverflow) != 10 {
		t.Errorf("Expected 10 items when taking more than available, got %d", len(resultOverflow))
	}

	resultFiltered := From(&numbers).Where(func(n int) bool {
		return n%2 != 0
	}).Take(3).Collect()

	/*UnSupported, needs stuff like divide and etc */

	if len(resultFiltered) != 3 {
		t.Errorf("Expected 3 items, got %d", len(resultFiltered))
	}
	if resultFiltered[0] != 1 || resultFiltered[1] != 3 || resultFiltered[2] != 5 {
		t.Errorf("Expected [1, 3, 5], got %v", resultFiltered)
	}
}

func TestGroupComprehensiveWithSifu(t *testing.T) {

	var employees []Employee
	employees = append(employees,
		Employee{Name: "Alice", Department: "IT", Age: 30},
		Employee{Name: "Bob", Department: "HR", Age: 25},
		Employee{Name: "Charlie", Department: "IT", Age: 35},
		Employee{Name: "David", Department: "HR", Age: 40},
		Employee{Name: "Eve", Department: "IT", Age: 28},
	)

	grouped := Group[string, Employee](
		From(&employees),
		func(e Employee) string { return e.Department },
	).Collect()

	if len(grouped.Items) != 2 {
		t.Errorf("Expected 2 departments, got %d", len(grouped.Items))
	}
	if len(grouped.Items["IT"]) != 3 {
		t.Errorf("Expected 3 employees in IT, got %d", len(grouped.Items["IT"]))
	}
	if len(grouped.Items["HR"]) != 2 {
		t.Errorf("Expected 2 employees in HR, got %d", len(grouped.Items["HR"]))
	}

	expr := Sifu.Expr[Employee]()
	groupedFiltered := Group[string, Employee](
		From(&employees).Where(expr.Prop("Age").NumBigger(28).Predicate()),
		func(e Employee) string { return e.Department },
	).Collect()

	if len(groupedFiltered.Items["IT"]) != 2 {
		t.Errorf("Expected 2 filtered employees in IT, got %d", len(groupedFiltered.Items["IT"]))
	}
	if len(groupedFiltered.Items["HR"]) != 1 {
		t.Errorf("Expected 1 filtered employee in HR, got %d", len(groupedFiltered.Items["HR"]))
	}

	var emptySlice []Employee
	groupedEmpty := Group[string, Employee](
		From(&emptySlice),
		func(e Employee) string { return e.Department },
	).Collect()

	if len(groupedEmpty.Items) != 0 {
		t.Errorf("Expected 0 groups for empty slice, got %d", len(groupedEmpty.Items))
	}
}

func TestProjectTakeWithSifu(t *testing.T) {
	type Employee struct {
		Name       string
		Department string
		Age        int
	}

	var employees []Employee
	employees = append(employees,
		Employee{Name: "Alice", Department: "IT", Age: 30},
		Employee{Name: "Bob", Department: "HR", Age: 25},
		Employee{Name: "Charlie", Department: "IT", Age: 35},
		Employee{Name: "David", Department: "HR", Age: 40},
		Employee{Name: "Eve", Department: "IT", Age: 28},
	)

	type InternalEmp struct {
		FullName string
		Dep      string
	}

	result := Project[Employee, InternalEmp](
		From(&employees).Take(2),
		func(e Employee) InternalEmp {
			return InternalEmp{
				FullName: e.Name,
				Dep:      e.Department,
			}
		},
	)

	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	result2 := Project[Employee, InternalEmp](
		From(&employees),
		func(e Employee) InternalEmp {
			return InternalEmp{
				FullName: e.Name,
				Dep:      e.Department,
			}
		},
	)

	if len(result2) != 5 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}
}

func TestProjectTakeSkipWithSifu(t *testing.T) {
	type Employee struct {
		Name       string
		Department string
		Age        int
	}

	var employees []Employee
	employees = append(employees,
		Employee{Name: "Alice", Department: "IT", Age: 30},
		Employee{Name: "Bob", Department: "HR", Age: 25},
		Employee{Name: "Charlie", Department: "IT", Age: 35},
		Employee{Name: "David", Department: "HR", Age: 40},
		Employee{Name: "Eve", Department: "IT", Age: 28},
	)

	type InternalEmp struct {
		FullName string
		Dep      string
	}

	result := Project[Employee, InternalEmp](
		From(&employees).Skip(2).Take(1),
		func(e Employee) InternalEmp {
			return InternalEmp{
				FullName: e.Name,
				Dep:      e.Department,
			}
		},
	)

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result[0].FullName != "Charlie" {
		t.Errorf("Expected Charlie, got %s", result[0].FullName)
	}

	if result[0].Dep != "IT" {
		t.Errorf("Expected IT, got %s", result[0].Dep)
	}

	result2 := Project[Employee, InternalEmp](
		From(&employees).Skip(4),
		func(e Employee) InternalEmp {
			return InternalEmp{
				FullName: e.Name,
				Dep:      e.Department,
			}
		},
	)

	if len(result2) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result2[0].FullName != "Eve" {
		t.Errorf("Expected Eve, got %s", result2[0].FullName)
	}

}

func TestProjectTakeSkipFilterWithSifu(t *testing.T) {
	type Employee struct {
		Name       string
		Department string
		Age        int
	}

	var employees []Employee
	employees = append(employees,
		Employee{Name: "Alice", Department: "IT", Age: 30},
		Employee{Name: "Bob", Department: "HR", Age: 25},
		Employee{Name: "Charlie", Department: "IT", Age: 35},
		Employee{Name: "David", Department: "HR", Age: 40},
		Employee{Name: "Eve", Department: "IT", Age: 28},
	)

	var expr *Sifu.TypeExpression[Employee]

	expr = Sifu.Expr[Employee]()

	result := Project[Employee, InternalEmp](
		From(&employees).Where(expr.Prop("Department").StrEq("IT").Predicate()).Skip(1).Take(1),
		func(e Employee) InternalEmp {
			return InternalEmp{
				FullName: e.Name,
				Dep:      e.Department,
			}
		},
	)

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result[0].FullName != "Charlie" {
		t.Errorf("Expected Charlie, got %s", result[0].FullName)
	}

	if result[0].Dep != "IT" {
		t.Errorf("Expected IT, got %s", result[0].Dep)
	}

	result2 := Project[Employee, InternalEmp](
		From(&employees).Where(expr.Prop("Department").StrEq("HR").Predicate()).Skip(1),
		func(e Employee) InternalEmp {
			return InternalEmp{
				FullName: e.Name,
				Dep:      e.Department,
			}
		},
	)

	if len(result2) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result2[0].FullName != "David" {
		t.Errorf("Expected David, got %s", result2[0].FullName)
	}

}

func TestCollectTakeSkipFilterWithSifu(t *testing.T) {
	type Employee struct {
		Name       string
		Department string
		Age        int
	}

	var employees []Employee
	employees = append(employees,
		Employee{Name: "Alice", Department: "IT", Age: 30},
		Employee{Name: "Bob", Department: "HR", Age: 25},
		Employee{Name: "Charlie", Department: "IT", Age: 35},
		Employee{Name: "David", Department: "HR", Age: 40},
		Employee{Name: "Eve", Department: "IT", Age: 28},
	)

	expr := Sifu.Expr[Employee]()
	result := From(&employees).Where(expr.Prop("Department").StrEq("IT").Predicate()).Take(1).Skip(1).Collect()

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result[0].Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", result[0].Name)
	}

	if result[0].Department != "IT" {
		t.Errorf("Expected IT, got %s", result[0].Department)
	}

	result2 := From(&employees).Where(expr.Prop("Department").StrEq("HR").Predicate()).Skip(1).Collect()

	if len(result2) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result2[0].Name != "David" {
		t.Errorf("Expected David, got %s", result2[0].Name)
	}

}

func TestCollectSortedTakeSkipWithSifu(t *testing.T) {

	var personList []Person
	personList = append(personList, Person{
		Name:       "Jane",
		LastName:   "Jane",
		Identifier: 5,
		Mail:       "Jane@gmail.com",
		Active:     true,
	})

	personList = append(personList, Person{
		Name:       "Jack",
		LastName:   "Jack",
		Identifier: 3,
		Mail:       "Jack@gmail.com",
		Active:     true,
	})

	personList = append(personList, Person{
		Name:       "Jack",
		LastName:   "Jack",
		Identifier: 1,
		Mail:       "Jack@gmail.com",
		Active:     true,
	})

	personList = append(personList, Person{
		Name:       "Martin",
		LastName:   "Martin",
		Identifier: 18,
		Mail:       "Jack@gmail.com",
		Active:     false,
	})

	personList = append(personList, Person{
		Name:       "Marcus",
		LastName:   "Marcus",
		Identifier: 2,
		Mail:       "Jack@gmail.com",
		Active:     true,
	})

	expr := Sifu.Expr[Person]()

	result := From(&personList).Skip(0).Take(1).Where(expr.Prop("Active").False().Predicate()).CollectSorted(expr.Prop("Identifier").Less().Predicate(), true)

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}
	if result[0].Name != "Martin" {
		t.Errorf("Expected Martin, got %s", result[0].Name)
	}

	result2 := From(&personList).Skip(1).Take(1).Where(expr.Prop("Active").False().Predicate()).CollectSorted(expr.Prop("Identifier").Less().Predicate(), true)

	if len(result2) != 0 {
		t.Errorf("Expected 0 items, got %d", len(result))
	}

	result3 := From(&personList).Skip(2).Take(4).Where(expr.Prop("Active").True().Predicate()).CollectSorted(expr.Prop("Identifier").Less().Predicate(), true)

	if result3[0].Identifier < result3[1].Identifier {
		t.Error("Expected item1, got ", result3[0].Identifier, ", ", result3[1].Identifier)
	}

	result4 :=
		From(&personList).Skip(2).Take(4).Where(expr.Prop("Active").True().Predicate()).CollectSorted(expr.Prop("Identifier").Less().Predicate(), false)

	if result4[0].Identifier > result4[1].Identifier {
		t.Error("Expected item1, got ", result3[0].Identifier, ", ", result3[1].Identifier)
	}

}

func TestGroupFilterTakeSkipWithSifu(t *testing.T) {

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

	students = append(students, Student{
		Name:     "Marry",
		Age:      22,
		Id:       3,
		Pressent: false,
	})

	students = append(students, Student{
		Name:     "Josh",
		Age:      22,
		Id:       3,
		Pressent: true,
	})

	expr := Sifu.Expr[Student]()
	GroupResult := Group[bool, Student](From(&students).Skip(2).Take(2).Where(expr.Prop("Age").NumBigger(0).Predicate()), Sifu.KeyAs[Student, bool](expr.Prop("Pressent")).Predicate()).Collect()

	if GroupResult.Items[true][0].Name != "Josh" {
		t.Error("Expected Josh, got ", GroupResult.Items[true][0].Name)
	}

	if GroupResult.Items[false][0].Name != "Marry" {
		t.Error("Expected Marry, got ", GroupResult.Items[false][0].Name)
	}

}

func TestUpdate2WithSifu(t *testing.T) {

	var CityList []city

	CityList = append(CityList, city{
		Name:   "Karaj",
		Id:     1,
		Active: true,
	})

	CityList = append(CityList, city{
		Name:   "Chaloos",
		Id:     2,
		Active: false,
	})

	CityList = append(CityList, city{
		Name:   "Tehran",
		Id:     3,
		Active: true,
	})

	CityList = append(CityList, city{
		Name:   "Isfahan",
		Id:     4,
		Active: true,
	})

	CityList = append(CityList, city{
		Name:   "Shiraz",
		Id:     5,
		Active: false,
	})

	expr := Sifu.Expr[city]()
	result := From(&CityList).Where(expr.Prop("Active").False().Predicate()).Skip(1).Take(1).
		Update(expr.Prop("Name").StrApp(" Deactivated").Predicate()).Collect()

	if len(result) != 1 {

		t.Errorf("Expected 1, got %d", len(result))

	}
	if result[0].Id != 5 {

		t.Errorf("Expected 5, got %d", result[0].Id)

	}

}

func TestFindParentNodeWithSifu(t *testing.T) {

	users := []User{
		{
			Name: "Ali",
			Age:  52,
			Id:   1,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Ahmad",
			Age:  52,
			Id:   184,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Reza",
			Age:  28,
			Id:   2,
			Addr: []Address{
				{City: "Karaj", Street: "Azadi", No: 8},
			},
			ParentId: 1,
		},
		{
			Name: "Dariush",
			Age:  52,
			Id:   185,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Sara",
			Age:  24,
			Id:   3,
			Addr: []Address{
				{City: "Shiraz", Street: "Chamran", No: 21},
			},
			ParentId: 1,
		},
		{
			Name: "Darvish",
			Age:  52,
			Id:   186,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Mina",
			Age:  31,
			Id:   4,
			Addr: []Address{
				{City: "Qom", Street: "Imam", No: 5},
			},
			ParentId: 0,
		},
		{
			Name: "Hossein",
			Age:  40,
			Id:   5,
			Addr: []Address{
				{City: "Mashhad", Street: "Sajjad", No: 18},
			},
			ParentId: 4,
		},
		{
			Name: "Niloofar",
			Age:  22,
			Id:   6,
			Addr: []Address{
				{City: "Isfahan", Street: "HashtBehesht", No: 33},
			},
			ParentId: 0,
		},
		{
			Name: "Amir",
			Age:  35,
			Id:   7,
			Addr: []Address{
				{City: "Qom", Street: "Bahonar", No: 9},
			},
			ParentId: 5,
		},
		{
			Name: "Fatemeh",
			Age:  27,
			Id:   8,
			Addr: []Address{
				{City: "Tehran", Street: "Kianpars", No: 44},
			},
			ParentId: 0,
		},
		{
			Name: "Mehdi",
			Age:  19,
			Id:   9,
			Addr: []Address{
				{City: "Tehran", Street: "Golha", No: 14},
			},
			ParentId: 8,
		},
		{
			Name: "Zahra",
			Age:  45,
			Id:   10,
			Addr: []Address{
				{City: "Tehran", Street: "Danesh", No: 2},
			},
			ParentId: 0,
		},
	}

	userExpr := Sifu.Expr[User]()
	addrExpr := Sifu.Expr[Address]()
	targetNode := From(&users).WhereEx(userExpr.Prop("Addr").Any(
		addrExpr.Prop("City").StrEq("Tehran"))).FindParentNode(userExpr.Prop("Id").NumEq(9).Predicate(),
		userExpr.Prop("ParentId").LinkEq("Id").Predicate())

	if targetNode.Id != 8 {
		t.Errorf("Expected 8, got %d", targetNode.Id)
	}

	targetNode2 := From(&users).WhereEx(userExpr.Prop("Addr").Any(
		addrExpr.Prop("City").StrEq("Zanjan"))).FindParentNode(userExpr.Prop("Id").NumEq(9).Predicate(),
		userExpr.Prop("ParentId").LinkEq("Id").Predicate())

	if targetNode2.Id > 0 {
		t.Errorf("Expected 0, got %d", targetNode2.Id)
	}

	targetNode3 := From(&users).WhereEx(userExpr.Prop("Addr").Any(
		addrExpr.Prop("City").StrEq("Tehran"))).FindParentNode(userExpr.Prop("Id").NumEq(1345).Predicate(),
		userExpr.Prop("ParentId").LinkEq("Id").Predicate())

	if targetNode3.Id > 0 {
		t.Errorf("Expected 0, got %d", targetNode3.Id)
	}

	targetNode4 :=
		From(&users).WhereEx(
			userExpr.Prop("Addr").Any(addrExpr.Prop("City").StrEq("Qom")),
		).FindParentNode(
			userExpr.Prop("Id").NumEq(7).Predicate(),
			userExpr.Prop("ParentId").LinkEq("Id").Predicate(),
		)

	if targetNode4.Id > 0 {
		t.Errorf("Expected 0, got %d", targetNode4.Id)
	}

}

func TestFindRootNodeWithSifu(t *testing.T) {

	Users := []User{
		{
			Name: "Ali",
			Age:  52,
			Id:   1,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Ahmad",
			Age:  52,
			Id:   184,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Reza",
			Age:  28,
			Id:   2,
			Addr: []Address{
				{City: "Karaj", Street: "Azadi", No: 8},
			},
			ParentId: 1,
		},
		{
			Name: "Dariush",
			Age:  52,
			Id:   185,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Sara",
			Age:  24,
			Id:   3,
			Addr: []Address{
				{City: "Shiraz", Street: "Chamran", No: 21},
			},
			ParentId: 1,
		},
		{
			Name: "Darvish",
			Age:  52,
			Id:   186,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Mina",
			Age:  31,
			Id:   4,
			Addr: []Address{
				{City: "Qom", Street: "Imam", No: 5},
			},
			ParentId: 0,
		},
		{
			Name: "Hossein",
			Age:  40,
			Id:   5,
			Addr: []Address{
				{City: "Mashhad", Street: "Sajjad", No: 18},
			},
			ParentId: 4,
		},
		{
			Name: "Niloofar",
			Age:  22,
			Id:   6,
			Addr: []Address{
				{City: "Isfahan", Street: "HashtBehesht", No: 33},
			},
			ParentId: 0,
		},
		{
			Name: "Amir",
			Age:  35,
			Id:   7,
			Addr: []Address{
				{City: "Qom", Street: "Bahonar", No: 9},
			},
			ParentId: 5,
		},
		{
			Name: "Fatemeh",
			Age:  27,
			Id:   8,
			Addr: []Address{
				{City: "Tehran", Street: "Kianpars", No: 44},
			},
			ParentId: 0,
		},
		{
			Name: "Mehdi",
			Age:  19,
			Id:   9,
			Addr: []Address{
				{City: "Tehran", Street: "Golha", No: 14},
			},
			ParentId: 8,
		},
		{
			Name: "Zahra",
			Age:  45,
			Id:   10,
			Addr: []Address{
				{City: "Tehran", Street: "Danesh", No: 2},
			},
			ParentId: 0,
		},
	}

	userExpr := Sifu.Expr[User]()

	addrExpr := Sifu.Expr[Address]()

	targetNode1 := From(&Users).WhereEx(

		userExpr.Prop("Addr").Any(

			addrExpr.Prop("City").StrEq("Qom").Or(addrExpr.Prop("City").StrEq("Mashhad")),
		),
	).FindRootNode(

		userExpr.Prop("Id").NumEq(7).Predicate(),

		userExpr.Prop("ParentId").LinkEq("Id").Predicate(),

		userExpr.Prop("Id").Less().Predicate(),
	)

	if targetNode1.Id != 4 {
		t.Errorf("Expected 4, got %d", targetNode1.Id)
	}

	targetNode2 := From(&Users).Where(userExpr.Prop("Addr").Any(userExpr.Prop("City").StrEq("LA")).Predicate()).
		FindRootNode(userExpr.Prop("Id").NumEq(7).Predicate(),

			userExpr.Prop("ParentId").LinkEq("Id").Predicate(),

			userExpr.Prop("Id").Less().Predicate())

	if targetNode2.Id != 0 {
		t.Errorf("Expected 0, got %d", targetNode2.Id)
	}

	targetNode3 := From(&Users).Where(
		userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").StrEq("Qom").Or(addrExpr.Prop("City").StrEq("Mashhad")),
		).Predicate(),
	).FindRootNode(
		userExpr.Prop("Id").NumEq(4).Predicate(),
		userExpr.Prop("ParentId").LinkEq("Id").Predicate(),
		userExpr.Prop("Id").Less().Predicate())

	if targetNode3.Id != 4 {
		t.Errorf("Expected 4, got %d", targetNode3.Id)
	}

	targetNode4 := From(&Users).Where(
		userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").StrIn([]string{"Qom", "Mashhad"}),
		).Predicate(),
	).FindRootNode(
		userExpr.Prop("Id").NumEq(18610).Predicate(),
		userExpr.Prop("ParentId").LinkEq("Id").Predicate(),
		userExpr.Prop("Id").Less().Predicate())

	if targetNode4.Id != 0 {
		t.Errorf("Expected 4, got %d", targetNode4.Id)
	}

	ctx, cancel := context.WithCancel(context.Background())

	targetNode5 := From(&Users).Where(
		userExpr.Prop("Addr").Any(
			addrExpr.Prop("City").StrIn([]string{"Qom", "Mashhad"}),
		).Predicate(),
	).TraverseRootNode(
		userExpr.Prop("Id").NumEq(7).Predicate(),
		userExpr.Prop("ParentId").LinkEq("Id").Predicate(),
		userExpr.Prop("Id").Less().Predicate(),
		ctx,
	)

	firstItemChecked := false

	for v := range targetNode5 {

		if !firstItemChecked {
			firstItemChecked = true
			if v.Id != 5 {
				t.Errorf("Expected 5, got %d", v.Id)
			}
		} else {
			if v.Id != 4 {
				t.Errorf("Expected 4, got %d", v.Id)
			}
		}
		time.Sleep(time.Millisecond * 500)

		fmt.Println(v)
		/*cancel()
		break*/
	}

	defer cancel()

}

func TestSetStr(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	updatedResult := From(&items).Where(expr.Prop("Id").NumEq(55).Predicate()).Update(expr.Prop("Name").SetString("mohammad").Predicate()).Collect()

	if updatedResult[0].Name != "mohammad" {
		t.Errorf("Expected mohammad, got %s", updatedResult[0].Name)
	}
}

func TestUpdateAppStruct(t *testing.T) {

	Users := []User{
		{
			Name: "Ali",
			Age:  52,
			Id:   1,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Ahmad",
			Age:  52,
			Id:   184,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Reza",
			Age:  28,
			Id:   2,
			Addr: []Address{
				{City: "Karaj", Street: "Azadi", No: 8},
			},
			ParentId: 1,
		},
		{
			Name: "Dariush",
			Age:  52,
			Id:   185,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Sara",
			Age:  24,
			Id:   3,
			Addr: []Address{
				{City: "Shiraz", Street: "Chamran", No: 21},
			},
			ParentId: 1,
		},
		{
			Name: "Darvish",
			Age:  52,
			Id:   186,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Mina",
			Age:  31,
			Id:   4,
			Addr: []Address{
				{City: "Qom", Street: "Imam", No: 5},
			},
			ParentId: 0,
		},
		{
			Name: "Hossein",
			Age:  40,
			Id:   5,
			Addr: []Address{
				{City: "Mashhad", Street: "Sajjad", No: 18},
			},
			ParentId: 4,
		},
		{
			Name: "Niloofar",
			Age:  22,
			Id:   6,
			Addr: []Address{
				{City: "Isfahan", Street: "HashtBehesht", No: 33},
			},
			ParentId: 0,
		},
		{
			Name: "Amir",
			Age:  35,
			Id:   7,
			Addr: []Address{
				{City: "Qom", Street: "Bahonar", No: 9},
			},
			ParentId: 5,
		},
		{
			Name: "Fatemeh",
			Age:  27,
			Id:   8,
			Addr: []Address{
				{City: "Tehran", Street: "Kianpars", No: 44},
			},
			ParentId: 0,
		},
		{
			Name: "Mehdi",
			Age:  19,
			Id:   9,
			Addr: []Address{
				{City: "Tehran", Street: "Golha", No: 14},
			},
			ParentId: 8,
		},
		{
			Name: "Zahra",
			Age:  45,
			Id:   10,
			Addr: []Address{
				{City: "Tehran", Street: "Danesh", No: 2},
			},
			ParentId: 0,
		},
	}

	user := Sifu.Expr[User]()

	updated_result := From(&Users).Where(user.Prop("Id").NumEq(10).Predicate()).Update(user.Prop("Addr").AppStruct(Address{
		Street: "La",
		City:   "La",
		State:  "La",
		Zip:    "La",
		No:     20,
	}).Predicate()).Collect()
	if updated_result[0].Addr[1].City != "La" {
		t.Errorf("Failed to set struct")
	}
	fmt.Println(updated_result[0].Addr)
}

func TestUpdateSetStruct(t *testing.T) {

	Users := []ForeignUser{

		{
			Name: "Ahmad",
			Age:  52,
			Id:   184,
			Addr: Address{
				City: "Tehran", Street: "Valiasr", No: 12,
			},
			ParentId: 0,
		},
	}

	user := Sifu.Expr[ForeignUser]()

	updatedResult := From(&Users).Where(user.Prop("Id").NumEq(184).Predicate()).Update(user.Prop("Addr").SetStruct(Address{
		Street: "La",
		City:   "La",
		State:  "La",
		Zip:    "La",
		No:     20,
	}).Predicate()).Collect()

	if updatedResult[0].Addr.City != "La" {
		t.Errorf("Failed to set struct")
	}
}
