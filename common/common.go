package common

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// Assert is used to assert a condition, if the condition is false, it will panic.
// The panic message will contain the file and line number of the assertion, pointing to the Assert function position.
func Assert(condition bool, message string) {
	if !condition {
		_, file, line, ok := runtime.Caller(1) // situation of caller of Assert
		if ok {
			panic(fmt.Sprintf("%s:%d: %s", file, line, message)) // nolint
		} else {
			panic(fmt.Sprintf("unknown: %s", message)) // nolint
		}
	}
}

// HandlePanic is mainly used to handle panic at the beginning postion of a goroutine,
// ensuring that the panic of the goroutine does not affect other goroutines;
// It is must be called by defer.
// The action function will be called with the file, line number and error information of the panic,
// pointing to the panic position.
func HandlePanic(action func(file string, line int, err any, stack []byte)) {
	if r := recover(); r != nil {
		_, file, line, ok := runtime.Caller(2) // situation of goroutine panic
		if !ok {
			file = "unknown"
			line = 0
		}

		buf := make([]byte, 1024)
		n := runtime.Stack(buf, false)

		action(file, line, r, buf[:n])
	}
}

func MustJsonMarshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%T is unsupported type: %v", v, v)
	}
	return string(bytes)
}

func Zero[T any]() (zero T) {
	return
}
