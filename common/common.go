package common

import (
	"fmt"
	"runtime"
)

func Assert(condition bool, message string) {
	if !condition {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "???"
			line = 0
		}
		panic(fmt.Sprintf("%s:%d: %s", file, line, message))
	}
}
