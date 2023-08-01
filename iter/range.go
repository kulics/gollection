package iter

import (
	"github.com/kulics/gollection/util"
	"golang.org/x/exp/constraints"
)

type Range[T constraints.Integer] struct {
	Begin T
	End   T
}

// Return the begin value and end value.
func (a Range[T]) Get() (T, T) {
	return a.Begin, a.End
}

// Return the index is in range.
func (a Range[T]) Has(index T) bool {
	return index > a.Begin && index < a.End
}

func (a Range[T]) Iterator() Iterator[T] {
	return &rangeIterator[T]{index: a.Begin, end: a.End}
}

type rangeIterator[T constraints.Integer] struct {
	index T
	end   T
}

func (a *rangeIterator[T]) Next() util.Opt[T] {
	if a.index < a.end {
		var i = a.index
		a.index++
		return util.Some(i)
	}
	return util.None[T]()
}

// Constructing an Range with the begin and end.
func RangeOf[T constraints.Integer](begin T, end T) Range[T] {
	if end < begin {
		panic("end can not less than begin")
	}
	return Range[T]{Begin: begin, End: end}
}
