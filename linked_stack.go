package main

func LinkedStackOf[T any](elements ...T) LinkedStack[T] {
	var stack = LinkedStack[T]{&struct {
		size int
		head *Node[T]
	}{0, nil}}
	for _, v := range elements {
		stack.Push(v)
	}
	return stack
}

func LinkedStackFrom[T any](collection Collection[T]) LinkedStack[T] {
	var stack = LinkedStack[T]{&struct {
		size int
		head *Node[T]
	}{0, nil}}
	ForEach(stack.Push, collection)
	return stack
}

type LinkedStack[T any] struct {
	inner *struct {
		size int
		head *Node[T]
	}
}

func (a LinkedStack[T]) Size() int {
	return a.inner.size
}

func (a LinkedStack[T]) IsEmpty() bool {
	return a.inner.size == 0
}

func (a LinkedStack[T]) Push(element T) {
	if a.inner.head == nil {
		a.inner.head = &Node[T]{element, nil}
	} else {
		a.inner.head = &Node[T]{element, a.inner.head}
	}
	a.inner.size++
}

func (a LinkedStack[T]) Pop() (value T, ok bool) {
	if a.inner.head == nil {
		return
	}
	a.inner.size--
	var item = a.inner.head.Value
	a.inner.head = a.inner.head.Next
	return item, true
}

func (a LinkedStack[T]) Peek() (value T, ok bool) {
	if a.inner.head == nil {
		return
	}
	return a.inner.head.Value, true
}
