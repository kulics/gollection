package gollection

func HashSetOf[T comparable](elements ...T) *HashSet[T] {
	var length = len(elements)
	var set = MakeHashSet[T](length)
	for _, v := range elements {
		set.Put(v)
	}
	return set
}

func MakeHashSet[T comparable](capacity int) *HashSet[T] {
	return (*HashSet[T])(MakeHashMap[T, Void](capacity))
}

func MakeHashSetWithHasher[T comparable](hasher func(data T) uint64, capacity int) *HashSet[T] {
	return (*HashSet[T])(MakeHashMapWithHasher[T, Void](hasher, capacity))
}

func HashSetFrom[T comparable](collection Collection[T]) *HashSet[T] {
	var length = collection.Count()
	var set = MakeHashSet[T](length)
	ForEach(func(t T) {
		set.Put(t)
	}, collection.Iter())
	return set
}

type HashSet[T comparable] HashMap[T, Void]

func (a *HashSet[T]) Count() int {
	return (*HashMap[T, Void])(a).Count()
}

func (a *HashSet[T]) IsEmpty() bool {
	return (*HashMap[T, Void])(a).IsEmpty()
}

func (a *HashSet[T]) Put(element T) bool {
	return (*HashMap[T, Void])(a).Put(element, Void{}).IsSome()
}

func (a *HashSet[T]) PutAll(elements Collection[T]) {
	var iter = elements.Iter()
	for item, ok := iter.Next().Get(); ok; item, ok = iter.Next().Get() {
		a.Put(item)
	}
}

func (a *HashSet[T]) Remove(element T) Option[T] {
	if (*HashMap[T, Void])(a).Remove(element).IsSome() {
		Some(element)
	}
	return None[T]()
}

func (a *HashSet[T]) Contains(element T) bool {
	return (*HashMap[T, Void])(a).Contains(element)
}

func (a *HashSet[T]) ContainsAll(elements Collection[T]) bool {
	var iter = elements.Iter()
	for item, ok := iter.Next().Get(); ok; item, ok = iter.Next().Get() {
		if !a.Contains(item) {
			return false
		}
	}
	return true
}

func (a *HashSet[T]) Clear() {
	(*HashMap[T, Void])(a).Clear()
}

func (a *HashSet[T]) Iter() Iterator[T] {
	return (*hashSetIterator[T])(&hashMapIterator[T, Void]{-1, (*HashMap[T, Void])(a)})
}

func (a *HashSet[T]) ToSlice() []T {
	var arr = make([]T, a.Count())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a.Iter())
	return arr
}

func (a *HashSet[T]) Clone() *HashSet[T] {
	return (*HashSet[T])((*HashMap[T, Void])(a).Clone())
}

type hashSetIterator[T comparable] hashMapIterator[T, Void]

func (a *hashSetIterator[T]) Next() Option[T] {
	var item = (*hashMapIterator[T, Void])(a).Next()
	if v, ok := item.Get(); ok {
		return Some(v.First)
	}
	return None[T]()
}

func CollectToHashSet[T comparable](it Iterator[T]) *HashSet[T] {
	var r = HashSetOf[T]()
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		r.Put(v)
	}
	return r
}
