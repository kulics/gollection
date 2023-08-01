package set

import (
	"github.com/kulics/gollection/dict"
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

func HashSetOf[T comparable](elements ...T) *HashSet[T] {
	var length = len(elements)
	var set = MakeHashSet[T](length)
	for _, v := range elements {
		set.Put(v)
	}
	return set
}

func MakeHashSet[T comparable](capacity int) *HashSet[T] {
	return (*HashSet[T])(dict.MakeHashDict[T, util.Void](capacity))
}

func MakeHashSetWithHasher[T comparable](hasher func(data T) uint64, capacity int) *HashSet[T] {
	return (*HashSet[T])(dict.MakeHashDictWithHasher[T, util.Void](hasher, capacity))
}

func HashSetFrom[T comparable](collection iter.Collection[T]) *HashSet[T] {
	var length = collection.Count()
	var set = MakeHashSet[T](length)
	iter.ForEach(func(t T) {
		set.Put(t)
	}, collection.Iterator())
	return set
}

type HashSet[T comparable] dict.HashDict[T, util.Void]

func (a *HashSet[T]) Count() int {
	return (*dict.HashDict[T, util.Void])(a).Count()
}

func (a *HashSet[T]) Put(element T) bool {
	return (*dict.HashDict[T, util.Void])(a).Put(element, util.Void{}).IsSome()
}

func (a *HashSet[T]) Remove(element T) util.Opt[T] {
	if (*dict.HashDict[T, util.Void])(a).Remove(element).IsSome() {
		util.Some(element)
	}
	return util.None[T]()
}

func (a *HashSet[T]) Contains(element T) bool {
	return (*dict.HashDict[T, util.Void])(a).Contains(element)
}

func (a *HashSet[T]) Clear() {
	(*dict.HashDict[T, util.Void])(a).Clear()
}

func (a *HashSet[T]) Iterator() iter.Iterator[T] {
	return &hashSetIterator[T]{(*dict.HashDict[T, util.Void])(a).Iterator()}
}

func (a *HashSet[T]) Clone() *HashSet[T] {
	return (*HashSet[T])((*dict.HashDict[T, util.Void])(a).Clone())
}

type hashSetIterator[T comparable] struct {
	it iter.Iterator[util.Pair[T, util.Void]]
}

func (a *hashSetIterator[T]) Next() util.Opt[T] {
	if v, ok := a.it.Next().Val(); ok {
		return util.Some(v.First)
	}
	return util.None[T]()
}

func HashSetCollector[T comparable]() iter.Collector[*HashSet[T], T, *HashSet[T]] {
	return hashSetCollector[T]{}
}

type hashSetCollector[T comparable] struct{}

func (a hashSetCollector[T]) Builder() *HashSet[T] {
	return MakeHashSet[T](10)
}

func (a hashSetCollector[T]) Append(supplier *HashSet[T], element T) {
	supplier.Put(element)
}

func (a hashSetCollector[T]) Finish(supplier *HashSet[T]) *HashSet[T] {
	return supplier
}
