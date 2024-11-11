package functional

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestMapSerial(t *testing.T) {
	result1 := MapSerial([]int{1, 2, 3, 4}, func(_ int, _ int) string {
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
	result1 := MapParallel([]int{1, 2, 3, 4}, func(_ int, _ int) string {
		return "Hello"
	})
	if len(result1) != 4 {
		t.Errorf("expected length of result1 to be 4, got %d", len(result1))
	}
	if !reflect.DeepEqual(result1, []string{"Hello", "Hello", "Hello", "Hello"}) {
		t.Errorf("expected result1 to be [Hello, Hello, Hello, Hello], got %v", result1)
	}
	result2 := MapParallel([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	if len(result2) != 4 {
		t.Errorf("expected length of result2 to be 4, got %d", len(result2))
	}
	if !reflect.DeepEqual(result2, []string{"1", "2", "3", "4"}) {
		t.Errorf("expected result2 to be [1, 2, 3, 4], got %v", result2)
	}
}

func TestMapParallelWithGoroutineUpperLimit(t *testing.T) {
	result1 := MapParallelWithGoroutineUpperLimit([]int{1, 2, 3, 4}, func(_ int, _ int) string {
		return "Hello"
	}, 2)
	if len(result1) != 4 {
		t.Errorf("expected length of result1 to be 4, got %d", len(result1))
	}
	if !reflect.DeepEqual(result1, []string{"Hello", "Hello", "Hello", "Hello"}) {
		t.Errorf("expected result1 to be [Hello, Hello, Hello, Hello], got %v", result1)
	}
	result2 := MapParallelWithGoroutineUpperLimit([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	}, 2)
	if len(result2) != 4 {
		t.Errorf("expected length of result2 to be 4, got %d", len(result2))
	}
	if !reflect.DeepEqual(result2, []string{"1", "2", "3", "4"}) {
		t.Errorf("expected result2 to be [1, 2, 3, 4], got %v", result2)
	}
}

func TestMapParallelWithGoroutineUpperLimitConcurrentSecurity(t *testing.T) {
	mapSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mapFunc := func(n int, _ int) string {
		return strconv.FormatInt(int64(n), 10)
	}
	mapResult := make([]string, len(mapSlice))
	for i, item := range mapSlice {
		mapResult[i] = mapFunc(item, i)
	}
	checkMapResult := func(r []string) {
		if len(r) != len(mapSlice) {
			t.Errorf("expected length of r to be %d, got %d", len(mapSlice), len(r))
		}
		if !reflect.DeepEqual(r, mapResult) {
			t.Errorf("expected r to be %v, got %v", mapResult, r)
		}
	}

	t.Run("param goroutineNum", func(t *testing.T) {
		for i := 0; i <= len(mapSlice); i++ {
			r := MapParallelWithGoroutineUpperLimit(mapSlice, mapFunc, i)
			checkMapResult(r)
		}
	})

	t.Run("panic", func(t *testing.T) {
		mapSlice[3] = 0
		defer func() {
			mapSlice[3] = 3
		}()
		mapResult[3] = ""
		defer func() {
			mapResult[3] = "3"
		}()

		for i := 0; i <= len(mapSlice); i++ {
			r := MapParallelWithGoroutineUpperLimit(mapSlice, func(item int, idx int) string {
				if idx == 3 {
					panic("panic")
				}
				return mapFunc(item, idx)
			}, i)
			checkMapResult(r)
		}
	})

	t.Run("latency", func(t *testing.T) {
		for i := 1; i <= len(mapSlice); i++ {
			now := time.Now()
			r := MapParallelWithGoroutineUpperLimit(mapSlice, func(item int, idx int) string {
				time.Sleep(time.Second)
				return mapFunc(item, idx)
			}, i)
			checkMapResult(r)
			t.Logf("goroutineNum: %d, latency: %s", i, time.Since(now))
		}
	})
}
