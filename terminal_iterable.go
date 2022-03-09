package gollection

func Sum[T Number, I Iterable[T]](it I) T {
	var result T
	ForEach(func(item T) { result += item }, it)
	return result
}

func Product[T Number, I Iterable[T]](it I) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else {
			result *= item.Second
		}
	}, WithIndex[T](it))
	return result
}

func Average[T Number, I Iterable[T]](it I) float64 {
	var result float64
	ForEach(func(item Pair[int, T]) {
		result += (float64(item.Second) - result) / float64(item.First+1)
	}, WithIndex[T](it))
	return result
}

func Count[T any, I Iterable[T]](it I) int {
	var result int
	ForEach(func(item T) { result++ }, it)
	return result
}

func Max[T Number, I Iterable[T]](it I) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else if result < item.Second {
			result = item.Second
		}
	}, WithIndex[T](it))
	return result
}

func Min[T Number, I Iterable[T]](it I) T {
	var result T
	ForEach(func(item Pair[int, T]) {
		if item.First == 0 {
			result = item.Second
		} else if result > item.Second {
			result = item.Second
		}
	}, WithIndex[T](it))
	return result
}

func ForEach[T any, I Iterable[T]](action func(T), it I) {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		action(v)
	}
}

func AllMatch[T any, I Iterable[T]](predicate func(T) bool, it I) bool {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func NoneMatch[T any, I Iterable[T]](predicate func(T) bool, it I) bool {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		if predicate(v) {
			return false
		}
	}
	return true
}

func AnyMatch[T any, I Iterable[T]](predicate func(T) bool, it I) bool {
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		if predicate(v) {
			return true
		}
	}
	return false
}

func First[T any, I Iterable[T]](it I) Option[T] {
	return it.Iter().Next()
}

func Last[T any, I Iterable[T]](it I) Option[T] {
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

func Reduce[T any, R any, I Iterable[T]](initial R, operation func(R, T) R, it I) R {
	var iter = it.Iter()
	var result = initial
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		result = operation(result, v)
	}
	return result
}

func Fold[T any, R any, I Iterable[T]](initial R, operation func(T, R) R, it I) R {
	var reverse = make([]T, 0)
	var iter = it.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		reverse = append(reverse, v)
	}
	var result = initial
	for i := len(reverse) - 1; i > 0; i-- {
		result = operation(reverse[i], result)
	}
	return result
}
