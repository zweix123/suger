package json

import (
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	err := Save[int]("test.json", []int{1, 2, 3})
	if err != nil {
		t.Errorf("Save() error = %v", err)
	}
}

func TestLoad(t *testing.T) {
	got, err := Load[int]("test.json")
	if err != nil {
		t.Errorf("Load() error = %v", err)
	}
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Load() = %v, want %v", got, []int{1, 2, 3})
	}
}
