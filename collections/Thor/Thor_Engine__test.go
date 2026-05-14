package Thor

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
func init() {
	LoadLargeData()
}

func TestQueryEngine(t *testing.T) {

	result := from(items).where(func(search ComplexObjectToSearch) bool {
		return search.Name == "Jane" && search.Flag == false
	}).Collect()

	result2 := from(result).any(func(search ComplexObjectToSearch) bool {
		return (search.Name != "Jane") || (search.Flag != false)
	}).Assert()

	if result2 {
		t.Error("result should be false")
	}

}

func TestGroupByNew(t *testing.T) {

	/*	ctx, _ := context.WithCancel(context.Background())
	 */
	res := Collect(Group[bool, ComplexObjectToSearch](from(items).where(func(search ComplexObjectToSearch) bool {

		return search.Age > 50000

	}), "Flag"))

	fmt.Println(res.Items[false][1])
	fmt.Println(res.Items[true][1])

	/*for _, v := range res.Items[false] {
		fmt.Println(v)
	}*/
	fmt.Println("==========================================================")
	fmt.Println("==========================================================")
	fmt.Println("==========================================================")
	/*for _, v := range res.Items[true] {
		fmt.Println(v)
	}*/

	/*fmt.Println(len(res.Items[true]))*/
	/*


		res2 := res.Items[true]

		res3 := res.Items[false]

		for i := range streams.FromData(ctx, 256, res2) {
			fmt.Println(i)
		}

		for i := range streams.FromData(ctx, 256, res3) {
			fmt.Println(i)
		}*/

}
