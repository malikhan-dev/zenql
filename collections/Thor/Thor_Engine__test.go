package Thor

import (
	"fmt"
	"testing"
)

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
	for i := 0; i < 200000; i++ {

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
	LoadLargeData()
}

func TestQueryEngine(t *testing.T) {

	result := From(items).Where(func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane" && search.Flag == false
	}).Collect()

	result2 := From(result).Any(func(search ComplexObjectToSearch) bool {
		return (search.Name != "Jane") || (search.Flag != false)
	}).Assert()

	if result2 {
		t.Error("result should be false")
	}

}

func TestGroupByNew(t *testing.T) {

	res :=

		Group[bool, ComplexObjectToSearch](
			From(items).Where(func(search ComplexObjectToSearch) bool {

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

	results := From(students).Where(func(search Student) bool {
		return search.Name == "Jane" && search.Pressent == false
	}).Collect()

	if len(results) > 0 {
		t.Error("result should be empty")
	}

	result2 := From(students).Any(func(search Student) bool {
		return search.Name == "Jane" && search.Pressent == true
	}).Assert()

	if !result2 {
		t.Error("student should exists")
	}

	GroupResult := Group[bool, Student](From(students).Where(func(student Student) bool {
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
		From(UserList).Where(func(user Users) bool {

			return From(user.Addr).Any(func(address Address) bool {
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

	mat := From(UserList).Where(func(user Users) bool {
		return user.Id < 5
	}).Where(func(user Users) bool {
		return user.Username == "mat"
	}).Collect()

	if len(mat) <= 0 {
		t.Error("Find Failed")
	} else {
		fmt.Println(mat)
	}

	assertion2 := From(UserList).Where(func(user Users) bool {
		return user.Id < 5
	}).Any(func(user Users) bool {
		return user.Username == "Wade"
	}).Assert()

	if !assertion2 {
		t.Error("Wade should exists")
	}
}
