package common

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type getPanicInfo func(skip int) (file string, line int, ok bool)

var currentGetPanicInfo getPanicInfo = func(skip int) (file string, line int, ok bool) {
	_, file, line, ok = runtime.Caller(skip + 1)
	return
}

// Assert is used to assert a condition, if the condition is false, it will panic.
// The panic message will contain the file and line number of the assertion, pointing to the Assert function position.
func Assert(condition bool, message string) {
	if !condition {
		file, line, ok := currentGetPanicInfo(1)
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
func HandlePanic(action func(file string, line int, err any)) {
	if r := recover(); r != nil {
		file, line, ok := currentGetPanicInfo(2)
		if !ok {
			file = "unknown"
			line = 0
		}
		action(file, line, r)
	}
}

func LogStr(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%T is unsupported type: %v", v, v)
	}
	return string(bytes)
}
