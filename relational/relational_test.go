package relational

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	if len(result1) != 3 {
		t.Errorf("expected result1 to be 3, got %v", len(result1))
	}
	if !reflect.DeepEqual(result1, map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	}) {
		t.Errorf("expected result1 to be %v, got %v", map[int][]int{
			0: {0, 3},
			1: {1, 4},
			2: {2, 5},
		}, result1)
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := GroupBy(allStrings, func(_ string) int {
		return 42
	})
	if reflect.TypeOf(nonempty[42]).String() != "relational.myStrings" {
		t.Errorf("expected type of nonempty[42] to be relational.myStrings, got %v", reflect.TypeOf(nonempty[42]))
	}
}
