package stack

import (
	"testing"
)

func TestLinkedStack(t *testing.T) {
	var stack = LinkedStackOf[int]()
	if stack.Count() != 0 {
		t.Fatal("stack count not eq 0")
	}
	stack.Push(1)
	if stack.Count() != 1 {
		t.Fatal("stack count not eq 1")
	}
	if stack.Peek().Get() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if stack.Count() != 1 {
		t.Fatal("stack count not eq 1")
	}
	if v := stack.Pop().OrPanic(); v != 1 {
		t.Fatal("element of top is not 1")
	}
	if stack.Pop().IsSome() {
		t.Fatal("stack must has not element")
	}
	for i := 0; i <= 10; i++ {
		stack.Push(i)
	}
	if stack.Count() != 11 {
		t.Fatal("stack count not eq 11")
	}
	stack = stack.Clone()
	if stack.Count() != 11 {
		t.Fatal("stack count not eq 11")
	}
	stack.Clear()
	if stack.Count() != 0 {
		t.Fatal("stack count not eq 0")
	}
	var stackB = LinkedStackFrom[int](LinkedStackOf(3, 2, 1))
	if stackB.Count() != 3 {
		t.Fatal("stack count not eq 3")
	}
	var iter = stackB.Iterator()
	for i := 3; i >= 1; i-- {
		if item, ok := iter.Next().Val(); ok && i != item {
			t.Fatal("element error")
		}
	}
	var _ Stack[int] = stack
}
