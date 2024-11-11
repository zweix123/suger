package monadic

import "errors"

var ErrNoErrButValueNone = errors.New("no error but value is none") //! It won't happen in theory

func Ok[T any](t T) Result[T] {
	return Result[T]{t: Some(t), e: nil}
}

func Err[T any](e error) Result[T] {
	return Result[T]{t: None[T](), e: e}
}

type Result[T any] struct {
	t Option[T]
	e error
}

func (r Result[T]) IsErr() bool {
	return r.e != nil
}

func (r Result[T]) IsOk() bool {
	return r.e == nil
}

func (r Result[T]) Unwrap() (T, error) {
	if r.IsErr() {
		var zero T
		return zero, r.e
	}
	if r.t.IsNone() {
		var zero T
		return zero, ErrNoErrButValueNone
	}
	return r.t.Unwrap(), nil
}
