package main

func Some[T any](a T) Option[T] {
	return Option[T]{ a, true }
}

func None[T any]() Option[T] {
	var a T
	return Option[T]{ a, false }
}

type Option[T any] struct {
	value T
	ok bool
}

func (a Option[T]) Get() (value T, ok bool)  {
	return a.value, a.ok
}

func (a Option[T]) GetOrPanic() T {
	if !a.ok {
		panic("none value of option")
	}
	return a.value
}

func (a Option[T]) GetOrDefault(value T) T {
	if !a.ok {
		return value
	}
	return a.value
}

func (a Option[T]) GetOrElse(action func() T) T {
	if !a.ok {
		return action()
	}
	return a.value
}

func (a Option[T]) IsSome() bool {
	return a.ok
}

func (a Option[T]) IsNone() bool {
	return !a.ok
}

func (a Option[T]) IfSome(action func(value T)) {
	if a.ok {
		action(a.value)
	}
}

func (a Option[T]) IfNone(action func()) {
	if !a.ok {
		action()
	}
}