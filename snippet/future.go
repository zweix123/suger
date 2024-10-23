package snippet

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/zweix123/suger/common"
)

type Future[T any] struct {
	done   chan T
	result T
	err    error
	once   sync.Once
}

func NewFuture[T any](promise func() (T, error)) *Future[T] {
	f := &Future[T]{
		done: make(chan T, 1),
	}

	go func() { // nolint
		defer common.HandlePanic(func(_ string, _ int, r any) {
			f.reject(fmt.Errorf("panic in promise: %v", r))
		})
		defer f.once.Do(func() { close(f.done) })
		result, err := promise()
		if err != nil {
			f.reject(err)
		} else {
			f.resolve(result)
		}
	}()

	return f
}

func (f *Future[T]) resolve(t T) {
	f.once.Do(func() {
		f.result = t
		f.err = nil
		f.done <- t
		close(f.done)
	})
}

func (f *Future[T]) reject(err error) {
	f.once.Do(func() {
		var zero T
		f.result = zero
		f.err = err
		close(f.done)
	})
}

func (f *Future[T]) Get() (T, error) {
	<-f.done
	return f.result, f.err
}

func (f *Future[T]) GetWithTimeout(timeout time.Duration) (T, error) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-f.done:
		return f.result, f.err
	case <-timer.C:
		var zero T
		return zero, ErrTimeout
	default:
		// f.done had been closed
		var zero T
		return zero, ErrTimeout
	}
}

var ErrTimeout = errors.New("future: operation timed out")
