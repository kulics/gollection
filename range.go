package gollection

import "golang.org/x/exp/constraints"

type Range[T constraints.Integer] struct {
	begin T
	end   T
}

// Return the begin value and end value.
func (a Range[T]) Get() (T, T) {
	return a.begin, a.end
}

// Return the begin value.
func (a Range[T]) Begin() T {
	return a.begin
}

// Return the end value.
func (a Range[T]) End() T {
	return a.end
}

// Return the index is in range.
func (a Range[T]) Has(index T) bool {
	return index > a.begin && index < a.end
}

func (a Range[T]) Iter() Iterator[T] {
	return &rangeIter[T]{index: a.begin, end: a.end}
}

type rangeIter[T constraints.Integer] struct {
	index T
	end   T
}

func (a *rangeIter[T]) Next() Option[T] {
	if a.index < a.end {
		var i = a.index
		a.index++
		return Some(i)
	}
	return None[T]()
}

// Constructing an Option with the begin and end.
func RangeOf[T constraints.Integer](begin T, end T) Range[T] {
	if end < begin {
		panic("end can not less than begin")
	}
	return Range[T]{begin: begin, end: end}
}
