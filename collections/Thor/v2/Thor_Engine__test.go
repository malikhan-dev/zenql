package collections

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/malikhan-dev/zenql/contracts/v2"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

type Address struct {
	City string
	Id   int
	Flag bool
}
type Users struct {
	Username string
	Id       int32
	Addr     []Address
}

type ComplexObjectToSearch struct {
	Name string
	Age  int
	Id   int
	Flag bool
}

var items []ComplexObjectToSearch

func LoadLargeData() {
	randFlag := false
	for i := 0; i < 50; i++ {

		items = append(items, ComplexObjectToSearch{
			Name: "Jane",
			Flag: randFlag,
			Id:   i,
			Age:  i,
		})
		randFlag = !randFlag
	}
}
func init() {

	contracts.SetMaxAllocGuard(25000000)
	LoadLargeData()

}

func BenchmarkQueryEngine(b *testing.B) {

	result := From(&items).Where(func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane" && search.Flag == false
	}).Collect()

	result2 := From(&result).Any(func(search ComplexObjectToSearch) bool {
		return (search.Name != "Jane") || (search.Flag != false)
	}).Assert()

	if result2 {
		b.Error("result should be false")
	}

}

func TestGroupByNew(t *testing.T) {

	res :=

		Group[bool, ComplexObjectToSearch](
			From(&items).Where(func(search ComplexObjectToSearch) bool {

				return search.Age > 20

			}),
			func(item ComplexObjectToSearch) bool {
				return item.Flag
			}).Collect()

	fmt.Println(res.Items[false][1])
	fmt.Println(res.Items[true][1])

	fmt.Println("==========================================================")
	fmt.Println("==========================================================")
	fmt.Println("==========================================================")

}

