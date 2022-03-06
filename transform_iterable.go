package main

type Pair[T any, R any] struct {
	First T
	Second R
}

func WithIndex[T any](it Iterable[T]) Iterator[Pair[int, T]] {
	return &indexStream[T]{-1, it.Iter()}
}

type indexStream[T any] struct {
	index int
	iterator Iterator[T]
}

func (i *indexStream[T]) Next() (value Pair[int, T], ok bool) {
	if v, ok := i.iterator.Next(); ok {
		i.index++
		return Pair[int, T]{i.index, v}, true
	}
	return
}

func (i *indexStream[T]) Iter() Iterator[Pair[int, T]] {
	return i
}

func Map[T any, R any](transform func(T) R, it Iterable[T]) Iterator[R] {
	return mapStream[T, R]{transform, it.Iter()}
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iterator Iterator[T]
}

func (m mapStream[T, R]) Next() (value R, ok bool) {
	if v, ok := m.iterator.Next(); ok {
		return m.transform(v), true
	}
	return
}

func (m mapStream[T, R]) Iter() Iterator[R] {
	return m
}

func Filter[T any](predecate func(T) bool, it Iterable[T]) Iterator[T] {
	return filterStream[T]{predecate, it.Iter()}
}

type filterStream[T any] struct {
	predecate func(T) bool
	iterator Iterator[T]
}

func (f filterStream[T]) Next() (value T, ok bool) {
	for v, ok := f.iterator.Next(); ok; v, ok = f.iterator.Next() {
		if f.predecate(v) {
			return v, true
		}
	}
	return
}

func (f filterStream[T]) Iter() Iterator[T] {
	return f
}

func Limit[T any](count int, it Iterable[T]) Iterator[T] {
	return &limitStream[T]{count, 0, it.Iter()}
}

type limitStream[T any] struct {
	limit int
	index int
	iterator Iterator[T]
}

func (l *limitStream[T]) Next() (value T, ok bool) {
	if v, ok := l.iterator.Next(); ok && l.index < l.limit {
		l.index++
		return v, true
	}
	return
}

func (l *limitStream[T]) Iter() Iterator[T] {
	return l
}

func Skip[T any](count int, it Iterable[T]) Iterator[T] {
	return &skipStream[T]{count, 0, it.Iter()}
}

type skipStream[T any] struct {
	skip int
	index int
	iterator Iterator[T]
}

func (l *skipStream[T]) Next() (value T, ok bool) {
	for v, ok := l.iterator.Next(); ok; v, ok = l.iterator.Next() {
		if l.index < l.skip {
			l.index++
			continue
		}
		return v, true
	}
	return
}

func (l *skipStream[T]) Iter() Iterator[T] {
	return l
}

func Step[T any](count int, it Iterable[T]) Iterator[T] {
	return &stepStream[T]{count, 0, it.Iter()}
}

type stepStream[T any] struct {
	step int
	index int
	iterator Iterator[T]
}

func (l *stepStream[T]) Next() (value T, ok bool) {
	for v, ok := l.iterator.Next(); ok; v, ok = l.iterator.Next() {
		if l.index < l.step {
			l.index++
			continue
		}
		l.index = 0
		return v, true
	}
	return
}

func (l *stepStream[T]) Iter() Iterator[T] {
	return l
}

func Concat[T any](left Iterable[T], right Iterable[T]) Iterator[T] {
	return &concatStream[T]{false, left.Iter(), right.Iter()}
}

type concatStream[T any] struct {
	firstok bool
	first Iterator[T]
	last Iterator[T]
}

func (l *concatStream[T]) Next() (value T, ok bool) {
	if l.firstok {
		if v, ok := l.first.Next(); ok {
			return v, true
		}
		l.firstok = false
		return l.Next()
	}
	return l.last.Next()
}

func (l *concatStream[T]) Iter() Iterator[T] {
	return l
}
