package util

// Constructing an Option with a value.
func Some[T any](a T) Opt[T] {
	return Opt[T]{a, true}
}

// Constructing an Option without a value
func None[T any]() Opt[T] {
	return Opt[T]{}
}

// Type-safe nullable types.
// Provides a safe way to manipulate values that may be null.
// It is nestable.
type Opt[T any] struct {
	value T
	ok    bool
}

// Val can use go's customary deconstructed Option,
// which is used like the built-in map, and can safely use value when ok is true.
func (a Opt[T]) Val() (value T, ok bool) {
	return a.value, a.ok
}

// Get the value in an unsafe way, and execute panic when ok is false.
func (a Opt[T]) OrPanic() T {
	if !a.ok {
		panic("none value of option")
	}
	return a.value
}

// Get the value in an safe way, and get else value when ok is false.
func (a Opt[T]) Or(value T) T {
	if !a.ok {
		return value
	}
	return a.value
}

// Get the value in an safe way, and get else value when ok is false.
func (a Opt[T]) OrDefault() (v T) {
	if !a.ok {
		return
	}
	return a.value
}

// Returns true when ok is true.
func (a Opt[T]) IsSome() bool {
	return a.ok
}

// Returns true when ok is false.
func (a Opt[T]) IsNone() bool {
	return !a.ok
}

// Execute the action when ok is true.
func (a Opt[T]) IfSome(action func(value T)) {
	if a.ok {
		action(a.value)
	}
}

// Execute the action when ok is false.
func (a Opt[T]) IfNone(action func()) {
	if !a.ok {
		action()
	}
}

func (a Opt[T]) Next() Opt[T] {
	return a
}
