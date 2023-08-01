package set

import (
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

// Collection's extended interfaces, can provide more functional abstraction for sets.
type Set[T any] interface {
	iter.Collection[T]

	Contains(element T) bool
	Put(element T) bool
	Remove(element T) util.Opt[T]
	Clear()
}
