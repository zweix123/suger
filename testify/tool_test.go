package testify

import (
	"fmt"
	"testing"
)

func TestHandleStdout(t *testing.T) {
	buf := HandleStdout(func() {
		fmt.Println("Hello, world!")
	})
	if buf.String() != "Hello, world!\n" {
		t.Errorf("Expected output to be %q, but got %q", "Hello, world!\n", buf.String())
	}
}

func TestHandlePanic(t *testing.T) {
	panicMsg := HandlePanic(func() {
		panic("TestHandlePanic") // nolint
	})
	if panicMsg != "TestHandlePanic" {
		t.Errorf("Expected panic message to be %q, but got %q", "TestHandlePanic", panicMsg)
	}
	noPanic := HandlePanic(func() {})
	if noPanic != "" {
		t.Errorf("Expected no panic, but got %q", noPanic)
	}
}
