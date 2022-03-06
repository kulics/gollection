package main

func Sum[T Number](it Iterable[T]) T {
	var result T
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		result += v
	}
	return result
}

func Count[T any](it Iterable[T]) int {
	var result int
	var iter = it.Iter()
	for _, ok := iter.Next(); ok; _, ok = iter.Next() {
		result++
	}
	return result
}

func Max[T Number](it Iterable[T]) T {
	var result T
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if result < v {
			result = v
		}
	}
	return result
}

func Min[T Number](it Iterable[T]) T {
	var result T
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if result > v {
			result = v
		}
	}
	return result
}

func ForEach[T any](action func(T), it Iterable[T]) {
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		action(v)
	}
}

func AllMatch[T any](predicate func(T) bool, it Iterable[T]) bool {
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func NoneMatch[T any](predicate func(T) bool, it Iterable[T]) bool {
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if predicate(v) {
			return false
		}
	}
	return true
}

func AnyMatch[T any](predicate func(T) bool, it Iterable[T]) bool {
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		if predicate(v) {
			return true
		}
	}
	return false
}

func First[T any](it Iterable[T]) (value T, ok bool) {
	return it.Iter().Next()
}

func Last[T any](it Iterable[T]) (value T, ok bool) {
	var iter = it.Iter()
	v, ok := iter.Next()
	for ok {
		v, ok = iter.Next()
	}
	return v, ok
}

func At[T any](index int, it Iterable[T]) (value T, ok bool) {
	var iter = it.Iter()
	v, ok := iter.Next()
	var i = 0
	for i < index && ok {
		v, ok = iter.Next()
		i++
	}
	return v, ok
}

func Reduce[T any, R any](initial R, operation func(R, T) R, it Iterable[T]) R {
	var iter = it.Iter()
	var result = initial
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		result = operation(result, v)
	}
	return result
}

func Fold[T any, R any](initial R, operation func(T, R) R, it Iterable[T]) R {
	var reverse = make([]T, 0)
	var iter = it.Iter()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		reverse = append(reverse, v)
	}
	var result = initial
	for i := len(reverse)-1; i > 0; i-- {
		result = operation(reverse[i], result)
	}
	return result
}