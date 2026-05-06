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
	} else {
		for _, item := range *foundItem {

			fmt.Println(item)
		}
	}

	foundItem2 := FindByPredicate(items, func(search ComplexObjectToSearch) bool {
		return search.Flag == false
	})

	fmt.Println("==================================================")
	if foundItem2 == nil {
		t.Error("Find by Predicate failed")
	} else if len(*foundItem2) == 0 {
		t.Error("Find by Predicate failed")
	} else {
		for _, item := range *foundItem2 {

			fmt.Println(item)
		}
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
	} else {
		for _, item := range items {

			fmt.Println(item)
		}
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
	} else {
		for _, item := range items {

			fmt.Println(item)
		}
	}

	fmt.Println(oldCount)
	fmt.Println(newCount)

}

func TestChainedSyntax(t *testing.T) {

	// must set the heavy_load const to false

	result, err1 := From(items).Where("Name", "John").Where("Flag", true).First().Collect()

	result2, err2 := From(items).Where("Name", "John").Where("Flag", true).All().Collect()

	result3, err3 := From(items).Where("Flag", true).AllOrDefault().Collect()

	result4, err4 := From(items).Where("Name", "John").Where("Flag", 2).FirstOrDefault().Collect()

	result5, err5 := From(items).Filter(func(search ComplexObjectToSearch) bool {

		return search.Name == "Jane" && search.Flag == true
	}).All().Collect()

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
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

	if result {
		fmt.Println(result)
	} else {
		t.Error("Any Function Malfunction")
	}

	assertion2 := Any(items, func(item ComplexObjectToSearch) bool {
		return item.Flag
	})

	if assertion2 {
		fmt.Println(assertion2)
	} else {
		t.Error("Any Function Malfunction")
	}

}

func Test_Should_Return_Errors(t *testing.T) {

	// must set the heavy_load const to false

	result, err := From(items).Where("NNNNAAAMMMEEE", 12).Where("FLLLAAAGGGG", "trvue").FirstOrDefault().Collect()

	if err == nil {
		t.Error("Error Collector Malfunction")
	} else {

		fmt.Println(result)

		if len(err) < 3 {
			t.Error("Error Collector Malfunction")
		}

		fmt.Println(err)
	}

	result2, err2 := From(items).Where("Name", 12).Where("Flag", true).FirstOrDefault().Collect()

	if err2 == nil {
		t.Error("Error Collector Malfunction")
	} else {
		fmt.Println(result2, err2)
	}

	result3, err3 := From(items).Where("Name", "Jane").Where("Flag", true).FirstOrDefault().Collect()
	if err3 != nil {
		t.Error("Error Collector Malfunction")
	} else {
		fmt.Println(result3, err3)
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

	fmt.Println(len(res))
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
