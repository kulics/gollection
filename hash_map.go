package gollection

import (
	"hash/maphash"
	"unsafe"
)

func defaultHashCode[K comparable](k K) uint64 {
	var h maphash.Hash
	var strKey = *(*string)(unsafe.Pointer(&struct {
		data unsafe.Pointer
		len  int
	}{unsafe.Pointer(&k), int(unsafe.Sizeof(k))}))
	h.WriteString(strKey)
	return h.Sum64()
}

func HashMapOf[K comparable, V any](elements ...Pair[K, V]) HashMap[K, V] {
	var size = len(elements)
	var dict = MakeHashMapWithHasher[K, V](defaultHashCode[K], size)
	for _, v := range elements {
		dict.Put(v.First, v.Second)
	}
	return dict
}

func MakeHashMap[K comparable, V any](capacity int) HashMap[K, V] {
	return MakeHashMapWithHasher[K, V](defaultHashCode[K], capacity)
}

func MakeHashMapWithHasher[K comparable, V any](hasher func(K) uint64, capacity int) HashMap[K, V] {
	var size = capacity
	var buckets = make([]int, bucketsSizeFor(size))
	for i := 0; i < len(buckets); i++ {
		buckets[i] = -1
	}
	if size < defaultElementsSize {
		size = defaultElementsSize
	}
	var inner = &hashMap[K, V]{
		buckets:    buckets,
		entries:    make([]entry[K, V], size),
		hash:       hasher,
		loadFactor: 1,
	}
	return HashMap[K, V]{inner}
}

func HashMapFrom[K comparable, V any](collection Collection[Pair[K, V]]) HashMap[K, V] {
	var size = collection.Size()
	var dict = MakeHashMapWithHasher[K, V](defaultHashCode[K], size)
	ForEach(func(t Pair[K, V]) {
		dict.Put(t.First, t.Second)
	}, collection.Iter())
	return dict
}

func bucketsSizeFor(size int) int {
	var bucketsSize = 16
	for bucketsSize < size {
		bucketsSize = bucketsSize * 2
	}
	return bucketsSize
}

type HashMap[K comparable, V any] struct {
	inner *hashMap[K, V]
}

type hashMap[K comparable, V any] struct {
	buckets     []int
	entries     []entry[K, V]
	appendCount int
	freeCount   int
	freeSize    int
	hash        func(K) uint64
	loadFactor  float64
}

type entry[K any, V any] struct {
	hash  uint64
	key   K
	value V
	next  int
	alive bool
}

func (a HashMap[K, V]) Get(key K) V {
	if v, ok := a.TryGet(key).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a HashMap[K, V]) Put(key K, value V) Option[V] {
	var hash = a.inner.hash(key)
	var index = a.index(hash)
	for i := a.inner.buckets[index]; i >= 0; i = a.inner.entries[i].next {
		var item = a.inner.entries[i]
		if item.hash == hash && item.key == key {
			var newItem = entry[K, V]{
				hash:  item.hash,
				key:   item.key,
				value: value,
				next:  item.next,
				alive: item.alive,
			}
			a.inner.entries[i] = newItem
			return Some(item.value)
		}
	}
	var bucket int
	if a.inner.freeSize > 0 {
		bucket = a.inner.freeCount
		a.inner.freeCount = a.inner.entries[a.inner.freeCount].next
		a.inner.freeSize--
	} else {
		if a.grow(a.Size() + 1) {
			index = a.index(hash)
		}
		bucket = a.inner.appendCount
		a.inner.appendCount++
	}
	var newItem = entry[K, V]{
		hash:  hash,
		key:   key,
		value: value,
		next:  a.inner.buckets[index],
		alive: true,
	}
	a.inner.entries[bucket] = newItem
	a.inner.buckets[index] = bucket
	return None[V]()
}

func (a HashMap[K, V]) PutAll(elements Collection[Pair[K, V]]) {
	var iter = elements.Iter()
	if size, addSize := a.Size(), elements.Size(); size < addSize {
		a.grow(addSize)
	}
	for item, ok := iter.Next().Get(); ok; item, ok = iter.Next().Get() {
		var k, v = item.Get()
		a.Put(k, v)
	}
}

func (a HashMap[K, V]) GetAndPut(key K, set func(oldValue Option[V]) V) Pair[V, Option[V]] {
	var hash = a.inner.hash(key)
	var index = a.index(hash)
	for i := a.inner.buckets[index]; i >= 0; i = a.inner.entries[i].next {
		var item = a.inner.entries[i]
		if item.hash == hash && item.key == key {
			var newValue = set(Some(item.value))
			var newItem = entry[K, V]{
				hash:  item.hash,
				key:   item.key,
				value: newValue,
				next:  item.next,
				alive: item.alive,
			}
			a.inner.entries[i] = newItem
			return PairOf(newValue, Some(item.value))
		}
	}
	var bucket int
	if a.inner.freeSize > 0 {
		bucket = a.inner.freeCount
		a.inner.freeCount = a.inner.entries[a.inner.freeCount].next
		a.inner.freeSize--
	} else {
		if a.grow(a.Size() + 1) {
			index = a.index(hash)
		}
		bucket = a.inner.appendCount
		a.inner.appendCount++
	}
	var newValue = set(None[V]())
	var newItem = entry[K, V]{
		hash:  hash,
		key:   key,
		value: newValue,
		next:  a.inner.buckets[index],
		alive: true,
	}
	a.inner.entries[bucket] = newItem
	a.inner.buckets[index] = bucket
	return PairOf(newValue, None[V]())
}

