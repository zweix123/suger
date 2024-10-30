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

func TestLogStr(t *testing.T) {
	// common test
	if LogStr(1) != "1" {
		t.Errorf("LogStr(1) = %q, want %q", LogStr(1), "1")
	}
	if LogStr(3.14) != "3.14" {
		t.Errorf("LogStr(3.14) = %q, want %q", LogStr(3.14), "3.14")
	}
	if LogStr("hello") != "\"hello\"" {
		t.Errorf("LogStr(\"hello\") = %q, want %q", LogStr("hello"), "\"hello\"")
	}
	if LogStr([]int{1, 2, 3}) != "[1,2,3]" {
		t.Errorf("LogStr([1,2,3]) = %q, want %q", LogStr([]int{1, 2, 3}), "[1,2,3]")
	}
	if LogStr(map[string]int{"a": 1, "b": 2}) != "{\"a\":1,\"b\":2}" {
		t.Errorf("LogStr(map[string]int{\"a\":1,\"b\":2}) = %q, want %q", LogStr(map[string]int{"a": 1, "b": 2}), "{\"a\":1,\"b\":2}")
	}

	// custom type
	// 1. common
	type Common struct {
		A int
		B float64
		C string
		D []int
		E map[string]int
	}
	if LogStr(Common{A: 1, B: 2.0, C: "three", D: []int{4, 5, 6}, E: map[string]int{"seven": 7, "eight": 8}}) != "{\"A\":1,\"B\":2,\"C\":\"three\",\"D\":[4,5,6],\"E\":{\"eight\":8,\"seven\":7}}" {
		t.Errorf("LogStr(Common{A: 1, B: 2.0, C: \"three\", D: [4,5,6], E: map[seven:7,eight:8]}) = %q, want %q", LogStr(Common{A: 1, B: 2.0, C: "three", D: []int{4, 5, 6}, E: map[string]int{"seven": 7, "eight": 8}}), "{\"A\":1,\"B\":2,\"C\":\"three\",\"D\":[4,5,6],\"E\":{\"eight\":8,\"seven\":7}}")
	}
	// 2. private
	type Private struct {
		a int
		b float64
		c string
		d []int
		e map[string]int
	}
	if LogStr(Private{a: 1, b: 2.0, c: "three", d: []int{4, 5, 6}, e: map[string]int{"seven": 7, "eight": 8}}) != "{}" { // Non exported fields are not displayed
		t.Errorf("LogStr(Private{a: 1, b: 2.0, c: \"three\", d: [4,5,6], e: map[seven:7,eight:8]}) = %q, want %q", LogStr(Private{a: 1, b: 2.0, c: "three", d: []int{4, 5, 6}, e: map[string]int{"seven": 7, "eight": 8}}), "{}")
	}

	// panic type
	// 1. circular reference
	type Node struct {
		Next *Node
	}
	n := &Node{}
	n.Next = n
	if !strings.Contains(LogStr(n), "is unsupported type") {
		t.Errorf("LogStr(&Node{Next: n}) = %q, want %q", LogStr(n), "is unsupported type")
	}
	// 2. unsupported type
	if !strings.Contains(LogStr(make(chan int)), "is unsupported type") {
		t.Errorf("LogStr(make(chan int)) = %q, want %q", LogStr(make(chan int)), "is unsupported type")
	}
}
