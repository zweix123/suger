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

func TestChunk(t *testing.T) {
	type args struct {
		src  []int
		size int
	}
	type wants struct {
		result [][]int
		err    error
	}
	tests := []struct {
		name string
		args args
		want wants
	}{
		{"nil", args{nil, 2}, wants{[][]int{}, nil}},
		{"empty", args{[]int{}, 2}, wants{[][]int{}, nil}},
		{"one", args{[]int{1}, 2}, wants{[][]int{{1}}, nil}},
		{"two", args{[]int{1, 2}, 2}, wants{[][]int{{1, 2}}, nil}},
		{"three", args{[]int{1, 2, 3}, 2}, wants{[][]int{{1, 2}, {3}}, nil}},
		{"zero", args{[]int{1, 2, 3}, 0}, wants{nil, ChunkErr}},
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
