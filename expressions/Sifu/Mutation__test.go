package Sifu

import (
	"testing"
)

func TestEqStr(t *testing.T) {

	type args struct {
		name string
	}

	var list []args

	list = append(list, args{name: "a"})
	list = append(list, args{name: "aB"})
	list = append(list, args{name: "aBc"})

	for x, v := range list {

		if x == 0 {

			expr := Expr[args]().Prop("name").EqStr("a").Gen()

			assert1 := expr(v)

			if assert1 != true {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}

		} else if x == 1 {
			expr := Expr[args]().Prop("name").EqStr("ab").Gen()

			assert2 := expr(v)

			if assert2 == true {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}
		} else if x == 2 {
			expr := Expr[args]().Prop("name").EqStr("aBc").Gen()

			assert2 := expr(v)

			if assert2 == false {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}
		}

	}

}

func TestAppStr(t *testing.T) {

	type Student struct {
		Name string
	}

	stdList := []Student{
		{Name: "a"},
	}

	expr := Expr[Student]().Prop("Name").AppStr(" bcD").Gen()

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

	expr := Expr[args]().Prop("name").AppStr(" bcD").Gen()

	list[0] = expr(list[0])

	if list[0].name != "a" {
		t.Errorf("expected %q, got %q", "a bcD", list[0].name)
	} // name is unexported, no changes should happen

}
