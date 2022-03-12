package set

import (
	. "github.com/kulics/gollection"
	. "github.com/kulics/gollection/dict"
	. "github.com/kulics/gollection/math"
)

func HashSetOf[T comparable](hasher func(data T) int, elements ...T) HashSet[T] {
	var size = len(elements)
	var set = MakeHashSet(hasher, size)
	for _, v := range elements {
		set.Put(v)
	}
	return set
}

func NumberSetOf[T Number](elements ...T) HashSet[T] {
	return HashSetOf(NumberHasher[T], elements...)
}

func StringSetOf[T ~string](elements ...T) HashSet[T] {
	return HashSetOf(StringHasher[T], elements...)
}

func MakeHashSet[T comparable](hasher func(data T) int, capacity int) HashSet[T] {
	return HashSet[T]{MakeHashDict[T, void](hasher, capacity)}
}

func MakeNumberSet[T Number](capacity int) HashSet[T] {
	return MakeHashSet(NumberHasher[T], capacity)
}

func MakeStringSet[T ~string](capacity int) HashSet[T] {
	return MakeHashSet(StringHasher[T], capacity)
}

func HashSetFrom[T comparable, I Collection[T]](hasher func(data T) int, collection I) HashSet[T] {
	var size = collection.Size()
	var set = MakeHashSet(hasher, size)
	ForEach(func(t T) {
		set.Put(t)
	}, collection)
	return set
}

func NumberSetFrom[T Number, I Collection[T]](collection I) HashSet[T] {
	return HashSetFrom(NumberHasher[T], collection)
}

func StringSetFrom[T ~string, I Collection[T]](collection I) HashSet[T] {
	return HashSetFrom(StringHasher[T], collection)
}

type HashSet[T comparable] struct {
	inner HashDict[T, void]
}

type void struct{}

var null = void{}

func (a HashSet[T]) Size() int {
	return a.inner.Size()
}

func (a HashSet[T]) IsEmpty() bool {
	return a.inner.IsEmpty()
}

func (a HashSet[T]) Put(element T) bool {
	return a.inner.Put(element, null).IsSome()
}

func (a HashSet[T]) PutAll(elements Collection[T]) {
	var iter = elements.Iter()
	for item, ok := iter.Next().Get(); ok; item, ok = iter.Next().Get() {
		a.Put(item)
	}
}

func (a HashSet[T]) Remove(element T) bool {
	return a.inner.Remove(element).IsSome()
}

func (a HashSet[T]) Contains(element T) bool {
	return a.inner.Contains(element)
}

func (a HashSet[T]) ContainsAll(elements Collection[T]) bool {
	var iter = elements.Iter()
	for item, ok := iter.Next().Get(); ok; item, ok = iter.Next().Get() {
		if !a.Contains(item) {
			return false
		}
	}
	return true
}

func (a HashSet[T]) Clear() {
	a.inner.Clear()
}

func (a HashSet[T]) Iter() Iterator[T] {
	return &hashSetIterator[T]{a.inner.Iter()}
}

func (a HashSet[T]) ToSlice() []T {
	var arr = make([]T, a.Size())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a)
	return arr
}

type hashSetIterator[T comparable] struct {
	source Iterator[Pair[T, void]]
}

func (a *hashSetIterator[T]) Next() Option[T] {
	var item = a.source.Next()
	if v, ok := item.Get(); ok {
		return Some(v.First)
	}
	return None[T]()
}

func (a *hashSetIterator[T]) Iter() Iterator[T] {
	return a
}
