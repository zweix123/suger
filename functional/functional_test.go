package functional

import (
	"reflect"
	"strconv"
	"testing"
)

func TestTimes(t *testing.T) {
	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	if len(result1) != 3 || !reflect.DeepEqual(result1, []string{"0", "1", "2"}) {
		t.Errorf("expected result1 to be [0, 1, 2], got %v", result1)
	}
}

func TestAll(t *testing.T) {
	if !All([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
		return x > 0
	}) {
		t.Errorf("expected All to be true")
	}
	if All([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
		return x > 3
	}) {
		t.Errorf("expected All to be false")
	}
}

func TestAny(t *testing.T) {
	if !Any([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
		return x > 3
	}) {
		t.Errorf("expected Any to be true")
	}
	if Any([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
		return x > 5
	}) {
		t.Errorf("expected Any to be false")
	}
}