func TestValidFilter(t *testing.T) {

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

	results := From(&students).Where(func(search Student) bool {
		return search.Name == "Jane" && search.Pressent == false
	}).Collect()

	if len(results) > 0 {
		t.Error("result should be empty")
	}

	result2 := From(&students).Any(func(search Student) bool {
		return search.Name == "Jane" && search.Pressent == true
	}).Assert()

	if !result2 {
		t.Error("student should exists")
	}

	GroupResult := Group[bool, Student](From(&students).Where(func(student Student) bool {
		return student.Age > 0
	}), func(student Student) bool {
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

func TestNestedSearch_Thor(t *testing.T) {

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

	res :=
		From(&UserList).Where(func(user Users) bool {

			return From(&user.Addr).Any(func(address Address) bool {
				return address.City == "Karaj"
			}).Assert()

		}).Collect()

	fmt.Println(res)

}
func TestWhereAny(t *testing.T) {
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

	mat := From(&UserList).Where(func(user Users) bool {
		return user.Id < 5
	}).Where(func(user Users) bool {
		return user.Username == "mat"
	}).Collect()

	if len(mat) <= 0 {
		t.Error("Find Failed")
	} else {
		fmt.Println(mat)
	}

	assertion2 := From(&UserList).Where(func(user Users) bool {
		return user.Id < 5
	}).Any(func(user Users) bool {
		return user.Username == "Wade"
	}).Assert()

	if !assertion2 {
		t.Error("Wade should exists")
	}
}

func TestHeapInitializer(t *testing.T) {

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

	result := From(&personList).Where(func(person Person) bool {
		return person.Active == true
	}).CollectSorted(func(person Person, person2 Person) bool {
		return person.Identifier < person2.Identifier
	}, true)

	fmt.Println(result)
}

func TestOpFusion(t *testing.T) {

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

	groupped := Group[bool, Person](

		From(&personList).Where(func(person Person) bool {
			return person.Identifier > 0
		}), func(t Person) bool {
			return t.Active
		},
	).Collect()

	fmt.Println(groupped.Items)
}

func TestFuseAny(t *testing.T) {

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

	assert1 := From(&personList).Where(func(person Person) bool {

		return From(&person.Address).Any(func(addr Addr) bool {
			return addr.City == "NYC"
		}).Assert()

	}).Where(func(person Person) bool {
		return person.Name == "Mark"
	}).Collect()

	assert2 := From(&personList).Where(func(person Person) bool {

		return From(&person.Address).Any(func(addr Addr) bool {
			return addr.City == "NYC"
		}).Assert()

	}).Any(func(person Person) bool {
		return person.Name == "Mark"
	}).Assert()

	if len(assert1) <= 0 {
		t.Error("Find Failed")
	}
	fmt.Println(assert1)

	if !assert2 {
		t.Error("Find Failed")
	}
	fmt.Println(assert2)
}

func TestProject1(t *testing.T) {

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

	newUsers = Project[Person, SysUser](
		From(&personList).Where(func(person Person) bool {
			return person.Identifier > 0
		}),
		MapPersonToSysUser,
	)

	fmt.Println(newUsers)
}

func TestTakeOperator(t *testing.T) {
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

func TestTakeEdgeCases(t *testing.T) {
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

	if len(resultFiltered) != 3 {
		t.Errorf("Expected 3 items, got %d", len(resultFiltered))
	}
	if resultFiltered[0] != 1 || resultFiltered[1] != 3 || resultFiltered[2] != 5 {
		t.Errorf("Expected [1, 3, 5], got %v", resultFiltered)
	}
}

func TestGroupComprehensive(t *testing.T) {
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

	groupedFiltered := Group[string, Employee](
		From(&employees).Where(func(e Employee) bool { return e.Age > 28 }),
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

func TestProjectTake(t *testing.T) {
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

func TestProjectTakeSkip(t *testing.T) {
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

func TestProjectTakeSkipFilter(t *testing.T) {
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
		From(&employees).Where(func(employee Employee) bool {
			return employee.Department == "IT"
		}).Skip(1).Take(1),
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
		From(&employees).Where(func(employee Employee) bool {
			return employee.Department == "HR"
		}).Skip(1),
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

func TestCollectTakeSkipFilter(t *testing.T) {
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

	result := From(&employees).Where(func(employee Employee) bool {
		return employee.Department == "IT"
	}).Take(1).Skip(1).Collect()

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result[0].Name != "Charlie" {
		t.Errorf("Expected Charlie, got %s", result[0].Name)
	}

	if result[0].Department != "IT" {
		t.Errorf("Expected IT, got %s", result[0].Department)
	}

	result2 := From(&employees).Where(func(employee Employee) bool {
		return employee.Department == "HR"
	}).Skip(1).Collect()

	if len(result2) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}

	if result2[0].Name != "David" {
		t.Errorf("Expected David, got %s", result2[0].Name)
	}

}
func TestEarlyExitAndTakeSkipOrders(t *testing.T) {
	item1 := From(&items).Skip(11).Take(1).Collect()
	item2 := From(&items).Take(1).Skip(11).Collect()

	if item1[0].Id != item2[0].Id {
		t.Error("Expected item1, got ", item1[0].Id, ", ", item2[0].Id)
	}

}

func TestCollectSortedTakeSkip(t *testing.T) {

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

	result := From(&personList).Skip(0).Take(1).Where(func(person Person) bool {
		return person.Active == false
	}).CollectSorted(func(person Person, person2 Person) bool {
		return person.Identifier < person2.Identifier
	}, true)

	if len(result) != 1 {
		t.Errorf("Expected 1 items, got %d", len(result))
	}
	if result[0].Name != "Martin" {
		t.Errorf("Expected Martin, got %s", result[0].Name)
	}

	result2 := From(&personList).Skip(1).Take(1).Where(func(person Person) bool {
		return person.Active == false
	}).CollectSorted(func(person Person, person2 Person) bool {
		return person.Identifier < person2.Identifier
	}, true)

	if len(result2) != 0 {
		t.Errorf("Expected 0 items, got %d", len(result))
	}

	result3 := From(&personList).Skip(2).Take(4).Where(func(person Person) bool {
		return person.Active == true
	}).CollectSorted(func(person Person, person2 Person) bool {
		return person.Identifier < person2.Identifier
	}, true)

	if result3[0].Identifier < result3[1].Identifier {
		t.Error("Expected item1, got ", result3[0].Identifier, ", ", result3[1].Identifier)
	}

	result4 :=
		From(&personList).Skip(2).Take(4).Where(func(person Person) bool {
			return person.Active == true
		}).
			CollectSorted(func(person Person, person2 Person) bool {
				return person.Identifier < person2.Identifier
			}, false)

	if result4[0].Identifier > result4[1].Identifier {
		t.Error("Expected item1, got ", result3[0].Identifier, ", ", result3[1].Identifier)
	}

}

func TestGroupFilterTakeSkip(t *testing.T) {

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

	GroupResult := Group[bool, Student](From(&students).Skip(2).Take(2).Where(func(student Student) bool {
		return student.Age > 0
	}), func(student Student) bool {
		return student.Pressent
	}).Collect()

	if GroupResult.Items[true][0].Name != "Josh" {
		t.Error("Expected Josh, got ", GroupResult.Items[true][0].Name)
	}

	if GroupResult.Items[false][0].Name != "Marry" {
		t.Error("Expected Marry, got ", GroupResult.Items[false][0].Name)
	}

}

func TestFindParentNode(t *testing.T) {

	type address struct {
		Street string
		City   string
		State  string
		Zip    string
		No     int
	}
	type User struct {
		Name     string
		Age      int
		Id       int
		addr     []address
		ParentId int
	}

	users := []User{
		{
			Name: "Ali",
			Age:  52,
			Id:   1,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Ahmad",
			Age:  52,
			Id:   184,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Reza",
			Age:  28,
			Id:   2,
			addr: []address{
				{City: "Karaj", Street: "Azadi", No: 8},
			},
			ParentId: 1,
		},
		{
			Name: "Dariush",
			Age:  52,
			Id:   185,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Sara",
			Age:  24,
			Id:   3,
			addr: []address{
				{City: "Shiraz", Street: "Chamran", No: 21},
			},
			ParentId: 1,
		},
		{
			Name: "Darvish",
			Age:  52,
			Id:   186,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Mina",
			Age:  31,
			Id:   4,
			addr: []address{
				{City: "Qom", Street: "Imam", No: 5},
			},
			ParentId: 0,
		},
		{
			Name: "Hossein",
			Age:  40,
			Id:   5,
			addr: []address{
				{City: "Mashhad", Street: "Sajjad", No: 18},
			},
			ParentId: 4,
		},
		{
			Name: "Niloofar",
			Age:  22,
			Id:   6,
			addr: []address{
				{City: "Isfahan", Street: "HashtBehesht", No: 33},
			},
			ParentId: 0,
		},
		{
			Name: "Amir",
			Age:  35,
			Id:   7,
			addr: []address{
				{City: "Qom", Street: "Bahonar", No: 9},
			},
			ParentId: 5,
		},
		{
			Name: "Fatemeh",
			Age:  27,
			Id:   8,
			addr: []address{
				{City: "Tehran", Street: "Kianpars", No: 44},
			},
			ParentId: 0,
		},
		{
			Name: "Mehdi",
			Age:  19,
			Id:   9,
			addr: []address{
				{City: "Tehran", Street: "Golha", No: 14},
			},
			ParentId: 8,
		},
		{
			Name: "Zahra",
			Age:  45,
			Id:   10,
			addr: []address{
				{City: "Tehran", Street: "Danesh", No: 2},
			},
			ParentId: 0,
		},
	}
	targetNode := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Tehran"
		}).Assert()

	}).FindParentNode(func(user User) bool {

		return user.Id == 9

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id
	})

	if targetNode.Id != 8 {
		t.Errorf("Expected 8, got %d", targetNode.Id)
	}

	targetNode2 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Zanjan"
		}).Assert()

	}).FindParentNode(func(users User) bool {

		return users.Id == 9

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id
	})

	if targetNode2.Id > 0 {
		t.Errorf("Expected 0, got %d", targetNode2.Id)
	}

	targetNode3 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Tehran"
		}).Assert()

	}).FindParentNode(func(users User) bool {

		return users.Id == 1345

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id
	})

	if targetNode3.Id > 0 {
		t.Errorf("Expected 0, got %d", targetNode3.Id)
	}

	targetNode4 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(a address) bool {
			return a.City == "Qom"
		}).Assert()

	}).FindParentNode(func(u User) bool {

		return u.Id == 7

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	})

	if targetNode4.Id > 0 {
		t.Errorf("Expected 0, got %d", targetNode4.Id)
	}

}

