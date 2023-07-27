package gollection

import "golang.org/x/exp/constraints"

// Returns true if the target is included in the iterator.
func Contains[T comparable](target T, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if v == target {
			return true
		}
	}
	return false
}

// Returns the sum of all the elements in the iterator.
func Sum[T constraints.Integer | constraints.Float](it Iterator[T]) T {
	return Fold(0, func(a, b T) T {
		return a + b
	}, it)
}

// Returns the product of all the elements in the iterator.
func Product[T constraints.Integer | constraints.Float](it Iterator[T]) T {
	return Fold(1, func(a, b T) T {
		return a * b
	}, it)
}

// Returns the average of all the elements in the iterator.
func Average[T constraints.Integer | constraints.Float](it Iterator[T]) float64 {
	return Fold(0.0, func(result float64, item Pair[int, T]) float64 {
		return result + (float64(item.Second)-result)/float64(item.First+1)
	}, Enumerate(it))
}

// Return the total number of iterators.
func Count[T any](it Iterator[T]) int {
	return Fold(0, func(v int, _ T) int {
		return v + 1
	}, it)
}

// Return the maximum value of all elements of the iterator.
func Max[T constraints.Ordered](it Iterator[T]) Option[T] {
	return Reduce(func(a T, b T) T {
		if a > b {
			return a
		} else {
			return b
		}
	}, it)
}

// Return the maximum value of all elements of the iterator.
func MaxBy[T any](greater func(T, T) bool, it Iterator[T]) Option[T] {
	return Reduce(func(a T, b T) T {
		if greater(a, b) {
			return a
		} else {
			return b
		}
	}, it)
}

// Return the minimum value of all elements of the iterator.
func Min[T constraints.Ordered](it Iterator[T]) Option[T] {
	return Reduce(func(a T, b T) T {
		if a < b {
			return a
		} else {
			return b
		}
	}, it)
}

// Return the minimum value of all elements of the iterator.
func MinBy[T any](less func(T, T) bool, it Iterator[T]) Option[T] {
	return Reduce(func(a T, b T) T {
		if less(a, b) {
			return a
		} else {
			return b
		}
	}, it)
}

// The action is executed for each element of the iterator, and the argument to the action is the element.
func ForEach[T any](action func(T), it Iterator[T]) {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		action(v)
	}
}

// Returns true if all elements in the iterator match the condition.
func AllMatch[T any](predicate func(T) bool, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Returns true if none elements in the iterator match the condition.
func NoneMatch[T any](predicate func(T) bool, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if predicate(v) {
			return false
		}
	}
	return true
}

// Returns true if any elements in the iterator match the condition.
func AnyMatch[T any](predicate func(T) bool, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Return the first element.
func First[T any](it Iterator[T]) Option[T] {
	return it.Next()
}

// Return the last element.
func Last[T any](it Iterator[T]) Option[T] {
	return Fold(None[T](), func(_ Option[T], next T) Option[T] {
		return Some(next)
	}, it)
}

// Return the element at index.
func At[T any](index int, it Iterator[T]) Option[T] {
	var result = it.Next()
	var i = 0
	for i < index && result.IsSome() {
		result = it.Next()
		i++
	}
	return result
}

// Return the value of the final composite, operates on the iterator from front to back.
func Reduce[T any](operation func(T, T) T, it Iterator[T]) Option[T] {
	if v, ok := it.Next().Get(); ok {
		return Some(Fold(v, operation, it))
	}
	return None[T]()
}

// Return the value of the final composite, operates on the iterator from back to front.
func Fold[T any, R any](initial R, operation func(R, T) R, it Iterator[T]) R {
	var result = initial
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		result = operation(result, v)
	}
	return result
}

// Splitting an iterator whose elements are pair into two lists.
func Unzip[A any, B any](it Iterator[Pair[A, B]]) Pair[*ArrayList[A], *ArrayList[B]] {
	var arrA = ArrayListOf[A]()
	var arrB = ArrayListOf[B]()
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		var a, b = v.Get()
		arrA.Append(a)
		arrB.Append(b)
	}
	return PairOf(arrA, arrB)
}

type Collector[S any, T any, R any] interface {
	Supply() S
	Accumulate(supplier S, element T)
	Finish(supplier S) R
}

// Collecting via Collector.
func Collect[T any, S any, R any](collector Collector[S, T, R], it Iterator[T]) R {
	var s = collector.Supply()
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		collector.Accumulate(s, v)
	}
	return collector.Finish(s)
}

// Collect to built-in map.
func CollectToMap[K comparable, V any](it Iterator[Pair[K, V]]) map[K]V {
	var r = make(map[K]V, 0)
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		r[v.First] = v.Second
	}
	return r
}
