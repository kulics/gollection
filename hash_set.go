package gollection

func HashSetOf[T comparable](elements ...T) HashSet[T] {
	var length = len(elements)
	var set = MakeHashSet[T](length)
	for _, v := range elements {
		set.Put(v)
	}
	return set
}

func MakeHashSet[T comparable](capacity int) HashSet[T] {
	return HashSet[T]{MakeHashMap[T, Void](capacity)}
}

func MakeHashSetWithHasher[T comparable](hasher func(data T) uint64, capacity int) HashSet[T] {
	return HashSet[T]{MakeHashMapWithHasher[T, Void](hasher, capacity)}
}

func HashSetFrom[T comparable](collection Collection[T]) HashSet[T] {
	var length = collection.Count()
	var set = MakeHashSet[T](length)
	ForEach(func(t T) {
		set.Put(t)
	}, collection.Iter())
	return set
}

type HashSet[T comparable] struct {
	inner HashMap[T, Void]
}

func (a HashSet[T]) Count() int {
	return a.inner.Count()
}

func (a HashSet[T]) IsEmpty() bool {
	return a.inner.IsEmpty()
}

func (a HashSet[T]) Put(element T) bool {
	return a.inner.Put(element, Void{}).IsSome()
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
	var arr = make([]T, a.Count())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a.Iter())
	return arr
}

func (a HashSet[T]) Clone() HashSet[T] {
	return HashSet[T]{a.inner.Clone()}
}

type hashSetIterator[T comparable] struct {
	source Iterator[Pair[T, Void]]
}

func (a *hashSetIterator[T]) Next() Option[T] {
	var item = a.source.Next()
	if v, ok := item.Get(); ok {
		return Some(v.First)
	}
	return None[T]()
}

func CollectToHashSet[T comparable](it Iterator[T]) HashSet[T] {
	var r = HashSetOf[T]()
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		r.Put(v)
	}
	return r
}
