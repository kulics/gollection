package gollection

import "testing"

func TestHashMap(t *testing.T) {
	var hashmap = HashMapOf[string, int]()
	if hashmap.Size() != 0 {
		t.Fatal("map size not eq 0")
	}
	hashmap.Put("1", 1)
	if hashmap.Size() != 1 {
		t.Fatal("map size not eq 1")
	}
	hashmap.Put("1", 2)
	if hashmap.Size() != 1 {
		t.Fatal("map size not eq 1")
	}
	if v, ok := hashmap.TryGet("1").Get(); !ok || v != 2 {
		t.Fatal("map value not eq 2")
	}
}
