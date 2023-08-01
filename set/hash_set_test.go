package set

import "testing"

func TestHashSet(t *testing.T) {
	var hashset = HashSetOf[int]()
	var _ Set[int] = hashset
}
