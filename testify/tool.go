/*
 * Testing utility methods, the code will be filled with panic, so it is not recommended for online services.
 */

package testify

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func HandleStdout(f func()) (buf bytes.Buffer) {
	tmp := os.Stdout
	defer func() {
		os.Stdout = tmp
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

// captureOut captures both stdout and stderr.
func captureOut(f func()) string {
	// Create a pipe to capture stdout
	custReader, custWriter, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// Save the original stdout and stderr to restore later
	origStdout := os.Stdout
	origStderr := os.Stderr

	// Restore stdout and stderr when done
	defer func() {
		os.Stdout = origStdout
		os.Stderr = origStderr
	}()

	// Set the stdout and stderr to the pipe
	os.Stdout, os.Stderr = custWriter, custWriter
	log.SetOutput(custWriter)

	// Create a channel to read the output from the pipe
	out := make(chan string)

	// Use a goroutine to read from the pipe and send the output to the channel
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, custReader)
		out <- buf.String()
	}()
	wg.Wait()

	// Call the function that writes to stdout
	f()

	// Close the writer to signal that we're done
	_ = custWriter.Close()

	// Wait for the goroutine to finish reading from the pipe
	return <-out
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
