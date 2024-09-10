package common

import (
	"sync"
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

func TestHandlePanic(t *testing.T) {
	// 测试正常情况
	t.Run("No Panic", func(t *testing.T) {
		actionCalled := false
		HandlePanic(func() {
			actionCalled = true
		})
		if actionCalled {
			t.Error("Action should not be called when no panic occurs")
		}
	})

	// 测试 panic 情况
	t.Run("With Panic", func(t *testing.T) {
		actionCalled := false
		// recovered := false

		func() {
			defer HandlePanic(func() {
				actionCalled = true
			})
			panic("Test panic")
		}()

		if !actionCalled {
			t.Error("Action should be called when panic occurs")
		}
		// if !recovered {
		// 	t.Error("Panic should be recovered")
		// }
	})

	// 测试在 goroutine 中的使用
	t.Run("In Goroutine", func(t *testing.T) {
		var wg sync.WaitGroup
		actionCalled := false

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer HandlePanic(func() {
				actionCalled = true
			})
			panic("Goroutine panic")
		}()

		wg.Wait()
		if !actionCalled {
			t.Error("Action should be called when panic occurs in goroutine")
		}
	})
}
