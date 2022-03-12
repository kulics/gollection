package union

func Ok[T any](a T) Result[T] {
	return Result[T]{a, nil}
}

func Err[T any](a error) Result[T] {
	var v T
	return Result[T]{v, a}
}

type Result[T any] struct {
	value T
	err   error
}

func (a Result[T]) Get() (value T, err error) {
	return a.value, a.err
}

func (a Result[T]) OrPanic() T {
	if a.err != nil {
		panic("error of result")
	}
	return a.value
}

func (a Result[T]) OrElse(value T) T {
	if a.err != nil {
		return value
	}
	return a.value
}

func (a Result[T]) OrGet(get func() T) T {
	if a.err != nil {
		return get()
	}
	return a.value
}

func (a Result[T]) IsOk() bool {
	return a.err == nil
}

func (a Result[T]) IsErr() bool {
	return a.err != nil
}

func (a Result[T]) IfOk(action func(value T)) {
	if a.err == nil {
		action(a.value)
	}
}

func (a Result[T]) IfErr(action func()) {
	if a.err != nil {
		action()
	}
}
