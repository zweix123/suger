package common

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/zweix123/suger/test"
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

func parse(output string) (file string, line int, message string, err error) {
	// file:line: message
	parts := strings.Split(output, ":")
	if len(parts) != 3 {
		return "", 0, output, fmt.Errorf("invalid panic message format: %q", output)
	}
	line, err = strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, output, fmt.Errorf("invalid line number: %q", parts[1])
	}
	parts[2] = strings.TrimSpace(parts[2])
	return parts[0], line, parts[2], nil
}

func TestAssertLine(t *testing.T) {
	msg := "TestAssertLine Assert Flag"
	defer func() {
		r := recover()
		if r == nil { // no panic
			t.Errorf("Expected panic, but did not occur")
		} else {
			file, line, message, err := parse(fmt.Sprintf("%v", r))
			if err != nil {
				t.Errorf("Unexpected panic: %v", err)
			}
			if !strings.HasSuffix(file, "common/common_test.go") {
				t.Errorf("Expected file to be %q, but got %q", "common/common_test.go", file)
			}
			if line != 67 {
				t.Errorf("Expected line to be %d, but got %d", 66, line)
			}
			if message != msg {
				t.Errorf("Expected message to be %q, but got %q", msg, message)
			}
		}
	}()
	Assert(false, msg)
}

