package functional

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		collection []int
		iteratee   func(item int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, func(item int) int {
			return item * 2
		}}, []int{2, 4, 6, 8, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
