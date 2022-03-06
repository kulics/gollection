package main

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func PairOf[T1 any, T2 any](f T1, s T2) Pair[T1, T2] {
	return Pair[T1, T2]{f, s}
}

func WithIndex[T any](it Iterable[T]) Iterator[Pair[int, T]] {
	return &indexStream[T]{-1, it.Iter()}
}

type indexStream[T any] struct {
	index    int
	iterator Iterator[T]
}

func (i *indexStream[T]) Next() Option[Pair[int, T]] {
	if v, ok := i.iterator.Next().Get(); ok {
		i.index++
		return Some(PairOf(i.index, v))
	}
	return None[Pair[int, T]]()
}

func (i *indexStream[T]) Iter() Iterator[Pair[int, T]] {
	return i
}

func Map[T any, R any](transform func(T) R, it Iterable[T]) Iterator[R] {
	return mapStream[T, R]{transform, it.Iter()}
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (m mapStream[T, R]) Next() Option[R] {
	if v, ok := m.iterator.Next().Get(); ok {
		return Some(m.transform(v))
	}
	return None[R]()
}

func (m mapStream[T, R]) Iter() Iterator[R] {
	return m
}

func Filter[T any](predecate func(T) bool, it Iterable[T]) Iterator[T] {
	return filterStream[T]{predecate, it.Iter()}
}

type filterStream[T any] struct {
	predecate func(T) bool
	iterator  Iterator[T]
}

func (f filterStream[T]) Next() Option[T] {
	for v, ok := f.iterator.Next().Get(); ok; v, ok = f.iterator.Next().Get() {
		if f.predecate(v) {
			return Some(v)
		}
	}
	return None[T]()
}

func (f filterStream[T]) Iter() Iterator[T] {
	return f
}

func Limit[T any](count int, it Iterable[T]) Iterator[T] {
	return &limitStream[T]{count, 0, it.Iter()}
}

type limitStream[T any] struct {
	limit    int
	index    int
	iterator Iterator[T]
}

func (l *limitStream[T]) Next() Option[T] {
	if v, ok := l.iterator.Next().Get(); ok && l.index < l.limit {
		l.index++
		return Some(v)
	}
	return None[T]()
}

func (l *limitStream[T]) Iter() Iterator[T] {
	return l
}

func Skip[T any](count int, it Iterable[T]) Iterator[T] {
	return &skipStream[T]{count, 0, it.Iter()}
}

type skipStream[T any] struct {
	skip     int
	index    int
	iterator Iterator[T]
}

func (l *skipStream[T]) Next() Option[T] {
	for v, ok := l.iterator.Next().Get(); ok; v, ok = l.iterator.Next().Get() {
		if l.index < l.skip {
			l.index++
			continue
		}
		return Some(v)
	}
	return None[T]()
}

func (l *skipStream[T]) Iter() Iterator[T] {
	return l
}

func Step[T any](count int, it Iterable[T]) Iterator[T] {
	return &stepStream[T]{count, 0, it.Iter()}
}

type stepStream[T any] struct {
	step     int
	index    int
	iterator Iterator[T]
}

func (l *stepStream[T]) Next() Option[T] {
	for v, ok := l.iterator.Next().Get(); ok; v, ok = l.iterator.Next().Get() {
		if l.index < l.step {
			l.index++
			continue
		}
		l.index = 0
		return Some(v)
	}
	return None[T]()
}

func (l *stepStream[T]) Iter() Iterator[T] {
	return l
}

func Concat[T any](left Iterable[T], right Iterable[T]) Iterator[T] {
	return &concatStream[T]{false, left.Iter(), right.Iter()}
}

type concatStream[T any] struct {
	firstok bool
	first   Iterator[T]
	last    Iterator[T]
}

func (l *concatStream[T]) Next() Option[T] {
	if l.firstok {
		if v, ok := l.first.Next().Get(); ok {
			return Some(v)
		}
		l.firstok = false
		return l.Next()
	}
	return l.last.Next()
}

func (l *concatStream[T]) Iter() Iterator[T] {
	return l
}
