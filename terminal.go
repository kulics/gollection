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
func Sum[T constraints.Integer](it Iterator[T]) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else {
			result += item.Second
		}
	}, Enumerate(it))
	return result
}

// Returns the product of all the elements in the iterator.
func Product[T constraints.Integer](it Iterator[T]) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else {
			result *= item.Second
		}
	}, Enumerate(it))
	return result
}

// Returns the average of all the elements in the iterator.
func Average[T constraints.Integer](it Iterator[T]) float64 {
	var result float64
	ForEach(func(item Pair[int, T]) {
		result += (float64(item.Second) - result) / float64(item.First+1)
	}, Enumerate(it))
	return result
}

// Return the total number of iterators.
func Count[T any](it Iterator[T]) int {
	var result int
	ForEach(func(item T) { result++ }, it)
	return result
}

// Return the maximum value of all elements of the iterator.
func Max[T constraints.Integer](it Iterator[T]) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else if result < item.Second {
			result = item.Second
		}
	}, Enumerate(it))
	return result
}

// Return the minimum value of all elements of the iterator.
func Min[T constraints.Integer](it Iterator[T]) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else if result > item.Second {
			result = item.Second
		}
	}, Enumerate(it))
	return result
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
	var curr = it.Next()
	var last = curr
	for curr.IsSome() {
		last = curr
		curr = it.Next()
	}
	return last
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
func Reduce[T any, R any](initial R, operation func(R, T) R, it Iterator[T]) R {
	var result = initial
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		result = operation(result, v)
	}
	return result
}

// Return the value of the final composite, operates on the iterator from back to front.
func Fold[T any, R any](initial R, operation func(T, R) R, it Iterator[T]) R {
	var reverse = make([]T, 0)
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		reverse = append(reverse, v)
	}
	var result = initial
	for i := len(reverse) - 1; i >= 0; i-- {
		result = operation(reverse[i], result)
	}
	return result
}
