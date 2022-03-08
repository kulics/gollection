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

func LinkedStackFrom[T any, I Collection[T]](collection I) LinkedStack[T] {
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

func (a LinkedStack[T]) Pop() T {
	if a.inner.head == nil {
		panic(OutOfBounds)
	}
	a.inner.size--
	var item = a.inner.head.Value
	a.inner.head = a.inner.head.Next
	return item
}

func (a LinkedStack[T]) Peek() T {
	if a.inner.head == nil {
		panic(OutOfBounds)
	}
	return a.inner.head.Value
}

func (a LinkedStack[T]) TryPop() Option[T] {
	if a.inner.head == nil {
		return None[T]()
	}
	a.inner.size--
	var item = a.inner.head.Value
	a.inner.head = a.inner.head.Next
	return Some(item)
}

func (a LinkedStack[T]) TryPeek() Option[T] {
	if a.inner.head == nil {
		return None[T]()
	}
	return Some(a.inner.head.Value)
}
