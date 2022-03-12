package gollection

import (
	. "github.com/kulics/gollection/tuple"
	. "github.com/kulics/gollection/union"
)

func Indexer[T any, I Iterable[T]](it I) Iterator[Pair[int, T]] {
	return &indexerStream[T]{-1, it.Iter()}
}

type indexerStream[T any] struct {
	index    int
	iterator Iterator[T]
}

func (a *indexerStream[T]) Next() Option[Pair[int, T]] {
	if v, ok := a.iterator.Next().Get(); ok {
		a.index++
		return Some(PairOf(a.index, v))
	}
	return None[Pair[int, T]]()
}

func (a *indexerStream[T]) Iter() Iterator[Pair[int, T]] {
	return a
}

func Mapper[T any, R any, I Iterable[T]](transform func(T) R, it I) Iterator[R] {
	return &mapperStream[T, R]{transform, it.Iter()}
}

type mapperStream[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (a *mapperStream[T, R]) Next() Option[R] {
	if v, ok := a.iterator.Next().Get(); ok {
		return Some(a.transform(v))
	}
	return None[R]()
}

func (a *mapperStream[T, R]) Iter() Iterator[R] {
	return a
}

func Filter[T any, I Iterable[T]](predecate func(T) bool, it I) Iterator[T] {
	return &filterStream[T]{predecate, it.Iter()}
}

type filterStream[T any] struct {
	predecate func(T) bool
	iterator  Iterator[T]
}

func (a *filterStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.predecate(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (a *filterStream[T]) Iter() Iterator[T] {
	return a
}

func Limit[T any, I Iterable[T]](count int, it I) Iterator[T] {
	return &limitStream[T]{count, 0, it.Iter()}
}

type limitStream[T any] struct {
	limit    int
	index    int
	iterator Iterator[T]
}

func (a *limitStream[T]) Next() Option[T] {
	if v, ok := a.iterator.Next().Get(); ok && a.index < a.limit {
		a.index++
		return Some(v)
	}
	return None[T]()
}

func (a *limitStream[T]) Iter() Iterator[T] {
	return a
}

func Skip[T any, I Iterable[T]](count int, it I) Iterator[T] {
	return &skipStream[T]{count, 0, it.Iter()}
}

type skipStream[T any] struct {
	skip     int
	index    int
	iterator Iterator[T]
}

func (a *skipStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.index < a.skip {
			a.index++
			continue
		}
		return Some(v)
	}
	return None[T]()
}

func (a *skipStream[T]) Iter() Iterator[T] {
	return a
}

func Step[T any, I Iterable[T]](count int, it I) Iterator[T] {
	return &stepStream[T]{count, count, it.Iter()}
}

type stepStream[T any] struct {
	step     int
	index    int
	iterator Iterator[T]
}

func (a *stepStream[T]) Next() Option[T] {
	for v, ok := a.iterator.Next().Get(); ok; v, ok = a.iterator.Next().Get() {
		if a.index < a.step {
			a.index++
			continue
		}
		a.index = 1
		return Some(v)
	}
	return None[T]()
}

func (a *stepStream[T]) Iter() Iterator[T] {
	return a
}

func Concat[T any, I Iterable[T]](left I, right I) Iterator[T] {
	return &concatStream[T]{false, left.Iter(), right.Iter()}
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

func (a *concatStream[T]) Iter() Iterator[T] {
	return a
}
