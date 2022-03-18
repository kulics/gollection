package gollection

func Contains[T comparable](target T, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if v == target {
			return true
		}
	}
	return false
}

func Sum[T Number](it Iterator[T]) T {
	var result T
	ForEach(func(item T) { result += item }, it)
	return result
}

func Product[T Number](it Iterator[T]) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else {
			result *= item.Second
		}
	}, Indexer[T](it))
	return result
}

func Average[T Number](it Iterator[T]) float64 {
	var result float64
	ForEach(func(item Pair[int, T]) {
		result += (float64(item.Second) - result) / float64(item.First+1)
	}, Indexer[T](it))
	return result
}

func Count[T any](it Iterator[T]) int {
	var result int
	ForEach(func(item T) { result++ }, it)
	return result
}

func Max[T Number](it Iterator[T]) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else if result < item.Second {
			result = item.Second
		}
	}, Indexer[T](it))
	return result
}

func Min[T Number](it Iterator[T]) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else if result > item.Second {
			result = item.Second
		}
	}, Indexer[T](it))
	return result
}

func ForEach[T any](action func(T), it Iterator[T]) {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		action(v)
	}
}

func AllMatch[T any](predicate func(T) bool, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func NoneMatch[T any](predicate func(T) bool, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if predicate(v) {
			return false
		}
	}
	return true
}

func AnyMatch[T any](predicate func(T) bool, it Iterator[T]) bool {
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		if predicate(v) {
			return true
		}
	}
	return false
}

func First[T any](it Iterator[T]) Option[T] {
	return it.Next()
}

func Last[T any](it Iterator[T]) Option[T] {
	var result = it.Next()
	for result.IsSome() {
		result = it.Next()
	}
	return result
}

func At[T any](index int, it Iterator[T]) Option[T] {
	var result = it.Next()
	var i = 0
	for i < index && result.IsSome() {
		result = it.Next()
		i++
	}
	return result
}

func Reduce[T any, R any](initial R, operation func(R, T) R, it Iterator[T]) R {
	var result = initial
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		result = operation(result, v)
	}
	return result
}

func Fold[T any, R any](initial R, operation func(T, R) R, it Iterator[T]) R {
	var reverse = make([]T, 0)
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		reverse = append(reverse, v)
	}
	var result = initial
	for i := len(reverse) - 1; i > 0; i-- {
		result = operation(reverse[i], result)
	}
	return result
}
