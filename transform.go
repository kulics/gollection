package gollection

// Add subscripts to the incoming iterators.
func Enumerate[T any](it Iterator[T]) Iterator[Pair[int, T]] {
	return &enumerateStream[T]{-1, it}
}

type enumerateStream[T any] struct {
	index    int
	iterator Iterator[T]
}

func (a *enumerateStream[T]) Next() Option[Pair[int, T]] {
	if v, ok := a.iterator.Next().Get(); ok {
		a.index++
		return Some(PairOf(a.index, v))
	}
	return None[Pair[int, T]]()
}

// Use transform to map an iterator to another iterator.
func Map[T any, R any](transform func(T) R, it Iterator[T]) Iterator[R] {
	return &mapStream[T, R]{transform, it}
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (a *mapStream[T, R]) Next() Option[R] {
	if v, ok := a.iterator.Next().Get(); ok {
		return Some(a.transform(v))
	}
	return None[R]()
}

// Use predicate to filter an iterator to another iterator。
func Filter[T any](predicate func(T) bool, it Iterator[T]) Iterator[T] {
	return &filterStream[T]{predicate, it}
}

type filterStream[T any] struct {
	predicate func(T) bool
	iterator  Iterator[T]
}

func (a *filterStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.predicate(v) {
			return Some(v)
		}
	}
	return None[T]()
}

// Convert an iterator to another iterator that limits the maximum number of iterations.
func Limit[T any](count int, it Iterator[T]) Iterator[T] {
	return &limitStream[T]{count, it}
}

type limitStream[T any] struct {
	limit    int
	iterator Iterator[T]
}

func (a *limitStream[T]) Next() Option[T] {
	if a.limit != 0 {
		a.limit -= 1
		return a.iterator.Next()
	}
	return None[T]()
}

// Converts an iterator to another iterator that skips a specified number of times.
func Skip[T any](count int, it Iterator[T]) Iterator[T] {
	return &skipStream[T]{count, it}
}

type skipStream[T any] struct {
	skip     int
	iterator Iterator[T]
}

func (a *skipStream[T]) Next() Option[T] {
	for a.skip > 0 {
		if a.iterator.Next().IsNone() {
			return None[T]()
		}
		a.skip -= 1
	}
	return a.iterator.Next()
}

// Converts an iterator to another iterator that skips a specified number of times each time.
func Step[T any](count int, it Iterator[T]) Iterator[T] {
	return &stepStream[T]{count - 1, true, it}
}

type stepStream[T any] struct {
	step      int
	firstTake bool
	iterator  Iterator[T]
}

func (a *stepStream[T]) Next() Option[T] {
	if a.firstTake {
		a.firstTake = false
		return a.iterator.Next()
	} else {
		return At(a.step, a.iterator)
	}
}

// By connecting two iterators in series,
// the new iterator will iterate over the first iterator before continuing with the second iterator.
func Concat[T any](left Iterator[T], right Iterator[T]) Iterator[T] {
	return &concatStream[T]{false, left, right}
}

type concatStream[T any] struct {
	firstNotFinished bool
	first            Iterator[T]
	last             Iterator[T]
}

func (a *concatStream[T]) Next() Option[T] {
	if a.firstNotFinished {
		if v, ok := a.first.Next().Get(); ok {
			return Some(v)
		}
		a.firstNotFinished = false
		return a.Next()
	}
	return a.last.Next()
}

// Converting a nested iterator to a flat iterator.
func Flatten[T Iterable[U], U any](it Iterator[T]) Iterator[U] {
	return &flattenStream[T, U]{it, None[Iterator[U]]()}
}

type flattenStream[T Iterable[U], U any] struct {
	iterator Iterator[T]
	subIter  Option[Iterator[U]]
}

func (a *flattenStream[T, U]) Next() Option[U] {
	if iter, ok := a.subIter.Get(); ok {
		if item, ok := iter.Next().Get(); ok {
			return Some(item)
		} else {
			a.subIter = None[Iterator[U]]()
			return a.Next()
		}
	} else if nextIter, ok := a.iterator.Next().Get(); ok {
		a.subIter = Some(nextIter.Iter())
		return a.Next()
	} else {
		return None[U]()
	}
}

// Compress two iterators into one iterator. The length is the length of the shortest iterator.
func Zip[T any, U any](left Iterator[T], right Iterator[U]) Iterator[Pair[T, U]] {
	return &zipStream[T, U]{left, right}
}

type zipStream[T any, U any] struct {
	first Iterator[T]
	last  Iterator[U]
}

func (a *zipStream[T, U]) Next() Option[Pair[T, U]] {
	if v1, ok1 := a.first.Next().Get(); ok1 {
		if v2, ok2 := a.last.Next().Get(); ok2 {
			return Some(PairOf(v1, v2))
		}
	}
	return None[Pair[T, U]]()
}
