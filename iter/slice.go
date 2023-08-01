package iter

import "github.com/kulics/gollection/util"

// Collection is implemented via Slice, which is isomorphic to the built-in slice.
type Slice[T any] []T

func (a Slice[T]) Iterator() Iterator[T] {
	return &sliceIterator[T]{-1, a}
}

func (a Slice[T]) Count() int {
	return len(a)
}

type sliceIterator[T any] struct {
	index  int
	source []T
}

func (a *sliceIterator[T]) Next() util.Opt[T] {
	if a.index < len(a.source)-1 {
		a.index++
		return util.Some(a.source[a.index])
	}
	return util.None[T]()
}

func CollectToSlice[T any](it Iterator[T]) []T {
	var r = make([]T, 0)
	for {
		if v, ok := it.Next().Val(); ok {
			r = append(r, v)
		} else {
			break
		}
	}
	return r
}
