package dict

import (
	"hash/maphash"
	"unsafe"

	"github.com/kulics/gollection/option"
	"github.com/kulics/gollection/ref"
	"github.com/kulics/gollection/seq"
)

const defaultElementsLength = 10

func arrayGrow(length int) int {
	var newLength = length + (length >> 1)
	if newLength < defaultElementsLength {
		newLength = defaultElementsLength
	}
	return newLength
}

func defaultHashCode[K comparable]() func(k K) uint64 {
	var h maphash.Hash
	var seed = h.Seed()
	var k K
	switch ((any)(k)).(type) {
	case string:
		return func(key K) uint64 {
			var strKey = *(*string)(unsafe.Pointer(&key))
			h.SetSeed(seed)
			h.WriteString(strKey)
			return h.Sum64()
		}
	default:
		return func(key K) uint64 {
			var strKey = *(*string)(unsafe.Pointer(&struct {
				data unsafe.Pointer
				len  int
			}{unsafe.Pointer(&k), int(unsafe.Sizeof(k))}))
			h.SetSeed(seed)
			h.WriteString(strKey)
			return h.Sum64()
		}
	}
}

func Of[K comparable, V any](elements ...Entry[K, V]) *Dict[K, V] {
	var length = len(elements)
	var dict = MakeWithHasher[K, V](defaultHashCode[K](), length)
	for _, v := range elements {
		dict.Add(v.Key, v.Value)
	}
	return dict
}

func Make[K comparable, V any](capacity int) *Dict[K, V] {
	return MakeWithHasher[K, V](defaultHashCode[K](), capacity)
}

func MakeWithHasher[K comparable, V any](hasher func(K) uint64, capacity int) *Dict[K, V] {
	var length = capacity
	var buckets = make([]int, bucketsLengthFor(length))
	for i := 0; i < len(buckets); i++ {
		buckets[i] = -1
	}
	if length < defaultElementsLength {
		length = defaultElementsLength
	}
	return &Dict[K, V]{
		buckets:    buckets,
		entries:    make([]entry[K, V], length),
		hash:       hasher,
		loadFactor: 1,
		seed:       maphash.MakeSeed(),
	}
}

func From[K comparable, V any](collection seq.Collection[Entry[K, V]]) *Dict[K, V] {
	var length = collection.Count()
	var dict = MakeWithHasher[K, V](defaultHashCode[K](), length)
	seq.ForEach[Entry[K, V]](func(t Entry[K, V]) {
		dict.Add(t.Key, t.Value)
	}, collection)
	return dict
}

func bucketsLengthFor(length int) int {
	var bucketsLength = 16
	for bucketsLength < length {
		bucketsLength = bucketsLength * 2
	}
	return bucketsLength
}

type Dict[K comparable, V any] struct {
	buckets     []int
	entries     []entry[K, V]
	appendCount int
	freeCount   int
	freeLength  int
	hash        func(K) uint64
	loadFactor  float64
	seed        maphash.Seed
}

type entry[K any, V any] struct {
	hash  uint64
	key   K
	value V
	next  int
	alive bool
}

func (a *Dict[K, V]) Count() int {
	return a.appendCount - a.freeLength
}

func (a *Dict[K, V]) Contains(key K) bool {
	return a.At(key).IsNotNil()
}

func (a *Dict[K, V]) At(key K) ref.Ref[V] {
	var hash = a.hash(key)
	var index = a.index(hash)
	for i := a.buckets[index]; i >= 0; i = a.entries[i].next {
		var item = a.entries[i]
		if item.hash == hash && item.key == key {
			return ref.Of(&a.entries[i].value)
		}
	}
	return ref.Of[V](nil)
}

func (a *Dict[K, V]) Add(key K, value V) option.Option[V] {
	var hash = a.hash(key)
	var index = a.index(hash)
	for i := a.buckets[index]; i >= 0; i = a.entries[i].next {
		var item = a.entries[i]
		if item.hash == hash && item.key == key {
			var newItem = entry[K, V]{
				hash:  item.hash,
				key:   item.key,
				value: value,
				next:  item.next,
				alive: item.alive,
			}
			a.entries[i] = newItem
			return option.Some(item.value)
		}
	}
	var bucket int
	if a.freeLength > 0 {
		bucket = a.freeCount
		a.freeCount = a.entries[a.freeCount].next
		a.freeLength--
	} else {
		if a.grow(a.Count() + 1) {
			index = a.index(hash)
		}
		bucket = a.appendCount
		a.appendCount++
	}
	var newItem = entry[K, V]{
		hash:  hash,
		key:   key,
		value: value,
		next:  a.buckets[index],
		alive: true,
	}
	a.entries[bucket] = newItem
	a.buckets[index] = bucket
	return option.None[V]()
}

