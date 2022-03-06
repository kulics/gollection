package main

func Sum[T Number](it Iterable[T]) T {
	var result T
	ForEach(func (item T) { result += item }, it)
	return result
}

func Count[T any](it Iterable[T]) int {
	var result int
	ForEach(func (item T) { result++ }, it)
	return result
}

func Max[T Number](it Iterable[T]) T {
	var result T
	ForEach(func (item T) { if result < item { result = item } }, it)
	return result
}

func Min[T Number](it Iterable[T]) T {
	var result T
	ForEach(func (item T) { if result > item { result = item } }, it)
	return result
}

func ForEach[T any](action func(T), it Iterable[T]) {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		action(v)
	}
}

func AllMatch[T any](predicate func(T) bool, it Iterable[T]) bool {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func NoneMatch[T any](predicate func(T) bool, it Iterable[T]) bool {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		if predicate(v) {
			return false
		}
	}
	return true
}

func AnyMatch[T any](predicate func(T) bool, it Iterable[T]) bool {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		if predicate(v) {
			return true
		}
	}
	return false
}

func First[T any](it Iterable[T]) Option[T] {
	return it.Iter().Next()
}

func Last[T any](it Iterable[T]) Option[T] {
	var iter = it.Iter()
	var result = iter.Next()
	for result.IsSome() {
		result = iter.Next()
	}
	return result
}

func At[T any](index int, it Iterable[T]) Option[T] {
	var iter = it.Iter()
	var result = iter.Next()
	var i = 0
	for i < index && result.IsSome() {
		result = iter.Next()
		i++
	}
	return result
}

func Reduce[T any, R any](initial R, operation func(R, T) R, it Iterable[T]) R {
	var iter = it.Iter()
	var result = initial
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		result = operation(result, v)
	}
	return result
}

func Fold[T any, R any](initial R, operation func(T, R) R, it Iterable[T]) R {
	var reverse = make([]T, 0)
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		reverse = append(reverse, v)
	}
	var result = initial
	for i := len(reverse)-1; i > 0; i-- {
		result = operation(reverse[i], result)
	}
	return result
}