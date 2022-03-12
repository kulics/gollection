package list

import (
	. "github.com/kulics/gollection"
	. "github.com/kulics/gollection/tuple"
	. "github.com/kulics/gollection/union"
)

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
