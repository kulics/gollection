package seq

import (
	"github.com/kulics/gollection/option"
	"golang.org/x/exp/constraints"
)

// Ruturns true if the count of collection is 0.
func IsEmpty[T any](c Collection[T]) bool {
	return c.Count() == 0
}

// Ruturns true if the count of collection is 0.
func IsNotEmpty[T any](c Collection[T]) bool {
	return c.Count() != 0
}

// Converts a collection to a Slice.
func ToSlice[T any](c Collection[T]) Slice[T] {
	var arr = make([]T, 0, c.Count())
	ForEach[T](func(t T) {
		arr = append(arr, t)
	}, c)
	return arr
}

// Returns true if the target is included in the Sequence.
func Contains[T comparable](target T, it Sequence[T]) bool {
	var iter = it.Iterator()
	for {
		if v, ok := iter.Next().Val(); ok {
			if v == target {
				return true
			}
		} else {
			break
		}
	}
	return false
}

// Returns the sum of all the elements in the Sequence.
func Sum[T constraints.Integer | constraints.Float](it Sequence[T]) T {
	return Fold(0, func(a, b T) T {
		return a + b
	}, it)
}

// Returns the product of all the elements in the Sequence.
func Product[T constraints.Integer | constraints.Float](it Sequence[T]) T {
	return Fold(1, func(a, b T) T {
		return a * b
	}, it)
}

// Returns the average of all the elements in the Sequence.
func Average[T constraints.Integer | constraints.Float](it Sequence[T]) float64 {
	return Fold(0.0, func(result float64, item Pair[int, T]) float64 {
		return result + (float64(item.Second)-result)/float64(item.First+1)
	}, Enumerate(it))
}

// Return the total number of Sequence.
func Count[T any](it Sequence[T]) int {
	return Fold(0, func(v int, _ T) int {
		return v + 1
	}, it)
}

// Return the maximum value of all elements of the Sequence.
func Max[T constraints.Ordered](it Sequence[T]) option.Option[T] {
	return Reduce(func(a T, b T) T {
		if a > b {
			return a
		} else {
			return b
		}
	}, it)
}

// Return the maximum value of all elements of the Sequence.
func MaxBy[T any](greater func(T, T) bool, it Sequence[T]) option.Option[T] {
	return Reduce(func(a T, b T) T {
		if greater(a, b) {
			return a
		} else {
			return b
		}
	}, it)
}

// Return the minimum value of all elements of the Sequence.
func Min[T constraints.Ordered](it Sequence[T]) option.Option[T] {
	return Reduce(func(a T, b T) T {
		if a < b {
			return a
		} else {
			return b
		}
	}, it)
}

// Return the minimum value of all elements of the Sequence.
func MinBy[T any](less func(T, T) bool, it Sequence[T]) option.Option[T] {
	return Reduce(func(a T, b T) T {
		if less(a, b) {
			return a
		} else {
			return b
		}
	}, it)
}

// The action is executed for each element of the Sequence, and the argument to the action is the element.
func ForEach[T any](action func(T), it Sequence[T]) {
	var iter = it.Iterator()
	for {
		if v, ok := iter.Next().Val(); ok {
			action(v)
		} else {
			break
		}
	}
}

// Returns true if all elements in the Sequence match the condition.
func AllMatch[T any](predicate func(T) bool, it Sequence[T]) bool {
	var iter = it.Iterator()
	for {
		if v, ok := iter.Next().Val(); ok {
			if !predicate(v) {
				return false
			}
		} else {
			break
		}
	}
	return true
}

// Returns true if none elements in the Sequence match the condition.
func NoneMatch[T any](predicate func(T) bool, it Sequence[T]) bool {
	var iter = it.Iterator()
	for {
		if v, ok := iter.Next().Val(); ok {
			if predicate(v) {
				return false
			}
		} else {
			break
		}
	}
	return true
}

// Returns true if any elements in the Sequence match the condition.
func AnyMatch[T any](predicate func(T) bool, it Sequence[T]) bool {
	var iter = it.Iterator()
	for {
		if v, ok := iter.Next().Val(); ok {
			if predicate(v) {
				return true
			}
		} else {
			break
		}
	}
	return false
}

// Return the first element.
func First[T any](it Sequence[T]) option.Option[T] {
	return it.Iterator().Next()
}

// Return the last element.
func Last[T any](it Sequence[T]) option.Option[T] {
	return Fold(option.None[T](), func(_ option.Option[T], next T) option.Option[T] {
		return option.Some(next)
	}, it)
}

// Return the element at index.
func At[T any](index int, it Sequence[T]) option.Option[T] {
	var iter = it.Iterator()
	var result = iter.Next()
	var i = 0
	for i < index && result.IsSome() {
		result = iter.Next()
		i++
	}
	return result
}

// Return the value of the final composite, operates on the Sequence from front to back.
func Reduce[T any](operation func(T, T) T, it Sequence[T]) option.Option[T] {
	var iter = it.Iterator()
	if v, ok := iter.Next().Val(); ok {
		var result = v
		for {
			if v, ok := iter.Next().Val(); ok {
				result = operation(result, v)
			} else {
				break
			}
		}
		return option.Some(result)
	}
	return option.None[T]()
}

// Return the value of the final composite, operates on the Sequence from back to front.
func Fold[T any, R any](initial R, operation func(R, T) R, it Sequence[T]) R {
	var result = initial
	var iter = it.Iterator()
	for {
		if v, ok := iter.Next().Val(); ok {
			result = operation(result, v)
		} else {
			break
		}
	}
	return result
}

type Collector[S any, T any, R any] interface {
	Builder() S
	Append(builder S, element T)
	Finish(builder S) R
}

// Collecting via Collector.
func Collect[T any, S any, R any](collector Collector[S, T, R], it Sequence[T]) R {
	var iter = it.Iterator()
	var s = collector.Builder()
	for {
		if v, ok := iter.Next().Val(); ok {
			collector.Append(s, v)
		} else {
			break
		}
	}
	return collector.Finish(s)
}

func FirstIndexOf[T comparable](li Sequence[T], element T) int {
	var iter = Enumerate(li).Iterator()
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
