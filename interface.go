// A collection library based on the go generic implementation,
// providing high performance and functional combinatorial capabilities.
package gollection

// Used to provide built-in computational constraints
type Integer interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~uint64 | ~uint32 | ~uint16 | ~uint8
}

// Used to provide built-in computational constraints, adding floating point numbers over Integer.
type Number interface {
	Integer | ~float64 | ~float32
}

// Iterator can be obtained iteratively through Iter, and each iterator should be independent.
type Iterable[T any] interface {
	Iter() Iterator[T]
}

// By implementing Next you can perform iterations and end them when the return value is None.
type Iterator[T any] interface {
	Next() Option[T]
}

// Iterable's extended interfaces, can provide more information to optimize performance.
type Collection[T any] interface {
	Iterable[T]

	Size() int
	IsEmpty() bool
	ToSlice() []T
}

// Collection's extended interfaces, can provide more functional abstraction for ordered sequences.
type List[T any] interface {
	Collection[T]

	Get(index int) T
	Set(index int, newElement T) T
	Update(index int, update func(oldElement T) T) T
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

// Collection's extended interfaces, can provide more functional abstraction for maps.
type Map[K any, V any] interface {
	Collection[Pair[K, V]]

	Get(key K) V
	Put(key K, value V) Option[V]
	PutAll(elements Collection[Pair[K, V]])
	Update(key K, update func(oldValue Option[V]) V) V
	TryGet(key K) Option[V]

	Remove(key K) Option[V]
	Contains(key K) bool
	Clear()
}

// Collection's extended interfaces, can provide more functional abstraction for sets.
type Set[T any] interface {
	Collection[T]

	Put(element T) bool
	PutAll(elements Collection[T])

	Remove(element T) bool
	Contains(element T) bool
	ContainsAll(elements Collection[T]) bool
	Clear()
}

// Collection's extended interfaces, can provide more functional abstraction for stacks.
type Stack[T any] interface {
	Collection[T]

	Push(element T)
	Pop() T
	Peek() T
	TryPop() Option[T]
	TryPeek() Option[T]
	Clear()
}
