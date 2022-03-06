package main

type Slice[T any] struct {
	Source []T
}

func (s Slice[T]) Iter() Iterator[T] {
	return SliceIter(s.Source)
}

func (s Slice[T]) Size() int {
	return len(s.Source)
}

func (s Slice[T]) IsEmpty() bool {
	return len(s.Source) == 0
}

func ToSlice[T any](s []T) Slice[T] {
	return Slice[T]{s}
}

func SliceIter[T any](source []T) Iterator[T] {
	return &sliceIterator[T]{-1, source}
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
