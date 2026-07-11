package Sifu

import "testing"

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
