package lingo

import (
	"fmt"
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
	for i := 0; i < 50000000; i++ {

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

	fmt.Println("==================================================")

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

	fmt.Println("==================================================")
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
	} else {

		fmt.Println(foundItem.Age)
		fmt.Println(foundItem.Id)
		fmt.Println(foundItem.Name)
		fmt.Println(foundItem.Flag)
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

	fmt.Println(oldCount)
	fmt.Println(newCount)

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

	fmt.Println(oldCount)
	fmt.Println(newCount)

}

func TestChainedSyntax(t *testing.T) {

	// must set the heavy_load const to false

	_, err1 := From(items).Where("Name", "John").Where("Flag", true).First().Collect()

	_, err2 := From(items).Where("Name", "John").Where("Flag", true).All().Collect()

	_, err3 := From(items).Where("Flag", true).AllOrDefault().Collect()

	_, err4 := From(items).Where("Name", "John").Where("Flag", 2).FirstOrDefault().Collect()

	_, err5 := From(items).Filter(func(search ComplexObjectToSearch) bool {

		return search.Name == "Jane" && search.Flag == true
	}).All().Collect()

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(err4)
	fmt.Println(err5)

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
		fmt.Println(err)
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

	if len(res.err) > 0 && res.err != nil {
		b.Error(res.err)

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

	res, err := From(Examples).Where("Id", uint32(2)).AllOrDefault().Collect()
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
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

	fmt.Println("================================")
	fmt.Println(res)

	fmt.Println(err)

	fmt.Println("================================")

	fmt.Println(res2)
	fmt.Println(err2)

	fmt.Println("================================")
}
