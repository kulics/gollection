package gollection

// Constructing an Result with a success value.
func Ok[T any](a T) Result[T] {
	return Result[T]{a, nil}
}

// Constructing an Result with a error value.
func Err[T any](a error) Result[T] {
	var v T
	return Result[T]{v, a}
}

// Type-safe errorable types.
// Provides a safe way to manipulate values that may be error.
// It is nestable.
type Result[T any] struct {
	value T
	err   error
}

// Get can use go's customary deconstructed Result,
// which is used like the built-in map, and can safely use value when error is nil.
func (a Result[T]) Get() (value T, err error) {
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
func (a Result[T]) OrElse(value T) T {
	if a.err != nil {
		return value
	}
	return a.value
}

// Get the value in an safe way, and get else value when error is not nil.
// The get function is lazy and can be used instead of OrElse when you need to avoid unnecessary computations.
func (a Result[T]) OrGet(get func() T) T {
	if a.err != nil {
		return get()
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
func (a Result[T]) IfErr(action func()) {
	if a.err != nil {
		action()
	}
}
