package functional

import (
	"strconv"
	"testing"

	"github.com/zweix123/suger/slice"
)

func TestMapSerial(t *testing.T) {
	result1 := MapSerial([]int{1, 2, 3, 4}, func(x int, _ int) string {
		return "Hello"
	})
	if len(result1) != 4 {
		t.Errorf("expected length of result1 to be 4, got %d", len(result1))
	}
	if result1[0] != "Hello" || result1[1] != "Hello" || result1[2] != "Hello" || result1[3] != "Hello" {
		t.Errorf("expected result1 to be [Hello, Hello, Hello, Hello], got %v", result1)
	}
	result2 := MapSerial([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	if len(result2) != 4 {
		t.Errorf("expected length of result2 to be 4, got %d", len(result2))
	}
	if result2[0] != "1" || result2[1] != "2" || result2[2] != "3" || result2[3] != "4" {
		t.Errorf("expected result2 to be [1, 2, 3, 4], got %v", result2)
	}
}

func TestMapParallel(t *testing.T) {
	result1 := MapParallel([]int{1, 2, 3, 4}, func(x int, _ int) string {
		return "Hello"
	})
	if len(result1) != 4 {
		t.Errorf("expected length of result1 to be 4, got %d", len(result1))
	}
	if !slice.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello"}) {
		t.Errorf("expected result1 to be [Hello, Hello, Hello, Hello], got %v", result1)
	}
	result2 := MapParallel([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	if len(result2) != 4 {
		t.Errorf("expected length of result2 to be 4, got %d", len(result2))
	}
	if !slice.Equal(result2, []string{"1", "2", "3", "4"}) {
		t.Errorf("expected result2 to be [1, 2, 3, 4], got %v", result2)
	}
}

func TestReduce(t *testing.T) {
	// TODO
}

func TestMapReduceSerial(t *testing.T) {
	srcArray := []int{1, 2, 3, 4}
	f := func(n int) []int {
		res := make([]int, 0, n)
		for i := 1; i <= n; i++ {
			res = append(res, i)
		}
		return res
	}
	result := Reduce(
		MapSerial(srcArray, func(n int, _ int) []int {
			return f(n)
		}),
		func(agg []int, item []int, _ int) []int {
			return append(agg, item...)
		},
		[]int{},
	)
	if !slice.Equal(result, []int{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}) {
		t.Errorf("expected result to be [1, 1, 2, 1, 2, 3, 1, 2, 3, 4], got %v", result)
	}
}

func TestMapReduceParallel(t *testing.T) {
	srcArray := []int{1, 2, 3, 4}
	f := func(n int) []int {
		res := make([]int, 0, n)
		for i := 1; i <= n; i++ {
			res = append(res, i)
		}
		return res
	}
	result := Reduce(
		MapParallel(srcArray, func(n int, _ int) []int {
			return f(n)
		}),
		func(agg []int, item []int, _ int) []int {
			return append(agg, item...)
		},
		[]int{},
	)
	if !slice.Equal(result, []int{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}) {
		t.Errorf("expected result to be [1, 1, 2, 1, 2, 3, 1, 2, 3, 4], got %v", result)
	}
}

func TestFilter(t *testing.T) {
	r1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	if len(r1) != 2 {
		t.Errorf("expected length of r1 to be 2, got %d", len(r1))
	}
	if !slice.Equal(r1, []int{2, 4}) {
		t.Errorf("expected r1 to be [2, 4], got %v", r1)
	}

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})
	if len(r2) != 2 {
		t.Errorf("expected length of r2 to be 2, got %d", len(r2))
	}
	if !slice.Equal(r2, []string{"foo", "bar"}) {
		t.Errorf("expected r2 to be [foo, bar], got %v", r2)
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Filter(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	if _, ok := interface{}(nonempty).(myStrings); !ok {
		t.Errorf("type preserved: expected nonempty to be of type []string, got %T", nonempty)
	}
}
