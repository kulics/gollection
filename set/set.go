package set

import (
	"github.com/kulics/gollection/dict"
	"github.com/kulics/gollection/option"
	"github.com/kulics/gollection/seq"
)

func Of[T comparable](elements ...T) *Set[T] {
	var length = len(elements)
	var set = Make[T](length)
	for _, v := range elements {
		set.Add(v)
	}
	return set
}

func Make[T comparable](capacity int) *Set[T] {
	return (*Set[T])(dict.Make[T, void](capacity))
}

func MakeWithHasher[T comparable](hasher func(data T) uint64, capacity int) *Set[T] {
	return (*Set[T])(dict.MakeWithHasher[T, void](hasher, capacity))
}

func From[T comparable](collection seq.Collection[T]) *Set[T] {
	var length = collection.Count()
	var set = Make[T](length)
	seq.ForEach[T](func(t T) {
		set.Add(t)
	}, collection)
	return set
}

type Set[T comparable] dict.Dict[T, void]

func (a *Set[T]) Count() int {
	return (*dict.Dict[T, void])(a).Count()
}

func (a *Set[T]) Add(element T) bool {
	return (*dict.Dict[T, void])(a).Add(element, void{}).IsSome()
}

func (a *Set[T]) Remove(element T) option.Option[T] {
	if (*dict.Dict[T, void])(a).Remove(element).IsSome() {
		option.Some(element)
	}
	return option.None[T]()
}

func (a *Set[T]) Contains(element T) bool {
	return (*dict.Dict[T, void])(a).Contains(element)
}

func (a *Set[T]) ContainsAll(elements seq.Collection[T]) bool {
	var d = (*dict.Dict[T, void])(a)
	var iter = elements.Iterator()
	for item, ok := iter.Next().Val(); ok; item, ok = iter.Next().Val() {
		if !d.Contains(item) {
			return false
		}
	}
	return true
}

func (a *Set[T]) Clear() {
	(*dict.Dict[T, void])(a).Clear()
}

func (a *Set[T]) Iterator() seq.Iterator[T] {
	return &hashSetIterator[T]{(*dict.Dict[T, void])(a).Iterator()}
}

func (a *Set[T]) Clone() *Set[T] {
	return (*Set[T])((*dict.Dict[T, void])(a).Clone())
}

type hashSetIterator[T comparable] struct {
	it seq.Iterator[dict.Entry[T, void]]
}

func (a *hashSetIterator[T]) Next() option.Option[T] {
	if v, ok := a.it.Next().Val(); ok {
		return option.Some(v.Key)
	}
	return option.None[T]()
}

func Collector[T comparable]() seq.Collector[*Set[T], T, *Set[T]] {
	return collector[T]{}
}

type collector[T comparable] struct{}

func (a collector[T]) Builder() *Set[T] {
	return Make[T](10)
}

func (a collector[T]) Append(supplier *Set[T], element T) {
	supplier.Add(element)
}

func (a collector[T]) Finish(supplier *Set[T]) *Set[T] {
	return supplier
}

// Indicates the type of empty.
type void struct{}
