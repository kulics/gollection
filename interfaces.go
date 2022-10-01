// A collection library based on the go generic implementation,
// providing high performance and functional combinatorial capabilities.
package gollection

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

	Count() int
	IsEmpty() bool
	ToSlice() []T
}

// Collection's extended interfaces, can provide more functional abstraction for ordered sequences.
type AnyList[T any] interface {
	Collection[T]

	LastIndex() int

	Get(index int) T
	TryGet(index int) Option[T]

	GetFirst() T
	TryGetFirst() Option[T]

	GetLast() T
	TryGetLast() Option[T]
}

// List's extended interfaces, can provide more mutable functional abstraction for ordered sequences.
type AnyMutableList[T any] interface {
	AnyList[T]

	Set(index int, newElement T) T
	TrySet(index int, newElement T) Option[T]

	Prepend(element T)
	PrependAll(elements Collection[T])

	Append(element T)
	AppendAll(elements Collection[T])

	Insert(index int, element T)
	InsertAll(index int, elements Collection[T])

	RemoveAt(index int) T
	RemoveRange(at Range[int])

	Clear()
}

// Collection's extended interfaces, can provide more functional abstraction for maps.
type AnyMap[K any, V any] interface {
	Collection[Pair[K, V]]

	Get(key K) V
	TryGet(key K) Option[V]

	Contains(key K) bool
}

// Map's extended interfaces, can provide more mutable functional abstraction for maps.
type AnyMutableMap[K any, V any] interface {
	AnyMap[K, V]

	Put(key K, value V) Option[V]
	PutAll(elements Collection[Pair[K, V]])

	Remove(key K) Option[V]

	Clear()
}

// Collection's extended interfaces, can provide more functional abstraction for sets.
type AnySet[T any] interface {
	Collection[T]

	Contains(element T) bool
	ContainsAll(elements Collection[T]) bool
}

// Set's extended interfaces, can provide more functional abstraction for sets.
type AnyMutableSet[T any] interface {
	AnySet[T]

	Put(element T) bool
	PutAll(elements Collection[T])

	Remove(element T) bool

	Clear()
}

// Collection's extended interfaces, can provide more functional abstraction for stacks.
type AnyStack[T any] interface {
	Collection[T]

	Push(element T)

	Pop() T
	TryPop() Option[T]

	Peek() T
	TryPeek() Option[T]

	Clear()
}

func EqualsList[T comparable](l AnyList[T], r AnyList[T]) bool {
	if l.Count() != r.Count() {
		return false
	}
	var lIter = l.Iter()
	var rIter = r.Iter()
	for v, ok := lIter.Next().Get(); ok; v, ok = lIter.Next().Get() {
		if v != rIter.Next().OrPanic() {
			return false
		}
	}
	return true
}

func FirstIndexOf[T comparable](li AnyList[T], element T) int {
	var iter = Enumerate(li.Iter())
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		if v.Second == element {
			return v.First
		}
	}
	return -1
}

func EqualsMap[K any, V comparable](l AnyMap[K, V], r AnyMap[K, V]) bool {
	if l.Count() != r.Count() {
		return false
	}
	var lIter = l.Iter()
	for pair, ok := lIter.Next().Get(); ok; pair, ok = lIter.Next().Get() {
		if v, ok := r.TryGet(pair.First).Get(); ok && v == pair.Second {
			return false
		}
	}
	return true
}
