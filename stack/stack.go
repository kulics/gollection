package stack

import (
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

// Collection's extended interfaces, can provide more functional abstraction for stacks.
type Stack[T any] interface {
	iter.Collection[T]

	Push(element T)
	Pop() util.Opt[T]
	Peek() util.Ref[T]
	Clear()
}
