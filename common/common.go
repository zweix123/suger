package common

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func Assert(condition bool, message string) {
	if !condition {
		_, file, line, ok := runtime.Caller(1)
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
func HandlePanic(action func(file string, line int, err any)) {
	if r := recover(); r != nil {
		_, file, line, ok := runtime.Caller(1)
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
