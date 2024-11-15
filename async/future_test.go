package async

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

func TestFutureWait(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		time.Sleep(3 * time.Second)
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

func TestFutureWaitTimeout(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		time.Sleep(1 * time.Second)
		return 1, nil
	})
	result, err := future.GetWithTimeout(3 * time.Second)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != 1 {
		t.Errorf("expected result 1, got %v", result)
	}
}

func TestFutureTimeout(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		time.Sleep(3 * time.Second)
		return 1, nil
	})
	_, err := future.GetWithTimeout(1 * time.Second)
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
		time.Sleep(3 * time.Second)
		return 0, errors.New("test error")
	})
	_, err := future.GetWithTimeout(1 * time.Second)
	if err != ErrTimeout {
		t.Errorf("expected timeout error, got %v", err)
	}
}

func TestFuturePanic(t *testing.T) {
	future := NewFuture[int](func() (int, error) {
		panic("test panic") // nolint
	})
	_, err := future.Get()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if !errors.Is(err, ErrPanic) {
		t.Errorf("expected panic error, got %v", err)
	}
}
