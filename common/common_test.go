package common

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func TestCurrentGetPanicInfo(t *testing.T) {
	var file string
	var line int
	var ok bool
	func() {
		file, line, ok = currentGetPanicInfo(1)
	}()
	if !strings.HasSuffix(file, "common_test.go") || ok != true {
		t.Errorf("Expected file to be common_test.go, ok to be true, got %q, %d, %t", file, line, ok)
	}
}

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

func TestAssertMock(t *testing.T) {
	// 默认的
	currentGetPanicInfo = func(skip int) (file string, line int, ok bool) {
		_, file, line, ok = runtime.Caller(skip + 1)
		return
	}
	defer func() {
		r := recover()
		if r == nil { // no panic
			t.Errorf("Expected panic, but did not occur")
		}
		s, ok := r.(string)
		if !ok {
			t.Errorf("Expected panic message to be string, got %T", r)
		}
		if !strings.Contains(s, "common_test.go") || !strings.Contains(s, "assert_message") {
			t.Errorf("Expected panic message to contain common_test.go and assert_message, got %q", s)
		}
	}()
	Assert(false, "assert_message") // should panic

	// 返回真的
	currentGetPanicInfo = func(skip int) (file string, line int, ok bool) {
		return "mock.go", 0, true
	}
	defer func() {
		r := recover()
		if r == nil { // no panic
			t.Errorf("Expected panic, but did not occur")
		}
		s, ok := r.(string)
		if !ok {
			t.Errorf("Expected panic message to be string, got %T", r)
		}
		if s != "mock.go:0: assert_message" {
			t.Errorf("Expected panic message to be mock.go:0: assert_message, got %q", s)
		}
	}()
	Assert(false, "assert_message") // should panic

	// 返回假的
	currentGetPanicInfo = func(skip int) (file string, line int, ok bool) {
		return "", 0, false
	}
	defer func() {
		r := recover()
		if r == nil { // no panic
			t.Errorf("Expected panic, but did not occur")
		}
		s, ok := r.(string)
		if !ok {
			t.Errorf("Expected panic message to be string, got %T", r)
		}
		if s != "unknown: assert_message" {
			t.Errorf("Expected panic message to be unknown: assert_message, got %q", s)
		}
	}()
	Assert(false, "assert_message") // should panic
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

func TestHandlePanicMock(t *testing.T) {
	currentGetPanicInfo = func(skip int) (file string, line int, ok bool) {
		_, file, line, ok = runtime.Caller(skip + 1)
		return
	}
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() { // nolint
		defer wg1.Done()
		defer HandlePanic(func(file string, _ int, err any) {
			if !strings.HasSuffix(file, "common_test.go") {
				t.Errorf("[1]Expected file to be common_test.go, got %q", file)
			}
			if err != "[1]TestHandlePanicMock Panic" {
				t.Errorf("[1]Expected err to be TestHandlePanicMock Panic1, got %q", err)
			}
		})
		panic("[1]TestHandlePanicMock Panic") // nolint
	}()
	wg1.Wait()

	// 返回假的
	currentGetPanicInfo = func(_ int) (file string, line int, ok bool) {
		return "", 0, false
	}
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() { // nolint
		defer wg2.Done()
		defer HandlePanic(func(file string, line int, err any) {
			if file != "unknown" {
				t.Errorf("[2]Expected file to be unknown, got %q", file)
			}
			if line != 0 {
				t.Errorf("[2]Expected line to be 0, got %d", line)
			}
			if err != "[2]TestHandlePanicMock Panic" {
				t.Errorf("[2]Expected err to be [2]TestHandlePanicMock Panic, got %q", err)
			}
		})
		panic("[2]TestHandlePanicMock Panic") // nolint
	}()
	wg2.Wait()
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
