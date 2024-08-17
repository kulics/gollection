package ref

func Of[T any](v *T) Ref[T] {
	return Ref[T]{v}
}

type Ref[T any] struct {
	Ptr *T
}

func (a Ref[T]) Val() (v T, ok bool) {
	if a.Ptr == nil {
		return
	}
	return *a.Ptr, true
}

func (a Ref[T]) Get() T {
	return *a.Ptr
}

func (a Ref[T]) Set(v T) T {
	var old = *a.Ptr
	*a.Ptr = v
	return old
}

func (a Ref[T]) IsNil() bool {
	return a.Ptr == nil
}

func (a Ref[T]) IsNotNil() bool {
	return a.Ptr != nil
}
