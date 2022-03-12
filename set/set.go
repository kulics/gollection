package set

import . "github.com/kulics/gollection"

type Set[T any] interface {
	Collection[T]

	Put(element T) bool
	PutAll(elements Collection[T])

	Remove(element T) bool
	Contains(element T) bool
	ContainsAll(elements Collection[T]) bool
	Clear()
}
