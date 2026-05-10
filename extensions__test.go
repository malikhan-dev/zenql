package lingo

import (
	"testing"
)

type ComplexObjectToSearch struct {
	Name string
	Age  int
	Id   int
	Flag bool
}

var items []ComplexObjectToSearch

const Heavy_Load = true

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

	if Heavy_Load {
		LoadLargeData()
	}

}
func LoadLargeData() {
	randFlag := false
	for i := 0; i < 2000000; i++ {

		items = append(items, ComplexObjectToSearch{
			Name: "Jane",
			Flag: randFlag,
			Id:   i,
			Age:  i,
		})
		randFlag = !randFlag
	}
}

func TestFindAllByPredicate(t *testing.T) {

	// must set the heavy_load const to false

	foundItem := FindByPredicate(items, func(search ComplexObjectToSearch) bool {
		return search.Flag == true
	})
	if foundItem == nil {
		t.Error("Find by Predicate failed")
	} else if len(*foundItem) == 0 {
		t.Error("Find by Predicate failed")
	}

	foundItem2 := FindByPredicate(items, func(search ComplexObjectToSearch) bool {
		return search.Flag == false
	})

	if foundItem2 == nil {
		t.Error("Find by Predicate failed")
	} else if len(*foundItem2) == 0 {
		t.Error("Find by Predicate failed")
	}
}

func TestFindFirstByPredicate(t *testing.T) {

	// must set the heavy_load const to false

	foundItem := FindFirstByPredicate(items, func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane"
	})

	if foundItem.Name != "Jane" {
		t.Error("Find First by Predicate Failed")
	} else if foundItem.Id <= 0 {
		t.Error("Find First by Predicate Failed")
	}

}

func TestRemoveFirstByPredicate(t *testing.T) {
	// must set the heavy_load const to false

	oldCount := len(items)

	var newItems *[]ComplexObjectToSearch

	newItems = RemoveFirstByPredicate(items, func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane"
	})

	newCount := len(*newItems)

	if newCount == oldCount {
		t.Error("remove first by Predicate failed")
	} else if (newCount + 1) != oldCount {
		t.Error("remove first by Predicate failed")
	}

}

func TestRemoveAllByPredicate(t *testing.T) {

	// must set the heavy_load const to false

	var newItems *[]ComplexObjectToSearch

	oldCount := len(items)

	newItems = RemoveByPredicate(items, func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane"
	})

	newCount := len(*newItems)

	if newCount == oldCount {
		t.Error("remove first by Predicate failed")
	}

}

func TestChainedSyntax(t *testing.T) {

	// must set the heavy_load const to false

	From(items).Where("Name", "John").Where("Flag", true).First().Collect()

	From(items).Where("Name", "John").Where("Flag", true).All().Collect()

	From(items).Where("Flag", true).AllOrDefault().Collect()

	From(items).Where("Name", "John").Where("Flag", 2).FirstOrDefault().Collect()

	From(items).Filter(func(search ComplexObjectToSearch) bool {

		return search.Name == "Jane" && search.Flag == true
	}).All().Collect()

}

func TestAny(t *testing.T) {

	// must set the heavy_load const to false

	result := Any(items, func(item ComplexObjectToSearch) bool {
		return item.Flag
	})

	if !result {
		t.Error("Any Function Malfunction")
	}

}

func Test_Should_Return_Errors(t *testing.T) {

	// must set the heavy_load const to false

	_, err := From(items).Where("NNNNAAAMMMEEE", 12).Where("FLLLAAAGGGG", "trvue").FirstOrDefault().Collect()

	if err == nil {
		t.Error(err)
	} else {
		if len(err) < 3 {
			t.Error("Error Collector Malfunction")
		}
	}

	_, err2 := From(items).Where("Name", 12).Where("Flag", true).FirstOrDefault().Collect()

	if err2 == nil {
		t.Error(err2)
	}

	_, err3 := From(items).Where("Name", "Jane").Where("Flag", true).FirstOrDefault().Collect()
	if err3 != nil {
		t.Error(err3)
	}

}

func BenchmarkAllOrDefaultCollector(b *testing.B) {

	// must set the heavy_load const to true

	res, err := From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}).AllOrDefault().Collect()

	if err != nil {
		b.Error(err)

	}

	if Any(res, func(item ComplexObjectToSearch) bool {
		return !item.Flag
	}) {
		b.Error("Wrong Data Fetched")
	}

}

func BenchmarkAllOrDefault(b *testing.B) {

	// must set the heavy_load const to true

	res := From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}).AllOrDefault()

	if len(res.Err) > 0 && res.Err != nil {
		b.Error(res.Err)

	}

	if Any(res.Items, func(item ComplexObjectToSearch) bool {
		return !item.Flag
	}) {
		b.Error("Wrong Data Fetched")
	}

}

