package gollection

// Constructing an Option with a value.
func Some[T any](a T) Option[T] {
	return Option[T]{a, true}
}

// Constructing an Option without a value
func None[T any]() Option[T] {
	var a T
	return Option[T]{a, false}
}

// Type-safe nullable types.
// Provides a safe way to manipulate values that may be null.
// It is nestable.
type Option[T any] struct {
	value T
	ok    bool
}

// Get can use go's customary deconstructed Option,
// which is used like the built-in map, and can safely use value when ok is true.
func (a Option[T]) Get() (value T, ok bool) {
	return a.value, a.ok
}

// Get the value in an unsafe way, and execute panic when ok is false.
func (a Option[T]) OrPanic() T {
	if !a.ok {
		panic("none value of option")
	}
	return a.value
}

// Get the value in an safe way, and get else value when ok is false.
func (a Option[T]) OrElse(value T) T {
	if !a.ok {
		return value
	}
	return a.value
}

// Get the value in an safe way, and get else value when ok is false.
// The get function is lazy and can be used instead of OrElse when you need to avoid unnecessary computations.
func (a Option[T]) OrGet(get func() T) T {
	if !a.ok {
		return get()
	}
	return a.value
}

// Returns true when ok is true.
func (a Option[T]) IsSome() bool {
	return a.ok
}

// Returns true when ok is false.
func (a Option[T]) IsNone() bool {
	return !a.ok
}

// Execute the action when ok is true.
func (a Option[T]) IfSome(action func(value T)) {
	if a.ok {
		action(a.value)
	}
}

// Execute the action when ok is false.
func (a Option[T]) IfNone(action func()) {
	if !a.ok {
		action()
	}
}
