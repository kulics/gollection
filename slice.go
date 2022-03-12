package gollection

import . "github.com/kulics/gollection/union"

func ToSlice[T any](a []T) Slice[T] {
	return Slice[T](a)
}

type Slice[T any] []T

func (a Slice[T]) Iter() Iterator[T] {
	return &sliceIterator[T]{-1, a}
}

func (a Slice[T]) Size() int {
	return len(a)
}

func (a Slice[T]) IsEmpty() bool {
	return len(a) == 0
}

func (a Slice[T]) ToSlice() []T {
	var slice = make([]T, len(a))
	copy(slice, a)
	return slice
}

type sliceIterator[T any] struct {
	index  int
	source []T
}

func (a *sliceIterator[T]) Next() Option[T] {
	if a.index < len(a.source)-1 {
		a.index++
		return Some(a.source[a.index])
	}
	return None[T]()
}

func (a *sliceIterator[T]) Iter() Iterator[T] {
	return a
}
