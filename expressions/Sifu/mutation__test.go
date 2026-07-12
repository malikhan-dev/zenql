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
