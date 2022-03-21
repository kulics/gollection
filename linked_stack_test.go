package gollection

import "testing"

func TestLinkedStack(t *testing.T) {
	var stack = LinkedStackOf[int]()
	if stack.Size() != 0 {
		t.Fatal("stack size not eq 0")
	}
	stack.Push(1)
	if stack.Size() != 1 {
		t.Fatal("stack size not eq 1")
	}
	if stack.Peek() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if stack.Size() != 1 {
		t.Fatal("stack size not eq 1")
	}
	if v := stack.Pop(); v != 1 {
		t.Fatal("element of top is not 1")
	}
	if stack.TryPop().IsSome() {
		t.Fatal("stack must has not element")
	}
	for i := 0; i <= 10; i++ {
		stack.Push(i)
	}
	if stack.Size() != 11 {
		t.Fatal("stack size not eq 11")
	}
	stack = stack.Clone()
	if stack.Size() != 11 {
		t.Fatal("stack size not eq 11")
	}
	stack.Clear()
	if stack.Size() != 0 {
		t.Fatal("stack size not eq 0")
	}
	var slice = stack.ToSlice()
	if len(slice) != 0 {
		t.Fatal("ToSlice size not eq to 0")
	}
	var stackB = LinkedStackFrom[int](ArrayListOf(1, 2, 3))
	if stackB.Size() != 3 {
		t.Fatal("stack size not eq 3")
	}
	var iter = stackB.Iter()
	for i := 3; i >= 1; i-- {
		var item = iter.Next()
		if i != item.value {
			t.Fatal("element error")
		}
	}
}
