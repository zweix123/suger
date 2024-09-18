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

// HandlePanic 处理panic, 主要用于goroutine的开始, 保证该goroutine的panic不会影响到其他goroutine
// 必须通过 defer 调用
func HandlePanic(action func()) {
	if r := recover(); r != nil {
		action()
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = "unknown"
			line = 0
		}
		s := fmt.Sprintf("%s:%d panic: %v", file, line, r)
		_ = s // TODO: add your log here
	}
}
