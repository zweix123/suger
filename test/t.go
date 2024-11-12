/*
 * Testing related functions, the code will be filled with panic, so it is not recommended for online services
 */

package test

import (
	"bytes"
	"fmt"
	"os"
)

func HandleStdout(f func()) (buf bytes.Buffer) {
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()

	r, w, err := os.Pipe()
	if err != nil {
		panic(err) // nolint
	}
	os.Stdout = w

	func() {
		defer w.Close()
		f()
	}()

	_, err = buf.ReadFrom(r)
	if err != nil {
		panic(err) // nolint
	}

	return buf
}

func HandlePanic(f func()) string {
	var buf bytes.Buffer
	func(buf *bytes.Buffer) {
		defer func() {
			r := recover()
			if r != nil {
				_, err := fmt.Fprintf(buf, "%v", r)
				if err != nil {
					panic(err) // nolint
				}
			}
		}()
		f()
	}(&buf)
	return buf.String()
}
