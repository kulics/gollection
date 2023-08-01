package util

// Constructing an Result with a success value.
func Ok[T any](a T) Result[T] {
	return Result[T]{value: a}
}

// Constructing an Result with a error value.
func Err[T any](a error) Result[T] {
	return Result[T]{err: a}
}

// Type-safe errorable types.
// Provides a safe way to manipulate values that may be error.
// It is nestable.
type Result[T any] struct {
	value T
	err   error
}

// Val can use go's customary deconstructed Result,
// which is used like the built-in map, and can safely use value when error is nil.
func (a Result[T]) Val() (value T, err error) {
	return a.value, a.err
}

// Get the value in an unsafe way, and execute panic when error is not nil.
func (a Result[T]) OrPanic() T {
	if a.err != nil {
		panic("error of result")
	}
	return a.value
}

// Get the value in an safe way, and get else value when error is not nil.
func (a Result[T]) Or(value T) T {
	if a.err != nil {
		return value
	}
	return a.value
}

func (a Result[T]) OrDefault() (v T) {
	if a.err != nil {
		return
	}
	return a.value
}

// Returns true when error is nil.
func (a Result[T]) IsOk() bool {
	return a.err == nil
}

// Returns true when error is not nil.
func (a Result[T]) IsErr() bool {
	return a.err != nil
}

// Execute the action when error is nil.
func (a Result[T]) IfOk(action func(value T)) {
	if a.err == nil {
		action(a.value)
	}
}

// Execute the action when error is not nil.
func (a Result[T]) IfErr(action func(err error)) {
	if a.err != nil {
		action(a.err)
	}
}
