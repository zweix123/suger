package common

import (
	"fmt"
	"runtime"
)

// Assert 断言
func Assert(condition bool, message string) {
	if !condition {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			panic(fmt.Sprintf("%s:%d: %s", file, line, message))
		} else {
			panic(fmt.Sprintf("unknown: %s", message))
		}
	}
}
