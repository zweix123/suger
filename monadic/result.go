package monadic

// copy from https://github.com/andeya/gust/blob/main/enum_result.go

import (
	"errors"
	"fmt"
)

var (
	ErrNotInitialized   = errors.New("not initialized")
	ErrImpossibleBranch = errors.New("impossible branch, program bug")
)

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

func (r Result[T]) String() string {
	var zero T
	if r.IsOk() {
		return fmt.Sprintf("Ok[%T](%v)", zero, r.t.Unwrap())
	}
	if r.t.IsNone() {
		var zero T
		if r.e == nil {
			return fmt.Sprintf("Err[%T](not initialized)", zero)
		}
		return fmt.Sprintf("Err[%T](%v)", zero, r.e)
	}
	return fmt.Sprintf("Err[%T](%v)", zero, ErrImpossibleBranch)
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
	panic(ErrImpossibleBranch) // nolint
}
