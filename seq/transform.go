package seq

import (
	"github.com/kulics/gollection/option"
)

// Add subscripts to the incoming Sequence.
func Enumerate[T any](it Sequence[T]) Sequence[Pair[int, T]] {
	return enumerateSequence[T]{it}
}

type enumerateSequence[T any] struct {
	seq Sequence[T]
}

func (a enumerateSequence[T]) Iterator() Iterator[Pair[int, T]] {
	return &enumerateIterator[T]{-1, a.seq.Iterator()}
}

type enumerateIterator[T any] struct {
	index    int
	iterator Iterator[T]
}

func (a *enumerateIterator[T]) Next() option.Option[Pair[int, T]] {
	if v, ok := a.iterator.Next().Val(); ok {
		a.index++
		return option.Some(Pair[int, T]{a.index, v})
	}
	return option.None[Pair[int, T]]()
}

// Use transform to map an Sequence to another Sequence.
func Map[T any, R any](transform func(T) R, it Sequence[T]) Sequence[R] {
	return mapSequence[T, R]{transform, it}
}

type mapSequence[T, R any] struct {
	transform func(T) R
	seq       Sequence[T]
}

func (a mapSequence[T, R]) Iterator() Iterator[R] {
	return &mapIterator[T, R]{a.transform, a.seq.Iterator()}
}

type mapIterator[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (a *mapIterator[T, R]) Next() option.Option[R] {
	if v, ok := a.iterator.Next().Val(); ok {
		return option.Some(a.transform(v))
	}
	return option.None[R]()
}

// Use predicate to filter an Sequence to another Sequence
func Filter[T any](predicate func(T) bool, it Sequence[T]) Sequence[T] {
	return filterSequence[T]{predicate, it}
}

type filterSequence[T any] struct {
	predicate func(T) bool
	seq       Sequence[T]
}

func (a filterSequence[T]) Iterator() Iterator[T] {
	return &filterIterator[T]{a.predicate, a.seq.Iterator()}
}

type filterIterator[T any] struct {
	predicate func(T) bool
	iterator  Iterator[T]
}

func (a *filterIterator[T]) Next() option.Option[T] {
	for {
		if v, ok := a.iterator.Next().Val(); ok {
			if a.predicate(v) {
				return option.Some(v)
			}
		} else {
			break
		}
	}
	return option.None[T]()
}

// Convert an Sequence to another Sequence that limits the maximum number of iterations.
func Limit[T any](count int, it Sequence[T]) Sequence[T] {
	return limitSequence[T]{count, it}
}

type limitSequence[T any] struct {
	limit int
	seq   Sequence[T]
}

func (a limitSequence[T]) Iterator() Iterator[T] {
	return &limitIterator[T]{a.limit, a.seq.Iterator()}
}

type limitIterator[T any] struct {
	limit    int
	iterator Iterator[T]
}

func (a *limitIterator[T]) Next() option.Option[T] {
	if a.limit != 0 {
		a.limit -= 1
		return a.iterator.Next()
	}
	return option.None[T]()
}

// Converts an Sequence to another Sequence that skips a specified number of times.
func Skip[T any](count int, it Sequence[T]) Sequence[T] {
	return skipSequence[T]{count, it}
}

type skipSequence[T any] struct {
	skip int
	seq  Sequence[T]
}

func (a skipSequence[T]) Iterator() Iterator[T] {
	return &skipIterator[T]{a.skip, a.seq.Iterator()}
}

type skipIterator[T any] struct {
	skip     int
	iterator Iterator[T]
}

func (a *skipIterator[T]) Next() option.Option[T] {
	for a.skip > 0 {
		if a.iterator.Next().IsNone() {
			return option.None[T]()
		}
		a.skip -= 1
	}
	return a.iterator.Next()
}

// Converts an Sequence to another Sequence that skips a specified number of times each time.
func Step[T any](count int, it Sequence[T]) Sequence[T] {
	return stepSequence[T]{count - 1, it}
}

type stepSequence[T any] struct {
	step int
	seq  Sequence[T]
}

func (a stepSequence[T]) Iterator() Iterator[T] {
	return &stepIterator[T]{a.step, true, a.seq.Iterator()}
}

type stepIterator[T any] struct {
	step      int
	firstTake bool
	iterator  Iterator[T]
}

func (a *stepIterator[T]) Next() option.Option[T] {
	if a.firstTake {
		a.firstTake = false
		return a.iterator.Next()
	} else {
		var iter = a.iterator
		var result = iter.Next()
		var i = 0
		for i < a.step && result.IsSome() {
			result = iter.Next()
			i++
		}
		return result
	}
}

// By connecting two Sequences in series,
// the new Sequence will iterate over the first Sequence before continuing with the second Sequence.
func Concat[T any](left Sequence[T], right Sequence[T]) Sequence[T] {
	return concatSequence[T]{left, right}
}

type concatSequence[T any] struct {
	first Sequence[T]
	last  Sequence[T]
}

func (a concatSequence[T]) Iterator() Iterator[T] {
	return &concatStream[T]{false, a.first.Iterator(), a.last.Iterator()}
}

type concatStream[T any] struct {
	firstNotFinished bool
	first            Iterator[T]
	last             Iterator[T]
}

func (a *concatStream[T]) Next() option.Option[T] {
	if a.firstNotFinished {
		if v, ok := a.first.Next().Val(); ok {
			return option.Some(v)
		}
		a.firstNotFinished = false
		return a.Next()
	}
	return a.last.Next()
}

// Converting a nested Sequence to a flat Sequence.
func Flatten[T Sequence[U], U any](it Sequence[T]) Sequence[U] {
	return flattenSequence[T, U]{it}
}

type flattenSequence[T Sequence[U], U any] struct {
	seq Sequence[T]
}

func (a flattenSequence[T, U]) Iterator() Iterator[U] {
	return &flattenIterator[T, U]{a.seq.Iterator(), option.None[Iterator[U]]()}
}

type flattenIterator[T Sequence[U], U any] struct {
	iterator Iterator[T]
	subIter  option.Option[Iterator[U]]
}

func (a *flattenIterator[T, U]) Next() option.Option[U] {
	if iter, ok := a.subIter.Val(); ok {
		if item, ok := iter.Next().Val(); ok {
			return option.Some(item)
		} else {
			a.subIter = option.None[Iterator[U]]()
			return a.Next()
		}
	} else if nextIter, ok := a.iterator.Next().Val(); ok {
		a.subIter = option.Some(nextIter.Iterator())
		return a.Next()
	} else {
		return option.None[U]()
	}
}

// Compress two Sequences into one Sequence. The length is the length of the shortest Sequence.
func Zip[T any, U any](left Sequence[T], right Sequence[U]) Sequence[Pair[T, U]] {
	return zipSequence[T, U]{left, right}
}

type zipSequence[T, U any] struct {
	first Sequence[T]
	last  Sequence[U]
}

func (a zipSequence[T, U]) Iterator() Iterator[Pair[T, U]] {
	return &zipIterator[T, U]{a.first.Iterator(), a.last.Iterator()}
}

type zipIterator[T any, U any] struct {
	first Iterator[T]
	last  Iterator[U]
}

func (a *zipIterator[T, U]) Next() option.Option[Pair[T, U]] {
	if v1, ok1 := a.first.Next().Val(); ok1 {
		if v2, ok2 := a.last.Next().Val(); ok2 {
			return option.Some(Pair[T, U]{v1, v2})
		}
	}
	return option.None[Pair[T, U]]()
}