func BenchmarkRawPerformance(b *testing.B) {

	_, err := From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}).AllOrDefault().Collect()

	if len(err) > 0 && err != nil {
		b.Error(err)

	}
}
func TestCollectRange(t *testing.T) {

	// must set the heavy_load const to true

	res, err := From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {

		return item.Id > 200000

	}).CollectRange(500)

	if len(err) > 0 && err != nil {
		t.Error(err)
	}

	if Any(res, func(item ComplexObjectToSearch) bool { return !item.Flag }) {
		t.Error("Wrong Data Fetched")
	}

}

func TestCollectRangeAndValidateData(t *testing.T) {

	res, err := From(items).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}).AllOrDefault().CollectRange(200)

	if err != nil {
		t.Error(err)

	}
	if len(res) > 200 {
		t.Error("Additional Data Fetched")
	}

	if res[0].Id != 200001 {
		t.Error("Wrong Data Fetched")
	}

	if res[199].Id != 200200 {
		t.Error("Wrong Data Fetched")
	}

}

func TestIssueWithUnsignedTypes(t *testing.T) {

	type Example struct {
		Id   uint32
		Name string
	}

	var Examples []Example

	Examples = append(Examples, Example{
		Id:   1,
		Name: "Sam",
	})

	Examples = append(Examples, Example{
		Id:   2,
		Name: "Dean",
	})

	_, err := From(Examples).Where("Id", uint32(2)).AllOrDefault().Collect()
	if err != nil {
		t.Error(err)
	}

}

func TestNestedSearch(t *testing.T) {

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

	res, err := From(UserList).Filter(func(user Users) bool {

		data, err := From(user.Addr).Where("City", "Karaj").Collect()

		if len(err) > 0 && err != nil {
			return true
		} else {
			return len(data) > 0
		}
	}).AllOrDefault().Collect()

	res2, err2 := From(UserList).Filter(func(user Users) bool {

		return Any(user.Addr, func(address Address) bool {
			return address.City == "Karaj"
		})

	}).AllOrDefault().Collect()

	if err2 != nil {
		t.Error(err2)
	}
	if err != nil {
		t.Error(err)
	}

	if len(res2) == 0 {
		t.Error("Data Fetch Failed")
	}

	if len(res) == 0 {
		t.Error("Data Fetch Failed")
	}
}

func TestGroupBy(t *testing.T) {

	type SysUser struct {
		Username    string
		Id          int
		Flag        bool
		AuthorityId uint32
	}

	users := []SysUser{
		{
			Username:    "removed user",
			Id:          -1,
			Flag:        true,
			AuthorityId: 1,
		},
		{
			Username:    "malikhan",
			Id:          1,
			Flag:        true,
			AuthorityId: 1,
		},
		{
			Username:    "Jackson",
			Id:          2,
			Flag:        true,
			AuthorityId: 1,
		},
		{
			Username:    "Miller",
			Id:          3,
			Flag:        false,
			AuthorityId: 14,
		},
		{
			Username:    "Alvarez",
			Id:          2,
			AuthorityId: 14,
			Flag:        false,
		},
	}

	res, err := GroupBy[bool, SysUser](From(users), "Flag").Collect()

	res2, err2 := GroupBy[uint32, SysUser](From(users).Filter(func(user SysUser) bool {

		return user.Id > 0

	}), "AuthorityId").Collect()

	if err2 != nil || err != nil {
		t.Error(err2)
		t.Error(err)
	}

	if len(res2[14]) != 2 {

		t.Error("Grouping Failed")
	}

	if len(res2[1]) != 2 {

		t.Error("Grouping Failed")
	}

	if len(res[true]) != 3 || len(res[false]) != 2 {

		t.Error("Grouping Failed")
	}
}

func TestPipeStream(t *testing.T) {

	for item := range From(items).Where("Flag", true).AllOrDefault().PipeStream(10) {

		if item.Err.Code != 0 {
			t.Error(item.Err)
		}
	}

}

func TestPipeGroupedStream(t *testing.T) {

	type student struct {
		Name    string
		Age     int
		Present bool
	}

	students := []student{
		{
			Name:    "jane",
			Age:     18,
			Present: true,
		},
		{
			Name:    "jack",
			Age:     19,
			Present: true,
		},
		{
			Name:    "james",
			Age:     20,
			Present: false,
		},
		{
			Name:    "john",
			Age:     20,
			Present: false,
		},
	}
	groupable := GroupBy[bool, student](From(students).AllOrDefault(), "Present")

	for item := range groupable.PipeStream(0) {

		for k, v := range item.Value {

			if k == true {

				for _, item := range v {
					if item.Name != "jane" && item.Name != "jack" {
						t.Error("Grouping Failed")
					}
				}

			} else {
				for _, item := range v {
					if item.Name != "james" && item.Name != "john" {
						t.Error("Grouping Failed")
					}
				}

			}
		}

		if item.Err.Code != 0 {
			t.Error(item.Err)
		}
	}

}
