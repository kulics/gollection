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

	Get(index int) T
	Set(index int, newElement T) T
	GetAndSet(index int, set func(oldElement T) T) Pair[T, T]
	TryGet(index int) Option[T]
	TrySet(index int, newElement T) Option[T]

	Prepend(element T)
	Append(element T)
	Insert(index int, element T) bool
	Remove(index int) Option[T]
}

type Map[K any, V any] interface {
	Collection[Pair[K, V]]

	Get(key K) V
	Put(key K, value V) Option[V]
	GetAndPut(key K, set func(oldValue Option[V]) V) Pair[V, Option[V]]
	TryGet(key K) Option[V]

	Remove(key K) Option[V]
	Contains(key K) bool
}

type Stack[T any] interface {
	Size() int
	IsEmpty() bool

	Push(element T)
	Pop() T
	Peek() T
	TryPop() Option[T]
	TryPeek() Option[T]
}
