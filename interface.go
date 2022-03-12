package gollection

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
