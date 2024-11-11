package monadic

import (
	"reflect"
	"testing"

	suger_test "github.com/zweix123/suger/test"
)

func TestOptionDefaultConstructor(t *testing.T) {
	type CustomType struct{}
	// The option default constructor is None
	var o1 Option[CustomType]
	if o1.IsSome() {
		t.Errorf("expected isSome to be false")
	}
	if !o1.IsNone() {
		t.Errorf("expected isNone to be true")
	}
	var o2 Option[*CustomType]
	if o2.IsSome() {
		t.Errorf("expected isSome to be false")
	}
	if !o2.IsNone() {
		t.Errorf("expected isNone to be true")
	}
}

func TestOptionString(t *testing.T) {
	type C struct {
		I int
		B bool
		F float64
		S string
	}

	tests := []struct {
		name string
		o    interface{} // Option
		s    string
	}{
		{name: "None[int]", o: None[int](), s: "None[int]"},
		{name: "Some[int](1)", o: Some(1), s: "Some[int](1)"},
		{name: "None[float64]", o: None[float64](), s: "None[float64]"},
		{name: "Some[float64](3.14)", o: Some(3.14), s: "Some[float64](3.14)"},
		{name: "None[string]", o: None[string](), s: "None[string]"},
		{name: "Some[string](Hello World)", o: Some("Hello World"), s: "Some[string](Hello World)"},
		{name: "None[monadic.C]", o: None[C](), s: "None[monadic.C]"},
		{name: "Some[monadic.C]({1 true 3.14 Hello World})", o: Some(C{I: 1, B: true, F: 3.14, S: "Hello World"}), s: "Some[monadic.C]({1 true 3.14 Hello World})"},
	}

	check := func(optionString, expectedString string) {
		if optionString != expectedString {
			t.Errorf("expected %s, got %s", expectedString, optionString)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch o := test.o.(type) {
			case Option[int]:
				check(o.String(), test.s)
			case Option[float64]:
				check(o.String(), test.s)
			case Option[string]:
				check(o.String(), test.s)
			case Option[C]:
				check(o.String(), test.s)
			default:
				t.Errorf("expected option, got %T", o)
			}
		})
	}
}

func TestIsSomeAndIsNone(t *testing.T) {
	type C struct {
		I int
		B bool
		F float64
		S string
	}

	tests := []struct {
		name   string
		o      interface{} // Option
		isSome bool
		isNone bool
	}{
		{name: "None[int]", o: None[int](), isSome: false, isNone: true},
		{name: "Some[int](1)", o: Some(1), isSome: true, isNone: false},
		{name: "None[float64]", o: None[float64](), isSome: false, isNone: true},
		{name: "Some[float64](3.14)", o: Some(3.14), isSome: true, isNone: false},
		{name: "None[string]", o: None[string](), isSome: false, isNone: true},
		{name: "Some[string](Hello World)", o: Some("Hello World"), isSome: true, isNone: false},
		{name: "None[monadic.C]", o: None[C](), isSome: false, isNone: true},
		{name: "Some[monadic.C]({1 true 3.14 Hello World})", o: Some(C{I: 1, B: true, F: 3.14, S: "Hello World"}), isSome: true, isNone: false},
	}

	check := func(isSome bool, isNone bool, expectedIsSome bool, expectedIsNone bool) {
		if isSome != expectedIsSome {
			t.Errorf("expected isSome to be %t, got %t", expectedIsSome, isSome)
		}
		if isNone != expectedIsNone {
			t.Errorf("expected isNone to be %t, got %t", expectedIsNone, isNone)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch o := test.o.(type) {
			case Option[int]:
				check(o.IsSome(), o.IsNone(), test.isSome, test.isNone)
			case Option[float64]:
				check(o.IsSome(), o.IsNone(), test.isSome, test.isNone)
			case Option[string]:
				check(o.IsSome(), o.IsNone(), test.isSome, test.isNone)
			case Option[C]:
				check(o.IsSome(), o.IsNone(), test.isSome, test.isNone)
			default:
				t.Errorf("expected option, got %T", o)
			}
		})
	}
}

func TestSomeUnwrap(t *testing.T) {
	type C struct {
		I int
		B bool
		F float64
		S string
	}

	tests := []struct {
		name  string
		o     interface{} // Option
		value interface{}
	}{
		{name: "Some[int](1)", o: Some(1), value: 1},
		{name: "Some[float64](3.14)", o: Some(3.14), value: 3.14},
		{name: "Some[string](Hello World)", o: Some("Hello World"), value: "Hello World"},
		{name: "Some[monadic.C]({1 true 3.14 Hello World})", o: Some(C{I: 1, B: true, F: 3.14, S: "Hello World"}), value: C{I: 1, B: true, F: 3.14, S: "Hello World"}},
	}

	check := func(value interface{}, expectedValue interface{}) {
		if !reflect.DeepEqual(value, expectedValue) {
			t.Errorf("expected value to be %v, got %v", expectedValue, value)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch o := test.o.(type) {
			case Option[int]:
				check(o.Unwrap(), test.value)
			case Option[float64]:
				check(o.Unwrap(), test.value)
			case Option[string]:
				check(o.Unwrap(), test.value)
			case Option[C]:
				check(o.Unwrap(), test.value)
			default:
				t.Errorf("expected option, got %T", o)
			}
		})
	}
}

func TestNoneUnwrap(t *testing.T) {
	type C struct {
		I int
		B bool
		F float64
		S string
	}

	tests := []struct {
		name     string
		o        interface{} // Option
		panicMsg string
	}{
		{name: "None[int]", o: None[int](), panicMsg: "Unwrap called on a None value: None[int]"},
		{name: "None[float64]", o: None[float64](), panicMsg: "Unwrap called on a None value: None[float64]"},
		{name: "None[string]", o: None[string](), panicMsg: "Unwrap called on a None value: None[string]"},
		{name: "None[monadic.C]", o: None[C](), panicMsg: "Unwrap called on a None value: None[monadic.C]"},
	}

	check := func(panicMsg string, expectedPanicMsg string) {
		if panicMsg != expectedPanicMsg {
			t.Errorf("expected panic message to be %s, got %s", expectedPanicMsg, panicMsg)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch o := test.o.(type) {
			case Option[int]:
				check(suger_test.HandlePanic(func() { o.Unwrap() }), test.panicMsg)
			case Option[float64]:
				check(suger_test.HandlePanic(func() { o.Unwrap() }), test.panicMsg)
			case Option[string]:
				check(suger_test.HandlePanic(func() { o.Unwrap() }), test.panicMsg)
			case Option[C]:
				check(suger_test.HandlePanic(func() { o.Unwrap() }), test.panicMsg)
			default:
				t.Errorf("expected option, got %T", o)
			}
		})
	}
}

func TestPointer(t *testing.T) {
	type C struct {
		I int
		B bool
		F float64
		S string
	}
	some := Some[*C](nil)
	if !some.IsSome() {
		t.Errorf("expected some to be some")
	}
	if some.IsNone() {
		t.Errorf("expected some to be none")
	}
	value := some.Unwrap()
	if value != nil {
		t.Errorf("expected value to be nil")
	}
}
