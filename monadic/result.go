package monadic

// copy from https://github.com/andeya/gust/blob/main/enum_result.go

import (
	"errors"
	"fmt"

	"github.com/zweix123/suger/common"
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
	if r.IsOk() {
		return fmt.Sprintf("Ok[%T](%v)", common.Zero[T](), r.t.Unwrap())
	}
	if r.t.IsNone() {
		if r.e == nil {
			return fmt.Sprintf("Err[%T](not initialized)", common.Zero[T]())
		}
		return fmt.Sprintf("Err[%T](%v)", common.Zero[T](), r.e)
	}
	return fmt.Sprintf("Err[%T](%v)", common.Zero[T](), ErrImpossibleBranch)
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
		if r.e == nil {
			return common.Zero[T](), ErrNotInitialized
		}
		return common.Zero[T](), r.e
	}
	panic(ErrImpossibleBranch) // nolint
}
