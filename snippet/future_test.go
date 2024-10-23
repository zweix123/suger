package snippet

import (
	"errors"
	"testing"
	"time"
)

func TestFutureNormal(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		return 1, nil
	})
	result, err := future.Get()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != 1 {
		t.Errorf("expected result 1, got %v", result)
	}
}

func TestFutureTimeout(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		time.Sleep(time.Second)
		return 1, nil
	})
	_, err := future.GetWithTimeout(time.Millisecond)
	if err != ErrTimeout {
		t.Errorf("expected timeout error, got %v", err)
	}
}

func TestFutureError(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		return 0, errors.New("test error")
	})
	_, err := future.Get()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestFutureTimeoutError(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		time.Sleep(time.Second)
		return 0, errors.New("test error")
	})
	_, err := future.GetWithTimeout(time.Millisecond)
	if err != ErrTimeout {
		t.Errorf("expected timeout error, got %v", err)
	}
}
