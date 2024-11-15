package slice

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		l []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil, 1}, false},
		{"empty", args{[]int{}, 1}, false},
		{"not found", args{[]int{1, 2, 3}, 4}, false},
		{"found", args{[]int{1, 2, 3}, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.l, tt.args.e); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil}, []int{}},
		{"empty", args{[]int{}}, []int{}},
		{"one", args{[]int{1}}, []int{1}},
		{"two", args{[]int{1, 2}}, []int{2, 1}},
		{"three", args{[]int{1, 2, 3}}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniq(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"nil", args{nil}, []int{}},
		{"empty", args{[]int{}}, []int{}},
		{"one", args{[]int{1}}, []int{1}},
		{"two", args{[]int{1, 2}}, []int{1, 2}},
		{"three", args{[]int{1, 2, 3}}, []int{1, 2, 3}},
		{"duplicate1", args{[]int{1, 1, 2, 2, 3, 3}}, []int{1, 2, 3}},
		{"duplicate2", args{[]int{1, 2, 3, 3, 2, 1}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uniq(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
