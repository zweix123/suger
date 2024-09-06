package common

import (
	"testing"
)

func TestAssertPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but did not occur")
		}
	}()
	Assert(false, "false")
}

func TestAssertNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Unexpected panic: %v", r)
		}
	}()
	Assert(true, "true")
}
