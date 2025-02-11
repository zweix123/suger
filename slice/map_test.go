package slice

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMapSerial(t *testing.T) {
	result1 := MapSerial([]int{1, 2, 3, 4}, func(_ int, _ int) string {
		return "Hello"
	})
	assert.Equal(t, 4, len(result1))
	assert.Equal(t, []string{"Hello", "Hello", "Hello", "Hello"}, result1)
	result2 := MapSerial([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	assert.Equal(t, []string{"1", "2", "3", "4"}, result2)
}

func TestMapParallel(t *testing.T) {
	result1 := MapParallel([]int{1, 2, 3, 4}, func(_ int, _ int) string {
		return "Hello"
	})
	assert.Equal(t, 4, len(result1))
	assert.Equal(t, []string{"Hello", "Hello", "Hello", "Hello"}, result1)
	result2 := MapParallel([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	assert.Equal(t, []string{"1", "2", "3", "4"}, result2)
}

func TestMapParallelWithGoroutineUpperLimit(t *testing.T) {
	result1 := MapParallelWithGoroutineUpperLimit([]int{1, 2, 3, 4}, func(_ int, _ int) string {
		return "Hello"
	}, 2)
	assert.Equal(t, 4, len(result1))
	assert.Equal(t, []string{"Hello", "Hello", "Hello", "Hello"}, result1)
	result2 := MapParallelWithGoroutineUpperLimit([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	}, 2)
	assert.Equal(t, []string{"1", "2", "3", "4"}, result2)
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
	checkMapResult := func(t *testing.T, r []string) {
		assert.Equal(t, len(r), len(mapSlice))
		assert.Equal(t, r, mapResult)
	}

	t.Run("param goroutineNum", func(t *testing.T) {
		for i := 0; i <= len(mapSlice); i++ {
			r := MapParallelWithGoroutineUpperLimit(mapSlice, mapFunc, i)
			checkMapResult(t, r)
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
					panic("panic") // nolint
				}
				return mapFunc(item, idx)
			}, i)
			checkMapResult(t, r)
		}
	})

	t.Run("latency", func(t *testing.T) {
		for i := 1; i <= len(mapSlice); i++ {
			now := time.Now()
			r := MapParallelWithGoroutineUpperLimit(mapSlice, func(item int, idx int) string {
				time.Sleep(time.Second)
				return mapFunc(item, idx)
			}, i)
			checkMapResult(t, r)
			t.Logf("goroutineNum: %d, latency: %s", i, time.Since(now))
		}
	})
}
