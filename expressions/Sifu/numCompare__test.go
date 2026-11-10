package Sifu

import (
	"fmt"
	"testing"
)

func TestNumCompareFloat64(t *testing.T) {

	type args struct {
		id float64
	}
	tests := []args{
		{id: 0},
		{id: 12.2},
		{id: 2.2},
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, 12.1).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 12.1)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 12.1)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 12.1)
			}
		}
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, float64(2)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, float64(2))
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, float64(2))
			}
		} else if ix == 2 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, float64(2))
			}
		}
	}

}

func TestNumCompareFloat32(t *testing.T) {

	type args struct {
		id float32
	}
	tests := []args{
		{id: 0},
		{id: 12.2},
		{id: 2.2},
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, float32(12.1)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 12.1)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 12.1)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 12.1)
			}
		}
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, float32(2)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, float32(2))
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, float32(2))
			}
		} else if ix == 2 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, float32(2))
			}
		}
	}

}

func TestNumCompareint32(t *testing.T) {

	type args struct {
		id int32
	}
	tests := []args{
		{id: 0},
		{id: 12},
		{id: 1},
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, int32(3)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 3)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 3)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 3)
			}
		}
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, int32(2)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 2)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 2)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 2)
			}
		}
	}

}

func TestNumCompareUnit32(t *testing.T) {

	type args struct {
		id uint32
	}
	tests := []args{
		{id: 0},
		{id: 12},
		{id: 1},
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, uint32(3)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 3)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 3)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 3)
			}
		}
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, uint32(2)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 2)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 2)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 2)
			}
		}
	}

}

func TestNumCompareUnit8(t *testing.T) {

	type args struct {
		id uint8
	}
	tests := []args{
		{id: 0},
		{id: 12},
		{id: 1},
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, uint8(3)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 3)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 3)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 3)
			}
		}
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").numcmp(1, uint8(2)).Predicate()

		result := fnc(tt)

		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 2)
			}
		} else if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 2)
			}
		} else if ix == 2 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 2)
			}
		}
	}

}

func TestNumSmallerOrEqual(t *testing.T) {
	type args struct {
		id uint8
	}
	tests := []args{
		{id: 0},
		{id: 12},
		{id: 1},
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").NumSmallerOrEq(uint8(1)).Predicate()

		result := fnc(tt)

		fmt.Printf("comparing %d is smaller or eq to %d \n", tt.id, 1)
		fmt.Println(result)
		if ix == 0 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 1)
			}
		}

		if ix == 1 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 1)
			}
		}
		if ix == 2 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 1)
			}
		}

	}

}

func TestNumBiggerOrEqual(t *testing.T) {
	type args struct {
		id uint8
	}
	tests := []args{
		{id: 0},
		{id: 12},
		{id: 1},
	}

	for ix, tt := range tests {

		fnc := Expr[args]().Prop("id").NumBiggerOrEq(uint8(1)).Predicate()

		result := fnc(tt)

		fmt.Printf("comparing %d is smaller or eq to %d \n", tt.id, 1)
		fmt.Println(result)
		if ix == 0 {
			if result == true {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 1)
			}
		}

		if ix == 1 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return false", tt.id, 1)
			}
		}
		if ix == 2 {
			if result == false {
				t.Errorf("numCompare(%v, %v) should return true", tt.id, 1)
			}
		}

	}
}
