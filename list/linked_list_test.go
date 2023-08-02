package list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	var list = LinkedListOf[int]()
	if list.Count() != 0 {
		t.Fatal("list count not eq 0")
	}
	list.PushBack(1)
	if list.Count() != 1 {
		t.Fatal("list count not eq 1")
	}
	if list.PeekFront().Get() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if v := list.PeekFront().Set(2); v != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if list.PeekFront().Get() != 2 {
		t.Fatal("element of index 0 is not 2")
	}
	for i := 0; i < 10; i++ {
		list.PushBack(i)
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
	var listB = LinkedListFrom[int](ArrayListOf(1, 2, 3))
	if listB.Count() != 3 {
		t.Fatal("list count not eq 3")
	}
	list.insertAll(0, listB)
	if list.Count() != 3 {
		t.Fatal("list count not eq 3")
	}
	var it = list.Iterator()
	for i := 1; i <= 3; i++ {
		if item, ok := it.Next().Val(); ok && i != item {
			t.Fatal("element error")
		}
	}
	list.insertAll(0, ArrayListOf(1, 2, 3))
	if list.Count() != 6 {
		t.Fatal("list count not eq 6")
	}
	if list.PeekFront().Get() != 1 {
		t.Fatal("first elements of list is not 1")
	}
	if list.PopFront().OrPanic() != 1 {
		t.Fatal("remove list first is not 1")
	}
	if list.PeekFront().Get() != 2 {
		t.Fatal("first elements of list is not 2")
	}
	if list.PeekBack().Get() != 3 {
		t.Fatal("last elements of list is not 3")
	}
	if list.PopBack().OrPanic() != 3 {
		t.Fatal("remove list last is not 3")
	}
	if list.PeekBack().Get() != 2 {
		t.Fatal("last elements of list is not 2")
	}
	list = LinkedListOf(1, 2, 3, 1, 2, 3)
	if list.Count() != 6 {
		t.Fatal("list count not eq 6")
	}
	var _ BiList[int] = list
}
