package list

import (
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

type List[T any] interface {
	iter.Collection[T]

	Peek() util.Ref[T]
	Push(element T)
	Pop() util.Opt[T]
	Clear()
}

// Collection's extended interfaces, can provide more functional abstraction for ordered sequences.
type BiList[T any] interface {
	List[T]

	PeekFront() util.Ref[T]
	PushFront(element T)
	PopFront() util.Opt[T]

	PeekBack() util.Ref[T]
	PushBack(element T)
	PopBack() util.Opt[T]
}

// Collection's extended interfaces, can provide more functional abstraction for ordered sequences.
type IdxList[T any] interface {
	List[T]

	At(index int) util.Ref[T]
	Insert(index int, element T)
	Remove(index int) T
	RemoveRange(at iter.Range[int])
}

func Equals[T comparable](l iter.Collection[T], r iter.Collection[T]) bool {
	if l.Count() != r.Count() {
		return false
	}
	var lIter = l.Iterator()
	var rIter = r.Iterator()
	for {
		if v1, ok1 := lIter.Next().Val(); ok1 {
			if v2, _ := rIter.Next().Val(); v1 != v2 {
				return false
			}
		} else {
			break
		}
	}
	return true
}

func FirstIndexOf[T comparable](li iter.Iterable[T], element T) int {
	var iter = iter.Enumerate(li.Iterator())
	for {
		if v, ok := iter.Next().Val(); ok {
			if v.Second == element {
				return v.First
			}
		} else {
			break
		}
	}
	return -1
}
