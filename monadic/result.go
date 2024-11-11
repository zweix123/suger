package monadic

// copy from https://github.com/andeya/gust/blob/main/enum_result.go

import "errors"

var ErrNotInitialized = errors.New("not initialized")

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

/*
|            | IsSome | IsNone               |
| ---------- | ------ | -------------------- |
| err == nil | IsOk   | Err(not initialized) |
| err != nil | bug    | IsErr                |
*/

func (r Result[T]) IsOk() bool {
	return r.t.IsSome() && r.e == nil
}

func (r Result[T]) IsErr() bool {
	return !r.IsOk()
}

func (r Result[T]) Unwrap() (T, error) {
	if r.IsOk() {
		return r.t.Unwrap(), nil
	}
	if r.t.IsNone() {
		var zero T
		if r.e == nil {
			return zero, ErrNotInitialized
		}
		return zero, r.e
	}
	panic("Impossible branch, program bug")
}
