package main

type Number interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~uint64 | ~uint32 | ~uint16 | ~uint8 |
		~float64 | ~float32
}

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
}

type List[T any] interface {
	Collection[T]
	Prepend(element T)
	Append(element T)
	Insert(index int, element T) bool
	Remove(index int) Option[T]
	Get(index int) Option[T]
	GetOrPanic(index int) T
	Set(index int, newElement T) Option[T]
}
