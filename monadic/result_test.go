package monadic

import (
	"errors"
	"reflect"
	"testing"
)

func TestIsOkAndIsErr(t *testing.T) {
	type C struct {
		I int
		B bool
		F float64
		S string
	}
	e := errors.New("error")

	tests := []struct {
		name  string
		r     interface{} // Result[T]
		isOk  bool
		isErr bool
	}{
		{name: "int-ok", r: Ok(1), isOk: true, isErr: false},
		{name: "int-err", r: Err[int](e), isOk: false, isErr: true},
		{name: "float-ok", r: Ok(3.14), isOk: true, isErr: false},
		{name: "float-err", r: Err[float64](e), isOk: false, isErr: true},
		{name: "string-ok", r: Ok("Hello World"), isOk: true, isErr: false},
		{name: "string-err", r: Err[string](e), isOk: false, isErr: true},
		{name: "C-ok", r: Ok(C{I: 1, B: true, F: 3.14, S: "Hello World"}), isOk: true, isErr: false},
		{name: "C-err", r: Err[C](e), isOk: false, isErr: true},
	}

	check := func(isOk, isErr, expectIsOk, expectIsErr bool) {
		if isOk != expectIsOk {
			t.Errorf("expected IsOk to be %v, got %v", expectIsOk, isOk)
		}
		if isErr != expectIsErr {
			t.Errorf("expected IsErr to be %v, got %v", expectIsErr, isErr)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch r := test.r.(type) {
			case Result[int]:
				check(r.IsOk(), r.IsErr(), test.isOk, test.isErr)
			case Result[float64]:
				check(r.IsOk(), r.IsErr(), test.isOk, test.isErr)
			case Result[string]:
				check(r.IsOk(), r.IsErr(), test.isOk, test.isErr)
			case Result[C]:
				check(r.IsOk(), r.IsErr(), test.isOk, test.isErr)
			default:
				t.Errorf("unexpected type: %T", r)
			}
		})
	}
}

func TestUnwrap(t *testing.T) {
	type C struct {
		I int
		B bool
		F float64
		S string
	}
	e := errors.New("error")

	tests := []struct {
		name string
		r    interface{} // Result[T]
		v    interface{}
		e    error
	}{
		{name: "int-ok", r: Ok(1), v: 1, e: nil},
		{name: "int-err", r: Err[int](e), v: 0, e: e},
		{name: "float-ok", r: Ok(3.14), v: 3.14, e: nil},
		{name: "float-err", r: Err[float64](e), v: 0.0, e: e},
		{name: "string-ok", r: Ok("Hello World"), v: "Hello World", e: nil},
		{name: "string-err", r: Err[string](e), v: "", e: e},
		{name: "C-ok", r: Ok(C{I: 1, B: true, F: 3.14, S: "Hello World"}), v: C{I: 1, B: true, F: 3.14, S: "Hello World"}, e: nil},
		{name: "C-err", r: Err[C](e), v: C{}, e: e},
	}

	check := func(v interface{}, e error, expectV interface{}, expectE error) {
		if !reflect.DeepEqual(v, expectV) {
			t.Errorf("expected value to be %v, got %v", expectV, v)
		}
		if e != expectE {
			t.Errorf("expected error to be %v, got %v", expectE, e)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch r := test.r.(type) {
			case Result[int]:
				v, e := r.Unwrap()
				check(v, e, test.v, test.e)
			case Result[float64]:
				v, e := r.Unwrap()
				check(v, e, test.v, test.e)
			case Result[string]:
				v, e := r.Unwrap()
				check(v, e, test.v, test.e)
			case Result[C]:
				v, e := r.Unwrap()
				check(v, e, test.v, test.e)
			default:
				t.Errorf("unexpected type: %T", r)
			}
		})
	}
}