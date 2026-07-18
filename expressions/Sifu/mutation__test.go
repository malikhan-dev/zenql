package Sifu

import (
	"testing"
)

func TestAppStr(t *testing.T) {

	type Student struct {
		Name string
	}

	stdList := []Student{
		{Name: "a"},
	}

	expr := Expr[Student]().Prop("Name").StrApp(" bcD").Predicate()

	stdList[0] = expr(stdList[0])

	if stdList[0].Name != "a bcD" {
		t.Errorf("expected %q, got %q", "a bcD", stdList[0].Name)
	}

}

func TestAppStrUnexported(t *testing.T) {

	type args struct {
		name string
	}

	var list []args

	list = append(list, args{name: "a"})

	expr := Expr[args]().Prop("name").StrApp(" bcD").Predicate()

	list[0] = expr(list[0])

	if list[0].name != "a" {
		t.Errorf("expected %q, got %q", "a bcD", list[0].name)
	} // name is unexported, no changes should happen

}

func TestSetBoolUnexported(t *testing.T) {

	type args struct {
		name     string
		pressent bool
	}

	var list []args

	list = append(list, args{name: "a", pressent: false})

	expr := Expr[args]().Prop("pressent").SetBool(true).Predicate()

	list[0] = expr(list[0])

	if list[0].pressent {
		t.Errorf("expected false, got true")
	} // pressent is unexported, no changes should happen

}

func TestSetBool(t *testing.T) {

	type args struct {
		name    string
		Present bool
	}

	var list []args

	list = append(list, args{name: "a", Present: false})

	expr := Expr[args]().Prop("Present").SetBool(true).Predicate()

	list[0] = expr(list[0])

	if !list[0].Present {
		t.Errorf("expected true got false")
	}

}

func TestShouldIgnoreWhenPropNotFound(t *testing.T) {

	type args struct {
		name    string
		Present bool
	}

	var list []args

	list = append(list, args{name: "a", Present: false})

	expr := Expr[args]().Prop("Lastname").SetBool(true).Predicate()

	list[0] = expr(list[0])

}

func TestSetStruct(t *testing.T) {

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

	expr := Expr[ForeignUser]().Prop("Addr").SetStruct(Address{City: "Tehran", Street: "Valiasr", No: 12}).Predicate()

	updateResult := expr(Users[0])

	if updateResult.Addr.City != "Tehran" {
		t.Errorf("Failed to set struct")
	}

	if updateResult.Addr.No != 12 {
		t.Errorf("Failed to set struct")
	}

	if updateResult.Addr.Street != "Valiasr" {
		t.Errorf("Failed to set struct")
	}
}

func TestAppStruct(t *testing.T) {

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
	}

	updater := Expr[User]().Prop("Addr").AppStruct(Address{City: "Karaj", Street: "Azimieh", No: 12}).Predicate()

	result := updater(Users[0])

	if len(result.Addr) != 2 {
		t.Errorf("Failed to append struct")
	}

	if result.Addr[1].City != "Karaj" {
		t.Errorf("Failed to append struct")
	}
	if result.Addr[0].City != "Tehran" {
		t.Errorf("Failed to append struct")
	}

}
