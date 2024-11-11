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
		panic(err)
	}
	os.Stdout = w

	func() {
		defer w.Close()
		f()
	}()

	_, err = buf.ReadFrom(r)
	if err != nil {
		panic(err)
	}

	return buf
}

func HandlePanic(f func()) string {
	var buf bytes.Buffer
	func(buf *bytes.Buffer) {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Fprintf(buf, "%v", r)
			}
		}()
		f()
	}(&buf)
	return buf.String()
}
