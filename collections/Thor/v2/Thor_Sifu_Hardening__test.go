package collections

import (
	"testing"

	"github.com/malikhan-dev/zenql/expressions/Sifu"
)

func TestBreakRuntime_InvalidPropName(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	query1 := expr.Prop("Nameasdasd").StrEq("asasdasd").And(expr.Prop("Flagwwwee").True()).Predicate()

	query2 := expr.Prop("Nam323e").StrEqNot("Janeff").Or(expr.Prop("Flagss").False()).Predicate()

	result := From(&items).Where(query1).Collect()

	result2 := From(&result).Any(query2).Assert()

	if result2 {
		t.Error("result should be false")
	}

}

func TestBreakRuntime_With_Trees(t *testing.T) {
	type Address struct {
		Street string
		City   string
		State  string
		Zip    string
		No     int
	}
	type User struct {
		Name     string
		Age      int
		Id       int
		Addr     []Address
		ParentId int
	}

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

	userExpr := Sifu.Expr[User]()

	addrExpr := Sifu.Expr[Address]()

	targetNode1 := From(&Users).WhereEx(

		userExpr.Prop("Addrdasd").Any(

			addrExpr.Prop("Citysdfsdf").StrEq("Qomqwewqe").Or(addrExpr.Prop("Cityfasfasf").StrEq("Maasdasdfasfadgshhad")),
		),
	).FindRootNode(

		userExpr.Prop("Id").NumEq(uint32(645)).Predicate(),

		userExpr.Prop("ParentIdqwee").LinkEq("Idffsad").Predicate(),

		userExpr.Prop("Idfasd").Less().Predicate(),
	)

	if targetNode1.Id != 0 {
		t.Errorf("Expected 4, got %d", targetNode1.Id)
	}

}

func TestBreakRuntimeWithInvalidFieldVal(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	query1 := expr.Prop("Id").StrEq("asasdasd").And(expr.Prop("Id").True()).Predicate()

	query2 := expr.Prop("Id").StrEqNot("Janeff").Or(expr.Prop("Name").False()).Predicate()

	result := From(&items).Where(query1).Collect()

	result2 := From(&result).Any(query2).Assert()

	if result2 {
		t.Error("result should be false")
	}

}

func TestBreakRuntimeWithInvalidTypeExpr(t *testing.T) {

	type Address struct {
		Street string
		City   string
		State  string
		Zip    string
		No     int
	}

	/*	result := From(&items).Where(Sifu.Expr[Address]().Prop("Id").StrEq("asasdasd").And(expr.Prop("Id").True()).Predicate()).Collect()

		result2 := From(&result).Any(
			expr.Prop("Id").StrEqNot("Janeff").Or(Sifu.Expr[Address]().Prop("Name").False()).Predicate()
		).Assert()
	*/ //Wont Compile

}
