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
	"testing"
)

func TestFilter(t *testing.T) {
	r1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	if len(r1) != 2 {
		t.Errorf("expected length of r1 to be 2, got %d", len(r1))
	}
	if !reflect.DeepEqual(r1, []int{2, 4}) {
		t.Errorf("expected r1 to be [2, 4], got %v", r1)
	}

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})
	if len(r2) != 2 {
		t.Errorf("expected length of r2 to be 2, got %d", len(r2))
	}
	if !reflect.DeepEqual(r2, []string{"foo", "bar"}) {
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
	if reflect.TypeOf(nonempty[0]) != reflect.TypeOf(allStrings) {
		t.Errorf("type not preserved")
	}

	// appending to a chunk should not affect original array
	originalArray := []int{0, 1, 2, 3, 4, 5}
	result := Chunk(originalArray, 2)
	result[0] = append(result[0], 6)
	if !reflect.DeepEqual(originalArray, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("original array is affected")
	}
}

func TestFlatten(t *testing.T) {
	result1 := Flatten([][]int{{0, 1}, {2, 3, 4, 5}})
	if !reflect.DeepEqual(result1, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("Flatten() = %v, want %v", result1, []int{0, 1, 2, 3, 4, 5})
	}

	type myStrings []string
	allStrings := myStrings{"", "foo", "bar"}
	nonempty := Flatten([]myStrings{allStrings})
	if reflect.TypeOf(nonempty) != reflect.TypeOf(allStrings) {
		t.Errorf("type not preserved")
	}
}
