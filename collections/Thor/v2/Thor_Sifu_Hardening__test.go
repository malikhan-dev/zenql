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

	/*	result := From(&items).Where(Sifu.Expr[Address]().Prop("Id").StrEq("asasdasd").And(expr.Prop("Id").True()).Predicate()).Collect()

		result2 := From(&result).Any(
			expr.Prop("Id").StrEqNot("Janeff").Or(Sifu.Expr[Address]().Prop("Name").False()).Predicate()
		).Assert()
	*/ //Wont Compile

}

func TestBreakRuntimeWithSetStrOnInvalidProp(t *testing.T) {

	expr := Sifu.Expr[ComplexObjectToSearch]()

	updatedResult := From(&items).Where(expr.Prop("Id").NumEq(55).Predicate()).Update(expr.Prop("Id").SetString("mohammad").Predicate()).Collect()

	if updatedResult[0].Id != 55 {
		t.Error("Expected 55, got ", updatedResult[0].Id)
	}
	updatedResult = From(&items).Where(expr.Prop("Id").NumEq(55).Predicate()).Update(expr.Prop("Id").SetBool(false).Predicate()).Collect()

	if updatedResult[0].Id != 55 {
		t.Error("Expected 55, got ", updatedResult[0].Id)
	}

	updatedResult = From(&items).Where(expr.Prop("Id").NumEq(55).Predicate()).Update(expr.Prop("Id").StrApp("yellow").Predicate()).Collect()

	if updatedResult[0].Id != 55 {
		t.Error("Expected 55, got ", updatedResult[0].Id)
	}
}

func TestBreakRuntime_InvalidStructAppend(t *testing.T) {

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
		{
			Name: "Ahmad",
			Age:  52,
			Id:   184,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 0,
		},
		{
			Name: "Reza",
			Age:  28,
			Id:   2,
			Addr: []Address{
				{City: "Karaj", Street: "Azadi", No: 8},
			},
			ParentId: 1,
		},
		{
			Name: "Dariush",
			Age:  52,
			Id:   185,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Sara",
			Age:  24,
			Id:   3,
			Addr: []Address{
				{City: "Shiraz", Street: "Chamran", No: 21},
			},
			ParentId: 1,
		},
		{
			Name: "Darvish",
			Age:  52,
			Id:   186,
			Addr: []Address{
				{City: "Tehran", Street: "Valiasr", No: 12},
			},
			ParentId: 184,
		},
		{
			Name: "Mina",
			Age:  31,
			Id:   4,
			Addr: []Address{
				{City: "Qom", Street: "Imam", No: 5},
			},
			ParentId: 0,
		},
		{
			Name: "Hossein",
			Age:  40,
			Id:   5,
			Addr: []Address{
				{City: "Mashhad", Street: "Sajjad", No: 18},
			},
			ParentId: 4,
		},
		{
			Name: "Niloofar",
			Age:  22,
			Id:   6,
			Addr: []Address{
				{City: "Isfahan", Street: "HashtBehesht", No: 33},
			},
			ParentId: 0,
		},
		{
			Name: "Amir",
			Age:  35,
			Id:   7,
			Addr: []Address{
				{City: "Qom", Street: "Bahonar", No: 9},
			},
			ParentId: 5,
		},
		{
			Name: "Fatemeh",
			Age:  27,
			Id:   8,
			Addr: []Address{
				{City: "Tehran", Street: "Kianpars", No: 44},
			},
			ParentId: 0,
		},
		{
			Name: "Mehdi",
			Age:  19,
			Id:   9,
			Addr: []Address{
				{City: "Tehran", Street: "Golha", No: 14},
			},
			ParentId: 8,
		},
		{
			Name: "Zahra",
			Age:  45,
			Id:   10,
			Addr: []Address{
				{City: "Tehran", Street: "Danesh", No: 2},
			},
			ParentId: 0,
		},
	}

	user := Sifu.Expr[User]()

	updated_result := From(&Users).Where(user.Prop("Id").NumEq(10).Predicate()).Update(user.Prop("Addr").AppStruct(ForeignAddress{
		Country: "USA",
	}).Predicate()).Collect()

	if updated_result[0].Addr[0].City != "Tehran" {
		t.Errorf("Failed to set struct")
	}

}

func TestBreakRuntime_InvalidStructSet(t *testing.T) {

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

	user := Sifu.Expr[ForeignUser]()

	updatedResult := From(&Users).Where(user.Prop("Id").NumEq(184).Predicate()).Update(user.Prop("Addr").SetStruct(ForeignAddress{
		Country: "USA",
	}).Predicate()).Collect()

	if updatedResult[0].Addr.City != "Tehran" {
		t.Errorf("Failed to set struct")
	}
}
