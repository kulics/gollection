package gollection

import . "github.com/kulics/gollection/union"

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Iterator[T any] interface {
	Iterable[T]

	Next() Option[T]
}

type Collection[T any] interface {
	Iterable[T]

	Size() int
	IsEmpty() bool
	ToSlice() []T
}
