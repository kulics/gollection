package list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	var list = Of[int]()
	if list.Count() != 0 {
		t.Fatal("list count not eq 0")
	}
	list.AddLast(1)
	if list.Count() != 1 {
		t.Fatal("list count not eq 1")
	}
	if list.First().Get() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if v := list.First().Set(2); v != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if list.First().Get() != 2 {
		t.Fatal("element of index 0 is not 2")
	}
	for i := 0; i < 10; i++ {
		list.AddLast(i)
	}
	if list.Count() != 11 {
		t.Fatal("list count not eq 11")
	}
	list = list.Clone()
	if list.Count() != 11 {
		t.Fatal("list count not eq 11")
	}
	list.Clear()
	if list.Count() != 0 {
		t.Fatal("list count not eq 0")
	}
	var listB = From[int](Of(1, 2, 3))
	if listB.Count() != 3 {
		t.Fatal("list count not eq 3")
	}
	list.AddAll(0, listB)
	if list.Count() != 3 {
		t.Fatal("list count not eq 3")
	}
	var it = list.Iterator()
	for i := 1; i <= 3; i++ {
		if item, ok := it.Next().Val(); ok && i != item {
			t.Fatal("element error")
		}
	}
	list.AddAll(0, Of(1, 2, 3))
	if list.Count() != 6 {
		t.Fatal("list count not eq 6")
	}
	if list.First().Get() != 1 {
		t.Fatal("first elements of list is not 1")
	}
	if list.RemoveFirst().OrPanic() != 1 {
		t.Fatal("remove list first is not 1")
	}
	if list.First().Get() != 2 {
		t.Fatal("first elements of list is not 2")
	}
	if list.Last().Get() != 3 {
		t.Fatal("last elements of list is not 3")
	}
	if list.RemoveLast().OrPanic() != 3 {
		t.Fatal("remove list last is not 3")
	}
	if list.Last().Get() != 2 {
		t.Fatal("last elements of list is not 2")
	}
	list = Of(1, 2, 3, 1, 2, 3)
	if list.Count() != 6 {
		t.Fatal("list count not eq 6")
	}
}
