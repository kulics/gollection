package gollection

type Integer interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~uint64 | ~uint32 | ~uint16 | ~uint8
}

type Number interface {
	Integer | ~float64 | ~float32
}

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Iterator[T any] interface {
	Next() Option[T]
}

type Collection[T any] interface {
	Iterable[T]

	Size() int
	IsEmpty() bool
	ToSlice() []T
}

type List[T any] interface {
	Collection[T]

	Get(index int) T
	Set(index int, newElement T) T
	GetAndSet(index int, set func(oldElement T) T) Pair[T, T]
	TryGet(index int) Option[T]
	TrySet(index int, newElement T) Option[T]

	Prepend(element T)
	PrependAll(elements Collection[T])
	Append(element T)
	AppendAll(elements Collection[T])
	Insert(index int, element T)
	InsertAll(index int, elements Collection[T])
	Remove(index int) T
	Clear()
}

type Map[K any, V any] interface {
	Collection[Pair[K, V]]

	Get(key K) V
	Put(key K, value V) Option[V]
	PutAll(elements Collection[Pair[K, V]])
	GetAndPut(key K, set func(oldValue Option[V]) V) Pair[V, Option[V]]
	TryGet(key K) Option[V]

	Remove(key K) Option[V]
	Contains(key K) bool
	Clear()
}

type Set[T any] interface {
	Collection[T]

	Put(element T) bool
	PutAll(elements Collection[T])

	Remove(element T) bool
	Contains(element T) bool
	ContainsAll(elements Collection[T]) bool
	Clear()
}

type Stack[T any] interface {
	Collection[T]

	Push(element T)
	Pop() T
	Peek() T
	TryPop() Option[T]
	TryPeek() Option[T]
	Clear()
}
