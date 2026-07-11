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
			expr.Prop("Name").EqStr("Jane"),
		).Gen(),
	).Take(1).Update(expr.Prop("Name").AppStr(" Updated").Gen()).Collect()

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
			expr.Prop("Name").EqStr("Jane").And(
				expr.Prop("Name").EqStr("Jack"),
			)).Gen(),
	).Take(20).Update(expr.Prop("Name").AppStr(" Updated").Gen()).Collect()

	if len(result2) > 0 {
		t.Errorf("Expected 0, got %d", len(result2))
	}
}

func BenchmarkQueryEngineWithSifu(b *testing.B) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	query1 := expr.Prop("Name").EqStr("Jane").And(expr.Prop("Flag").True()).Gen()

	query2 := expr.Prop("Name").NotEqStr("Jane").Or(expr.Prop("Flag").False()).Gen()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		result := From(&items).Where(query1).Collect()

		result2 := From(&result).Any(query2).Assert()

		if result2 {
			b.Error("result should be false")
		}

	}

}

func TestGroupByNewWithSifu(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()
	res :=

		Group[bool, ComplexObjectToSearch](
			From(&items).Where(expr.Prop("Age").NumBigger(20).Gen()),
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

	results := From(&students).Where(expr.Prop("Name").EqStr("Jane").And(expr.Prop("Pressent").False()).Gen()).Collect()

	if len(results) > 0 {
		t.Error("result should be empty")
	}

	result2 := From(&students).Any(expr.Prop("Name").EqStr("Jane").And(expr.Prop("Pressent").True()).Gen()).Assert()

	if !result2 {
		t.Error("student should exists")
	}

	GroupResult := Group[bool, Student](From(&students).Where(expr.Prop("Age").NumBigger(0).Gen()), func(student Student) bool {
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
			addrExpr.Prop("City").EqStr("Karaj"),
		).Gen(),
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

	mat := From(&UserList).Where(
		expr.Prop("Id").NumSmaller(5).Gen(),
	).Where(
		expr.Prop("Username").EqStr("mat").Gen(),
	).Collect()

	if len(mat) <= 0 {
		t.Error("Find Failed")
	} else {
		fmt.Println(mat)
	}

	assertion2 := From(&UserList).Where(expr.Prop("Id").NumSmaller(5).Gen()).Any(expr.Prop("Username").EqStr("Wade").Gen()).Assert()

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

	result := From(&personList).Where(expr.Prop("Active").True().Gen()).CollectSorted(expr.Prop("Identifier").Less().Gen(), true)

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
		From(&personList).Where(expr.Prop("Identifier").NumBigger(0).Gen()), Sifu.KeyAs[Person, bool](expr.Prop("Active")).Gen(),
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
		addrExpr.Prop("City").EqStr("NYC")).Gen(),
	).Where(
		expr.Prop("Name").EqStr("Mark").Gen(),
	).Collect()

	assert2 := From(&personList).Where(expr.Prop("Address").Any(
		addrExpr.Prop("City").EqStr("NYC")).Gen(),
	).Any(
		expr.Prop("Name").EqStr("Mark").Gen(),
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
		From(&personList).Where(expr.Prop("Identifier").NumBigger(0).Gen()),
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
		From(&employees).Where(expr.Prop("Age").NumBigger(28).Gen()),
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

	type InternalEmp struct {
		FullName string
		Dep      string
	}

	var expr *Sifu.TypeExpression[Employee]

	expr = Sifu.Expr[Employee]()

	result := Project[Employee, InternalEmp](
		From(&employees).Where(expr.Prop("Department").EqStr("IT").Gen()).Skip(1).Take(1),
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
		From(&employees).Where(expr.Prop("Department").EqStr("HR").Gen()).Skip(1),
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
	result := From(&employees).Where(expr.Prop("Department").EqStr("IT").Gen()).Take(1).Skip(1).Collect()

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result[0].Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", result[0].Name)
	}

	if result[0].Department != "IT" {
		t.Errorf("Expected IT, got %s", result[0].Department)
	}

	result2 := From(&employees).Where(expr.Prop("Department").EqStr("HR").Gen()).Skip(1).Collect()

	if len(result2) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result2[0].Name != "David" {
		t.Errorf("Expected David, got %s", result2[0].Name)
	}

}

func TestCollectSortedTakeSkipWithSifu(t *testing.T) {

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

	result := From(&personList).Skip(0).Take(1).Where(expr.Prop("Active").False().Gen()).CollectSorted(expr.Prop("Identifier").Less().Gen(), true)

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}
	if result[0].Name != "Martin" {
		t.Errorf("Expected Martin, got %s", result[0].Name)
	}

	result2 := From(&personList).Skip(1).Take(1).Where(expr.Prop("Active").False().Gen()).CollectSorted(expr.Prop("Identifier").Less().Gen(), true)

	if len(result2) != 0 {
		t.Errorf("Expected 0 items, got %d", len(result))
	}

	result3 := From(&personList).Skip(2).Take(4).Where(expr.Prop("Active").True().Gen()).CollectSorted(expr.Prop("Identifier").Less().Gen(), true)

	if result3[0].Identifier < result3[1].Identifier {
		t.Error("Expected item1, got ", result3[0].Identifier, ", ", result3[1].Identifier)
	}

	result4 :=
		From(&personList).Skip(2).Take(4).Where(expr.Prop("Active").True().Gen()).CollectSorted(expr.Prop("Identifier").Less().Gen(), false)

	if result4[0].Identifier > result4[1].Identifier {
		t.Error("Expected item1, got ", result3[0].Identifier, ", ", result3[1].Identifier)
	}

}

func TestGroupFilterTakeSkipWithSifu(t *testing.T) {

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
	GroupResult := Group[bool, Student](From(&students).Skip(2).Take(2).Where(expr.Prop("Age").NumBigger(0).Gen()), Sifu.KeyAs[Student, bool](expr.Prop("Pressent")).Gen()).Collect()

	if GroupResult.Items[true][0].Name != "Josh" {
		t.Error("Expected Josh, got ", GroupResult.Items[true][0].Name)
	}

	if GroupResult.Items[false][0].Name != "Marry" {
		t.Error("Expected Marry, got ", GroupResult.Items[false][0].Name)
	}

}

func TestUpdate2WithSifu(t *testing.T) {

	type city struct {
		Name   string
		Id     int
		Active bool
	}

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
	result := From(&CityList).Where(expr.Prop("Active").False().Gen()).Skip(1).Take(1).
		Update(expr.Prop("Name").AppStr(" Deactivated").Gen()).Collect()

	if len(result) != 1 {

		t.Errorf("Expected 1, got %d", len(result))

	}
	if result[0].Id != 5 {

		t.Errorf("Expected 5, got %d", result[0].Id)

	}

}
