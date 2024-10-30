package common

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestAssertPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil { // no panic
			t.Errorf("Expected panic, but did not occur")
		}
	}()
	Assert(false, "") // should panic
}

func TestAssertNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil { // happen panic
			t.Errorf("Unexpected panic: %v", r)
		}
	}()
	Assert(true, "") // should not panic
}

func TestHandlePanic(t *testing.T) {
	t.Run("No Panic", func(t *testing.T) {
		actionCalled := false

		HandlePanic(func(_ string, _ int, _ any) {
			actionCalled = true
		})
		if actionCalled {
			t.Error("Action should not be called when no panic occurs")
		}
	})

	t.Run("No Panic In Goroutine", func(t *testing.T) {
		actionCalled := false

		var wg sync.WaitGroup
		wg.Add(1)
		go func() { // nolint
			defer wg.Done()
			defer HandlePanic(func(_ string, _ int, _ any) {
				actionCalled = true
			})
			panic("panic") // nolint
		}()
		wg.Wait()

		if !actionCalled {
			t.Error("Action should be called when panic occurs in goroutine")
		}
	})

	t.Run("Happen Panic", func(t *testing.T) {
		actionCalled := false

		var wg sync.WaitGroup
		wg.Add(1)
		func() {
			defer wg.Done()
			defer HandlePanic(func(_ string, _ int, _ any) {
				actionCalled = true
			})
			panic("panic") // nolint
		}()
		wg.Wait()

		if !actionCalled {
			t.Error("Action should be called when panic occurs")
		}
	})
}

func TestHandlePanicOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe failed: %v", err)
	}
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // nolint
		defer wg.Done()
		defer HandlePanic(func(file string, line int, err any) {
			fmt.Printf("%s:%d: %v", file, line, err)
		})
		panic("TestHandlePanicOutput Panic") // nolint
	}()
	wg.Wait()

	w.Close() // nolint
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Errorf("ReadFrom failed: %v", err)
	}
	output := buf.String()

	expected := "TestHandlePanicOutput Panic"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but got %q", expected, output)
	}
}