func (a *Dict[K, V]) Remove(key K) option.Option[V] {
	var hash = a.hash(key)
	var index = a.index(hash)
	var last = -1
	for i := a.buckets[index]; i >= 0; i = a.entries[i].next {
		var item = a.entries[i]
		if item.hash == hash && item.key == key {
			if last < 0 {
				a.buckets[index] = a.entries[i].next
			} else {
				var item = a.entries[last]
				item.next = a.entries[i].next
				a.entries[last] = item
			}
			var nilK K
			var nilV V
			var empty = entry[K, V]{
				next:  a.freeCount,
				key:   nilK,
				value: nilV,
			}
			a.entries[i] = empty
			a.freeCount = i
			a.freeCount++
			return option.Some(item.value)
		}
	}
	return option.None[V]()
}

func (a *Dict[K, V]) Clear() {
	for i := 0; i < len(a.buckets); i++ {
		a.buckets[i] = -1
	}
	for i := 0; i < len(a.entries); i++ {
		a.entries[i] = entry[K, V]{}
	}
}

func (a *Dict[K, V]) Iterator() seq.Iterator[Entry[K, V]] {
	return &hashDictIterator[K, V]{-1, a}
}

func (a *Dict[K, V]) Clone() *Dict[K, V] {
	var buckets = make([]int, len(a.buckets))
	copy(buckets, a.buckets)
	var entries = make([]entry[K, V], len(a.entries))
	copy(entries, a.entries)
	return &Dict[K, V]{
		buckets:     buckets,
		entries:     entries,
		appendCount: a.appendCount,
		freeCount:   a.freeCount,
		freeLength:  a.freeLength,
		hash:        a.hash,
		loadFactor:  a.loadFactor,
	}
}

func (a *Dict[K, V]) grow(minCapacity int) bool {
	var entriesLength = len(a.entries)
	var bucketsLength = len(a.buckets)
	var isRehash = false
	if float64(minCapacity/bucketsLength) > a.loadFactor {
		var newBucketsLength = bucketsLength * 2
		var newBuckets = make([]int, newBucketsLength)
		for i := 0; i < len(newBuckets); i++ {
			newBuckets[i] = -1
		}
		for i, v := range a.entries {
			if v.alive {
				var bucket = int(v.hash % uint64(newBucketsLength))
				v.next = newBuckets[bucket]
				a.entries[i] = v
				newBuckets[bucket] = i
			}
		}
		a.buckets = newBuckets
		isRehash = true
	}
	if minCapacity > entriesLength {
		var newLength = entriesLength + (entriesLength >> 1)
		if newLength < minCapacity {
			newLength = minCapacity
		}
		var newEntries = make([]entry[K, V], newLength)
		copy(newEntries, a.entries)
		a.entries = newEntries
	}
	return isRehash
}

func (a *Dict[K, V]) index(hash uint64) int {
	return int(hash % uint64(len(a.buckets)))
}

type hashDictIterator[K comparable, V any] struct {
	index  int
	source *Dict[K, V]
}

func (a *hashDictIterator[K, V]) Next() option.Option[Entry[K, V]] {
	for a.index < len(a.source.entries)-1 {
		a.index++
		var item = a.source.entries[a.index]
		if item.alive {
			return option.Some(Entry[K, V]{item.key, item.value})
		}
	}
	return option.None[Entry[K, V]]()
}

func Collector[K comparable, V any]() seq.Collector[*Dict[K, V], Entry[K, V], *Dict[K, V]] {
	return collector[K, V]{}
}

type collector[K comparable, V any] struct{}

func (a collector[K, V]) Builder() *Dict[K, V] {
	return Make[K, V](10)
}

func (a collector[K, V]) Append(supplier *Dict[K, V], element Entry[K, V]) {
	supplier.Add(element.Key, element.Value)
}

func (a collector[K, V]) Finish(supplier *Dict[K, V]) *Dict[K, V] {
	return supplier
}

func Equals[K comparable, V comparable](l Dict[K, V], r Dict[K, V]) bool {
	if l.Count() != r.Count() {
		return false
	}
	var lIter = l.Iterator()
	for {
		if pair, ok := lIter.Next().Val(); ok {
			if v, ok := r.At(pair.Key).Val(); !ok || v != pair.Value {
				return false
			}
		} else {
			break
		}
	}
	return true
}
