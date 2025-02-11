package slice

/*
How to design test samples from what aspects?
1. base function test
2. type test
3. side effect
4. abnormal condition
*/

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	assert.Equal(t, []string{"0", "1", "2"}, result1)
}

func TestAll(t *testing.T) {
	assert.True(
		t,
		All([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
			return x > 0
		}),
	)
	assert.False(
		t,
		All([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
			return x > 3
		}),
	)
}

func TestAny(t *testing.T) {
	assert.True(
		t,
		Any([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
			return x > 3
		}),
	)
	assert.False(
		t,
		Any([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
			return x > 5
		}),
	)
}

func TestFilter(t *testing.T) {
	r1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	assert.Equal(t, []int{2, 4}, r1)

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})
	assert.Equal(t, []string{"foo", "bar"}, r2)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Filter(allStrings, func(x string, _ int) bool {
		return len(x) > 0
	})
	assert.IsType(t, myStrings{}, nonempty)
}

func TestChunk(t *testing.T) {
	type args struct {
		src  []int
		size int
	}
	type wants struct {
		result [][]int
	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{"nil", args{nil, 2}, wants{[][]int{}}},
		{"empty", args{[]int{}, 2}, wants{[][]int{}}},
		{"one", args{[]int{1}, 2}, wants{[][]int{{1}}}},
		{"two", args{[]int{1, 2}, 2}, wants{[][]int{{1, 2}}}},
		{"three", args{[]int{1, 2, 3}, 2}, wants{[][]int{{1, 2}, {3}}}},
		{"zero", args{[]int{1, 2, 3}, 0}, wants{nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.args.src, tt.args.size)
			if !reflect.DeepEqual(got, tt.want.result) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want.result)
			}
		})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Chunk(allStrings, 2)
	assert.IsType(t, allStrings, nonempty[0])

	// appending to a chunk should not affect original array
	originalArray := []int{0, 1, 2, 3, 4, 5}
	result := Chunk(originalArray, 2)
	result[0] = append(result[0], 6)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, originalArray)
}

func TestFlatten(t *testing.T) {
	result1 := Flatten([][]int{{0, 1}, {2, 3, 4, 5}})
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, result1)

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Flatten([]myStrings{allStrings})
	assert.IsType(t, allStrings, nonempty)
}