func TestHandlePanic(t *testing.T) {
	t.Run("No Panic", func(t *testing.T) {
		actionCalled := false
		HandlePanic(func(_ string, _ int, _ any, _ []byte) {
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
			defer HandlePanic(func(_ string, _ int, _ any, _ []byte) {
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
			defer HandlePanic(func(_ string, _ int, _ any, _ []byte) {
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

func TestHandlePanicOutput1(t *testing.T) {
	msg := "TestHandlePanicOutput1 Panic"

	buf := test.HandleStdout(func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() { // nolint
			defer wg.Done()
			defer HandlePanic(func(file string, line int, err any, _ []byte) {
				fmt.Printf("%s:%d: %v", file, line, err)
			})
			panic(msg) // nolint
		}()
		wg.Wait()
	})

	file, line, message, err := parse(buf.String())
	if err != nil {
		t.Errorf("Unexpected panic: %v", err)
	}
	if !strings.HasSuffix(file, "common/common_test.go") {
		t.Errorf("Expected file to be %q, but got %q", "common/common_test.go", file)
	}
	if line != 131 {
		t.Errorf("Expected line to be %d, but got %d", 131, line)
	}
	if message != msg {
		t.Errorf("Expected message to be %q, but got %q", msg, message)
	}
}

func TestHandlePanicOutput2(t *testing.T) {
	msg := "TestHandlePanicOutput2 Panic"

	buf := test.HandleStdout(func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() { // nolint
			defer wg.Done()
			defer HandlePanic(func(_ string, _ int, _ any, stack []byte) {
				fmt.Printf("%s", string(stack))
			})
			panic(msg) // nolint
		}()
		wg.Wait()
	})

	lines := strings.Split(buf.String(), "\n")
	check := func() bool {
		if len(lines) != 12 {
			fmt.Println(len(lines))
			return false
		}
		if !strings.HasSuffix(lines[0], "[running]:") {
			fmt.Printf("lines[0]: %q\n", lines[0])
			return false
		}
		if !strings.Contains(lines[3], "/suger/common.HandlePanic") {
			fmt.Printf("lines[3]: %q\n", lines[3])
			return false
		}
		if !strings.Contains(lines[4], "/suger/common/common.go:36") {
			fmt.Printf("lines[4]: %q\n", lines[4])
			return false
		}
		if !strings.Contains(lines[7], "/suger/common.TestHandlePanicOutput2.func1.1()") {
			fmt.Printf("lines[7]: %q\n", lines[7])
			return false
		}
		return true
	}()
	if !check {
		for idx, line := range lines {
			t.Error(idx, line)
		}
	}
}

func TestLogStr(t *testing.T) {
	// common test
	if MustJsonMarshal(1) != "1" {
		t.Errorf("LogStr(1) = %q, want %q", MustJsonMarshal(1), "1")
	}
	if MustJsonMarshal(3.14) != "3.14" {
		t.Errorf("LogStr(3.14) = %q, want %q", MustJsonMarshal(3.14), "3.14")
	}
	if MustJsonMarshal("hello") != "\"hello\"" {
		t.Errorf("LogStr(\"hello\") = %q, want %q", MustJsonMarshal("hello"), "\"hello\"")
	}
	if MustJsonMarshal([]int{1, 2, 3}) != "[1,2,3]" {
		t.Errorf("LogStr([1,2,3]) = %q, want %q", MustJsonMarshal([]int{1, 2, 3}), "[1,2,3]")
	}
	if MustJsonMarshal(map[string]int{"a": 1, "b": 2}) != "{\"a\":1,\"b\":2}" {
		t.Errorf("LogStr(map[string]int{\"a\":1,\"b\":2}) = %q, want %q", MustJsonMarshal(map[string]int{"a": 1, "b": 2}), "{\"a\":1,\"b\":2}")
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
	if MustJsonMarshal(Common{A: 1, B: 2.0, C: "three", D: []int{4, 5, 6}, E: map[string]int{"seven": 7, "eight": 8}}) != "{\"A\":1,\"B\":2,\"C\":\"three\",\"D\":[4,5,6],\"E\":{\"eight\":8,\"seven\":7}}" {
		t.Errorf("LogStr(Common{A: 1, B: 2.0, C: \"three\", D: [4,5,6], E: map[seven:7,eight:8]}) = %q, want %q", MustJsonMarshal(Common{A: 1, B: 2.0, C: "three", D: []int{4, 5, 6}, E: map[string]int{"seven": 7, "eight": 8}}), "{\"A\":1,\"B\":2,\"C\":\"three\",\"D\":[4,5,6],\"E\":{\"eight\":8,\"seven\":7}}")
	}
	// 2. private
	type Private struct {
		a int
		b float64
		c string
		d []int
		e map[string]int
	}
	if MustJsonMarshal(Private{a: 1, b: 2.0, c: "three", d: []int{4, 5, 6}, e: map[string]int{"seven": 7, "eight": 8}}) != "{}" { // Non exported fields are not displayed
		t.Errorf("LogStr(Private{a: 1, b: 2.0, c: \"three\", d: [4,5,6], e: map[seven:7,eight:8]}) = %q, want %q", MustJsonMarshal(Private{a: 1, b: 2.0, c: "three", d: []int{4, 5, 6}, e: map[string]int{"seven": 7, "eight": 8}}), "{}")
	}

	// panic type
	// 1. circular reference
	type Node struct {
		Next *Node
	}
	n := &Node{}
	n.Next = n
	if !strings.Contains(MustJsonMarshal(n), "is unsupported type") {
		t.Errorf("LogStr(&Node{Next: n}) = %q, want %q", MustJsonMarshal(n), "is unsupported type")
	}
	// 2. unsupported type
	if !strings.Contains(MustJsonMarshal(make(chan int)), "is unsupported type") {
		t.Errorf("LogStr(make(chan int)) = %q, want %q", MustJsonMarshal(make(chan int)), "is unsupported type")
	}
}

func TestZero(t *testing.T) {
	// 1. base function
	// 1.1 built-in type
	if Zero[bool]() != false {
		t.Errorf("Zero[bool]() = %t, want %t", Zero[bool](), false)
	}
	if Zero[int]() != 0 {
		t.Errorf("Zero[int]() = %d, want %d", Zero[int](), 0)
	}
	if Zero[float64]() != 0.0 {
		t.Errorf("Zero[float64]() = %f, want %f", Zero[float64](), 0.0)
	}
	if Zero[string]() != "" {
		t.Errorf("Zero[string]() = %q, want %q", Zero[string](), "")
	}
	// 1.2 custom type
	type Custom struct {
		A int
		B string
	}
	if Zero[Custom]().A != 0 {
		t.Errorf("Zero[Custom]().A = %d, want %d", Zero[Custom]().A, 0)
	}
	if Zero[Custom]().B != "" {
		t.Errorf("Zero[Custom]().B = %q, want %q", Zero[Custom]().B, "")
	}
	// 2. idempotent
	first := Zero[Custom]()
	second := Zero[Custom]()
	if first != second {
		t.Errorf("Zero[Custom]() = %v, want %v", first, second)
	}
}
