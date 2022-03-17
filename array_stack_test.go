package gollection

import "testing"

func TestArrayStack(t *testing.T) {
	var stack = ArrayStackOf[int]()
	if stack.Size() != 0 {
		t.Fatal("array stack size not eq 0")
	}
	if stack.Capacity() != defaultElementsSize {
		t.Fatal("array stack capacity not eq defaultElementsSize")
	}
	stack.Push(1)
	if stack.Size() != 1 {
		t.Fatal("array stack size not eq 1")
	}
	if stack.Capacity() != defaultElementsSize {
		t.Fatal("array stack capacity not eq defaultElementsSize")
	}
	if stack.Peek() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if stack.Size() != 1 {
		t.Fatal("array stack size not eq 1")
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
		t.Fatal("array stack size not eq 11")
	}
	if stack.Capacity() != 15 {
		t.Fatal("array stack capacity not grow *1.5")
	}
	stack = stack.Clone()
	if stack.Size() != 11 {
		t.Fatal("array stack size not eq 11")
	}
	if stack.Capacity() != 15 {
		t.Fatal("array stack capacity not grow *1.5")
	}
	stack.Clear()
	if stack.Size() != 0 {
		t.Fatal("array stack size not eq 0")
	}
	if stack.Capacity() != 15 {
		t.Fatal("array stack capacity not grow *1.5")
	}
	stack.Reserve(10)
	if stack.Size() != 0 {
		t.Fatal("array stack size not eq 0")
	}
	if stack.Capacity() != 15 {
		t.Fatal("array stack capacity not grow *1.5")
	}
	stack.Reserve(30)
	if stack.Size() != 0 {
		t.Fatal("array stack size not eq 0")
	}
	if stack.Capacity() != 30 {
		t.Fatal("array stack capacity not grow to 30")
	}
	var slice = stack.ToSlice()
	if len(slice) != 0 {
		t.Fatal("ToSlice size not eq to 0")
	}
	var stackB = ArrayStackFrom[int](ArrayListOf(1, 2, 3))
	if stackB.Size() != 3 {
		t.Fatal("array list size not eq 3")
	}
	if stackB.Capacity() != 3 {
		t.Fatal("array list capacity not eq 3")
	}
	var iter = stackB.Iter()
	for i := 3; i >= 1; i-- {
		var item = iter.Next()
		if i != item.value {
			t.Fatal("element error")
		}
	}
}
