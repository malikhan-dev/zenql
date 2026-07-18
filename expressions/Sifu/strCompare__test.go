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

			expr := Expr[args]().Prop("name").StrEq("a").Predicate()

			assert1 := expr(v)

			if assert1 != true {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}

		} else if x == 1 {
			expr := Expr[args]().Prop("name").StrEq("ab").Predicate()

			assert2 := expr(v)

			if assert2 == true {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}
		} else if x == 2 {
			expr := Expr[args]().Prop("name").StrEq("aBc").Predicate()

			assert2 := expr(v)

			if assert2 == false {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}
		}

	}

}

func TestStrIn(t *testing.T) {

	type args struct {
		name string
	}

	var list []args

	list = append(list, args{name: "a"})
	list = append(list, args{name: "aB"})
	list = append(list, args{name: "aBc"})

	for x, v := range list {

		expr := Expr[args]().Prop("name").StrIn([]string{"karaj", "Isfahan", "aB"}).Predicate()
		assertion := expr(v)

		if x == 0 {

			if assertion {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}

		} else if x == 1 {

			if !assertion {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}
		} else if x == 2 {

			if assertion {
				t.Errorf("Expected to be true %s %s", v.name, "a")
			}
		}

	}

}
