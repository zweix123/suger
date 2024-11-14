package snippet

import (
	"errors"
	"fmt"
	"time"

	"github.com/zweix123/suger/common"
)

var (
	ErrPanic   = errors.New("future: panic in promise")
	ErrTimeout = errors.New("future: operation timed out")
)

type Future[T any] struct {
	done   chan struct{} // Only used as a means of synchronization
	result T
	err    error
}

func NewFuture[T any](promise func() (T, error)) *Future[T] {
	f := &Future[T]{
		done: make(chan struct{}, 1),
	}

	go func() { // nolint
		defer close(f.done)
		defer common.HandlePanic(func(_ string, _ int, r any) {
			// if promise panic, the f.result and f.err not assigned correctly, so we need to assign them here
			f.result = common.Zero[T]()
			f.err = fmt.Errorf("%w: %v", ErrPanic, r)
		})
		f.result, f.err = promise()
	}()

	return f
}

func (f *Future[T]) Get() (T, error) {
	<-f.done
	return f.result, f.err
}

func (f *Future[T]) GetWithTimeout(timeout time.Duration) (T, error) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select { // Possible blockage
	case <-f.done:
		return f.result, f.err
	case <-timer.C:
		return common.Zero[T](), ErrTimeout
	}
}
