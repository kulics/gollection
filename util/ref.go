package util

func RefOf[T any](v *T) Ref[T] {
	return Ref[T]{v}
}

type Ref[T any] struct {
	value *T
}

func (a Ref[T]) Val() (v T, ok bool) {
	if a.value == nil {
		return
	}
	return *a.value, true
}

func (a Ref[T]) Get() T {
	return *a.value
}

func (a Ref[T]) Set(v T) T {
	var old = *a.value
	*a.value = v
	return old
}

func (a Ref[T]) IsNil() bool {
	return a.value == nil
}

func (a Ref[T]) IsNotNil() bool {
	return a.value != nil
}
