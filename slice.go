package gollection

func ToSlice[T any](a []T) Slice[T] {
	return Slice[T](a)
}

type Slice[T any] []T

func (s Slice[T]) Iter() Iterator[T] {
	return &sliceIterator[T]{-1, s}
}

func (s Slice[T]) Size() int {
	return len(s)
}

func (s Slice[T]) IsEmpty() bool {
	return len(s) == 0
}

type sliceIterator[T any] struct {
	index  int
	source []T
}

func (s *sliceIterator[T]) Next() Option[T] {
	if s.index < len(s.source)-1 {
		s.index++
		return Some(s.source[s.index])
	}
	return None[T]()
}

func (s *sliceIterator[T]) Iter() Iterator[T] {
	return s
}
