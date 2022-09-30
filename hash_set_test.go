package gollection

import "testing"

func TestHashSet(t *testing.T) {
	var hashset = HashSetOf[int]()
	var _ AnySet[int] = hashset
	var _ AnyMutableSet[int] = hashset
}
