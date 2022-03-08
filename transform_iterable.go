package main

func WithIndex[T any, I Iterable[T]](it I) IndexStream[T] {
	return IndexStream[T]{&struct {
		index    int
		iterator Iterator[T]
	}{-1, it.Iter()}}
}

type IndexStream[T any] struct {
	inner *struct {
		index    int
		iterator Iterator[T]
	}
}

func (i IndexStream[T]) Next() Option[Pair[int, T]] {
	if v, ok := i.inner.iterator.Next().Get(); ok {
		i.inner.index++
		return Some(PairOf(i.inner.index, v))
	}
	return None[Pair[int, T]]()
}

func (i IndexStream[T]) Iter() Iterator[Pair[int, T]] {
	return i
}

func Mapper[T any, R any, I Iterable[T]](transform func(T) R, it I) MapperStream[T, R] {
	return MapperStream[T, R]{transform, it.Iter()}
}

type MapperStream[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (m MapperStream[T, R]) Next() Option[R] {
	if v, ok := m.iterator.Next().Get(); ok {
		return Some(m.transform(v))
	}
	return None[R]()
}

func (m MapperStream[T, R]) Iter() Iterator[R] {
	return m
}

func Filter[T any, I Iterable[T]](predecate func(T) bool, it I) FilterStream[T] {
	return FilterStream[T]{predecate, it.Iter()}
}

type FilterStream[T any] struct {
	predecate func(T) bool
	iterator  Iterator[T]
}

func (f FilterStream[T]) Next() Option[T] {
	for v, ok := f.iterator.Next().Get(); ok; v, ok = f.iterator.Next().Get() {
		if f.predecate(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (f FilterStream[T]) Iter() Iterator[T] {
	return f
}

func Limit[T any, I Iterable[T]](count int, it I) LimitStream[T] {
	return LimitStream[T]{&struct {
		limit    int
		index    int
		iterator Iterator[T]
	}{count, 0, it.Iter()}}
}

type LimitStream[T any] struct {
	inner *struct {
		limit    int
		index    int
		iterator Iterator[T]
	}
}

func (l LimitStream[T]) Next() Option[T] {
	if v, ok := l.inner.iterator.Next().Get(); ok && l.inner.index < l.inner.limit {
		l.inner.index++
		return Some(v)
	}
	return None[T]()
}

func (l LimitStream[T]) Iter() Iterator[T] {
	return l
}

func Skip[T any, I Iterable[T]](count int, it I) SkipStream[T] {
	return SkipStream[T]{&struct {
		skip     int
		index    int
		iterator Iterator[T]
	}{count, 0, it.Iter()}}
}

type SkipStream[T any] struct {
	inner *struct {
		skip     int
		index    int
		iterator Iterator[T]
	}
}

func (l SkipStream[T]) Next() Option[T] {
	for v, ok := l.inner.iterator.Next().Get(); ok; v, ok = l.inner.iterator.Next().Get() {
		if l.inner.index < l.inner.skip {
			l.inner.index++
			continue
		}
		return Some(v)
	}
	return None[T]()
}

func (l SkipStream[T]) Iter() Iterator[T] {
	return l
}

func Step[T any, I Iterable[T]](count int, it I) StepStream[T] {
	return StepStream[T]{&struct {
		step     int
		index    int
		iterator Iterator[T]
	}{count, count, it.Iter()}}
}

type StepStream[T any] struct {
	inner *struct {
		step     int
		index    int
		iterator Iterator[T]
	}
}

func (l StepStream[T]) Next() Option[T] {
	for v, ok := l.inner.iterator.Next().Get(); ok; v, ok = l.inner.iterator.Next().Get() {
		if l.inner.index < l.inner.step {
			l.inner.index++
			continue
		}
		l.inner.index = 1
		return Some(v)
	}
	return None[T]()
}

func (l StepStream[T]) Iter() Iterator[T] {
	return l
}

func Concat[T any, I Iterable[T]](left I, right I) ConcatStream[T] {
	return ConcatStream[T]{&struct {
		firstNotFinished bool
		first            Iterator[T]
		last             Iterator[T]
	}{false, left.Iter(), right.Iter()}}
}

type ConcatStream[T any] struct {
	inner *struct {
		firstNotFinished bool
		first            Iterator[T]
		last             Iterator[T]
	}
}

func (l ConcatStream[T]) Next() Option[T] {
	if l.inner.firstNotFinished {
		if v, ok := l.inner.first.Next().Get(); ok {
			return Some(v)
		}
		l.inner.firstNotFinished = false
		return l.Next()
	}
	return l.inner.last.Next()
}

func (l ConcatStream[T]) Iter() Iterator[T] {
	return l
}