func (a HashMap[K, V]) TryGet(key K) Option[V] {
	var hash = a.inner.hash(key)
	var index = a.index(hash)
	for i := a.inner.buckets[index]; i >= 0; i = a.inner.entries[i].next {
		var item = a.inner.entries[i]
		if item.hash == hash && item.key == key {
			return Some(item.value)
		}
	}
	return None[V]()
}

func (a HashMap[K, V]) Remove(key K) Option[V] {
	var hash = a.inner.hash(key)
	var index = a.index(hash)
	var last = -1
	for i := a.inner.buckets[index]; i >= 0; i = a.inner.entries[i].next {
		var item = a.inner.entries[i]
		if item.hash == hash && item.key == key {
			if last < 0 {
				a.inner.buckets[index] = a.inner.entries[i].next
			} else {
				var item = a.inner.entries[last]
				item.next = a.inner.entries[i].next
				a.inner.entries[last] = item
			}
			var nilK K
			var nilV V
			var empty = entry[K, V]{
				next:  a.inner.freeCount,
				key:   nilK,
				value: nilV,
			}
			a.inner.entries[i] = empty
			a.inner.freeCount = i
			a.inner.freeCount++
			return Some(item.value)
		}
	}
	return None[V]()
}

func (a HashMap[K, V]) Contains(key K) bool {
	return a.TryGet(key).IsSome()
}

func (a HashMap[K, V]) Size() int {
	return a.inner.appendCount - a.inner.freeSize + 1
}

func (a HashMap[K, V]) IsEmpty() bool {
	return a.Size() == 0
}

func (a HashMap[K, V]) Clear() {
	for i := 0; i < len(a.inner.buckets); i++ {
		a.inner.buckets[i] = -1
	}
	for i := 0; i < len(a.inner.entries); i++ {
		a.inner.entries[i] = entry[K, V]{}
	}
}

func (a HashMap[K, V]) Iter() Iterator[Pair[K, V]] {
	return &hashMapIterator[K, V]{-1, a}
}

func (a HashMap[K, V]) ToSlice() []Pair[K, V] {
	var arr = make([]Pair[K, V], a.Size())
	ForEach(func(t Pair[K, V]) {
		arr = append(arr, t)
	}, a.Iter())
	return arr
}

func (a HashMap[K, V]) Clone() HashMap[K, V] {
	var buckets = make([]int, len(a.inner.buckets))
	copy(buckets, a.inner.buckets)
	var entries = make([]entry[K, V], len(a.inner.entries))
	copy(entries, a.inner.entries)
	var inner = &hashMap[K, V]{
		buckets:     buckets,
		entries:     entries,
		appendCount: a.inner.appendCount,
		freeCount:   a.inner.freeCount,
		freeSize:    a.inner.freeSize,
		hash:        a.inner.hash,
		loadFactor:  a.inner.loadFactor,
	}
	return HashMap[K, V]{inner}
}

func (a HashMap[K, V]) grow(minCapacity int) bool {
	var entriesSize = len(a.inner.entries)
	var bucketsSize = len(a.inner.buckets)
	var isRehash = false
	if float64(minCapacity/bucketsSize) > a.inner.loadFactor {
		var newBucketsSize = bucketsSize * 2
		var newBuckets = make([]int, newBucketsSize)
		for i := 0; i < len(newBuckets); i++ {
			newBuckets[i] = -1
		}
		for i, v := range a.inner.entries {
			if v.alive {
				var bucket = int(v.hash % uint64(newBucketsSize))
				v.next = newBuckets[bucket]
				a.inner.entries[i] = v
				newBuckets[bucket] = i
			}
		}
		a.inner.buckets = newBuckets
		isRehash = true
	}
	if minCapacity > entriesSize {
		var newSize = entriesSize + (entriesSize >> 1)
		if newSize < minCapacity {
			newSize = minCapacity
		}
		var newEntries = make([]entry[K, V], newSize)
		copy(newEntries, a.inner.entries)
		a.inner.entries = newEntries
	}
	return isRehash
}

func (a HashMap[K, V]) index(hash uint64) int {
	return int(hash % uint64(len(a.inner.buckets)))
}

type hashMapIterator[K comparable, V any] struct {
	index  int
	source HashMap[K, V]
}

func (a *hashMapIterator[K, V]) Next() Option[Pair[K, V]] {
	for a.index < len(a.source.inner.entries)-1 {
		a.index++
		var item = a.source.inner.entries[a.index]
		if item.alive {
			return Some(PairOf(item.key, item.value))
		}
	}
	return None[Pair[K, V]]()
}
