package dict

import (
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

// Collection's extended interfaces, can provide more functional abstraction for maps.
type Dict[K any, V any] interface {
	iter.Collection[util.Pair[K, V]]

	Contains(key K) bool
	At(key K) util.Ref[V]
	Put(key K, value V) util.Opt[V]
	Remove(key K) util.Opt[V]
	Clear()
}

func Equals[K any, V comparable](l Dict[K, V], r Dict[K, V]) bool {
	if l.Count() != r.Count() {
		return false
	}
	var lIter = l.Iterator()
	for {
		if pair, ok := lIter.Next().Val(); ok {
			if v, ok := r.At(pair.First).Val(); !ok || v != pair.Second {
				return false
			}
		} else {
			break
		}
	}
	return true
}