func TestFindRootNode(t *testing.T) {
	type address struct {
		Street string
		City   string
		State  string
		Zip    string
		No     int
	}
	type User struct {
		Name     string
		Age      int
		Id       int
		addr     []address
		ParentId int
	}

	users := []User{
		{
			Name: "Ali",
			Age:  52,
			Id:   1,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Ahmad",
			Age:  52,
			Id:   184,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Reza",
			Age:  28,
			Id:   2,
			addr: []address{
				{City: "Karaj", Street: "Azadi", No: 8},
			},
			ParentId: 1,
		},
		{
			Name: "Dariush",
			Age:  52,
			Id:   185,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Sara",
			Age:  24,
			Id:   3,
			addr: []address{
				{City: "Shiraz", Street: "Chamran", No: 21},
			},
			ParentId: 1,
		},
		{
			Name: "Darvish",
			Age:  52,
			Id:   186,
			addr: []address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Mina",
			Age:  31,
			Id:   4,
			addr: []address{
				{City: "Qom", Street: "Imam", No: 5},
			},
			ParentId: 0,
		},
		{
			Name: "Hossein",
			Age:  40,
			Id:   5,
			addr: []address{
				{City: "Mashhad", Street: "Sajjad", No: 18},
			},
			ParentId: 4,
		},
		{
			Name: "Niloofar",
			Age:  22,
			Id:   6,
			addr: []address{
				{City: "Isfahan", Street: "HashtBehesht", No: 33},
			},
			ParentId: 0,
		},
		{
			Name: "Amir",
			Age:  35,
			Id:   7,
			addr: []address{
				{City: "Qom", Street: "Bahonar", No: 9},
			},
			ParentId: 5,
		},
		{
			Name: "Fatemeh",
			Age:  27,
			Id:   8,
			addr: []address{
				{City: "Tehran", Street: "Kianpars", No: 44},
			},
			ParentId: 0,
		},
		{
			Name: "Mehdi",
			Age:  19,
			Id:   9,
			addr: []address{
				{City: "Tehran", Street: "Golha", No: 14},
			},
			ParentId: 8,
		},
		{
			Name: "Zahra",
			Age:  45,
			Id:   10,
			addr: []address{
				{City: "Tehran", Street: "Danesh", No: 2},
			},
			ParentId: 0,
		},
	}

	targetNode1 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Qom" || address.City == "Mashhad"
		}).Assert()

	}).FindRootNode(func(user User) bool {

		return user.Id == 7

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	}, func(user User, user2 User) bool {

		return user.Id < user2.Id

	})

	if targetNode1.Id != 4 {
		t.Errorf("Expected 4, got %d", targetNode1.Id)
	}

	targetNode2 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "LA"
		}).Assert()

	}).FindRootNode(func(user User) bool {

		return user.Id == 7

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	}, func(user User, user2 User) bool {

		return user.Id < user2.Id

	})

	if targetNode2.Id != 0 {
		t.Errorf("Expected 0, got %d", targetNode2.Id)
	}

	targetNode3 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Qom" || address.City == "Mashhad"
		}).Assert()

	}).FindRootNode(func(user User) bool {

		return user.Id == 4

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	}, func(user User, user2 User) bool {

		return user.Id < user2.Id

	})

	if targetNode3.Id != 4 {
		t.Errorf("Expected 4, got %d", targetNode3.Id)
	}

	targetNode4 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Qom" || address.City == "Mashhad"
		}).Assert()

	}).FindRootNode(func(user User) bool {

		return user.Id == 18610

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	}, func(user User, user2 User) bool {

		return user.Id < user2.Id

	})

	if targetNode4.Id != 0 {
		t.Errorf("Expected 0, got %d", targetNode4.Id)
	}

	ctx, cancel := context.WithCancel(context.Background())

	targetNode5 := From(&users).Where(func(user User) bool {

		return From(&user.addr).Any(func(address address) bool {
			return address.City == "Qom" || address.City == "Mashhad"
		}).Assert()

	}).TraverseRootNode(func(user User) bool {

		return user.Id == 7

	}, func(child User, parent User) bool {

		return child.ParentId == parent.Id

	}, func(user User, user2 User) bool {

		return user.Id < user2.Id

	}, ctx)

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

func TestUpdateCollect(t *testing.T) {

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

	result := From(&CityList).Where(func(search city) bool {

		return search.Active

	}).Skip(1).Take(1).CollectUpdated(func(search city) city {

		search.Name += " is active"

		return search

	})

	if len(CityList[2].Name) > 6 {

		t.Errorf("Expected 6, got %d", len(CityList[2].Name))

	}

	if len(result) != 1 {

		t.Errorf("Expected 1, got %d", len(result))

	}
	if result[0].Id != 3 {

		t.Errorf("Expected 3, got %d", result[0].Id)

	}

	for _, v := range result {

		fmt.Println(v.Name)

		fmt.Println(v.Id)

	}
}

func TestUpdate2(t *testing.T) {

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

	result := From(&CityList).Where(func(search city) bool {

		return !search.Active

	}).Skip(1).Take(1).
		Update(func(search city) city {
			search.Name += " Deactivated"
			return search
		}).Collect()

	if len(result) != 1 {

		t.Errorf("Expected 1, got %d", len(result))

	}
	if result[0].Id != 5 {

		t.Errorf("Expected 5, got %d", result[0].Id)

	}

}
