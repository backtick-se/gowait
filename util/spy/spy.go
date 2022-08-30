package spy

import "reflect"

type Spy interface {
	Calls() int
	Called() bool
	CalledWith(...any) bool
}

var _ Spy = &Void{}
var _ Spy = &E{}
var _ Spy = &E1[int]{}
var _ Spy = &E2[int, int]{}

type Call []any

func (c Call) Matches(args ...any) bool {
	if len(c) != len(args) {
		return false
	}
	for i := range c {
		if !reflect.DeepEqual(c[i], args[i]) {
			return false
		}
	}
	return true
}

//
// No return values
//

type Void struct {
	calls []Call
}

func (t *Void) Call(args ...any) {
	t.calls = append(t.calls, Call(args))
}

func (t Void) Calls() int {
	return len(t.calls)
}

func (t Void) Called() bool {
	return t.Calls() > 0
}

func (t Void) CalledWith(args ...any) bool {
	for _, call := range t.calls {
		if call.Matches(args...) {
			return true
		}
	}
	return false
}

//
// Error only
//

type E struct {
	Void
	Error error
}

func (t *E) Call(args ...any) error {
	t.Void.Call(args...)
	return t.Error
}

//
// 1 Return value, error
//

type E1[R any] struct {
	E
	Result R
}

func (t *E1[R]) Call(args ...any) (R, error) {
	if err := t.E.Call(args...); err != nil {
		var empty R
		return empty, err
	}
	return t.Result, nil
}

//
// 2 Return values, error
//

type E2[R1, R2 any] struct {
	E
	Result1 R1
	Result2 R2
}

func (t *E2[R1, R2]) Call(args ...any) (R1, R2, error) {
	if err := t.E.Call(args...); err != nil {
		var empty1 R1
		var empty2 R2
		return empty1, empty2, err
	}
	return t.Result1, t.Result2, nil
}
