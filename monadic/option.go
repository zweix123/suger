package monadic

// copy from https://github.com/andeya/gust/blob/main/option.go

import "fmt"

func Some[T any](t T) Option[T] {
	v := &t
	return Option[T]{value: &v}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

type Option[T any] struct {
	value **T
	// Why **T:
	// 1. Allows distinction between Some(*T(nil)) and None[*T](): Some(*T(nil)).IsSome() == true, Some(*T(nil)).Unwrap() == nil
	// 2. Default constructor of Option[T] is None[T]() in terms of behavior
}

func (o Option[T]) String() string {
	var zero T
	if o.IsNone() {
		return fmt.Sprintf("None[%T]", zero)
	}
	return fmt.Sprintf("Some[%T](%v)", zero, o.unwrapUnchecked())
}

func (o Option[T]) unwrapUnchecked() T {
	return **o.value
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) IsSome() bool {
	return !o.IsNone()
}

func (o Option[T]) Unwrap() T {
	if o.IsSome() {
		return o.unwrapUnchecked()
	}
	panic(fmt.Sprintf("Unwrap called on a None value: %v", o)) // nolint
}
