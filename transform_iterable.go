package gollection

import (
	. "github.com/kulics/gollection/tuple"
	. "github.com/kulics/gollection/union"
)

func Indexer[T any, I Iterable[T]](it I) IndexerStream[T] {
	return IndexerStream[T]{&struct {
		index    int
		iterator Iterator[T]
	}{-1, it.Iter()}}
}

type IndexerStream[T any] struct {
	inner *struct {
		index    int
		iterator Iterator[T]
	}
}

func (a IndexerStream[T]) Next() Option[Pair[int, T]] {
	if v, ok := a.inner.iterator.Next().Get(); ok {
		a.inner.index++
		return Some(PairOf(a.inner.index, v))
	}
	return None[Pair[int, T]]()
}

func (a IndexerStream[T]) Iter() Iterator[Pair[int, T]] {
	return a
}

func Mapper[T any, R any, I Iterable[T]](transform func(T) R, it I) MapperStream[T, R] {
	return MapperStream[T, R]{transform, it.Iter()}
}

type MapperStream[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (a MapperStream[T, R]) Next() Option[R] {
	if v, ok := a.iterator.Next().Get(); ok {
		return Some(a.transform(v))
	}
	return None[R]()
}

func (a MapperStream[T, R]) Iter() Iterator[R] {
	return a
}

func Filter[T any, I Iterable[T]](predecate func(T) bool, it I) FilterStream[T] {
	return FilterStream[T]{predecate, it.Iter()}
}

type FilterStream[T any] struct {
	predecate func(T) bool
	iterator  Iterator[T]
}

func (a FilterStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.predecate(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (a FilterStream[T]) Iter() Iterator[T] {
	return a
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

func (a LimitStream[T]) Next() Option[T] {
	if v, ok := a.inner.iterator.Next().Get(); ok && a.inner.index < a.inner.limit {
		a.inner.index++
		return Some(v)
	}
	return None[T]()
}

func (a LimitStream[T]) Iter() Iterator[T] {
	return a
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

func (a SkipStream[T]) Next() Option[T] {
	for v, ok := a.inner.iterator.Next().Get(); ok; v, ok = a.inner.iterator.Next().Get() {
		if a.inner.index < a.inner.skip {
			a.inner.index++
			continue
		}
		return Some(v)
	}
	return None[T]()
}

func (a SkipStream[T]) Iter() Iterator[T] {
	return a
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

func (a StepStream[T]) Next() Option[T] {
	for v, ok := a.inner.iterator.Next().Get(); ok; v, ok = a.inner.iterator.Next().Get() {
		if a.inner.index < a.inner.step {
			a.inner.index++
			continue
		}
		a.inner.index = 1
		return Some(v)
	}
	return None[T]()
}

func (a StepStream[T]) Iter() Iterator[T] {
	return a
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

func (a ConcatStream[T]) Next() Option[T] {
	if a.inner.firstNotFinished {
		if v, ok := a.inner.first.Next().Get(); ok {
			return Some(v)
		}
		a.inner.firstNotFinished = false
		return a.Next()
	}
	return a.inner.last.Next()
}

func (a ConcatStream[T]) Iter() Iterator[T] {
	return a
}
