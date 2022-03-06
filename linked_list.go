package main

func LinkedListOf[T any](elements ...T) LinkedList[T] {
	var list = LinkedList[T]{ &struct {
		size int
		head *Node[T]
	}{0, nil} }
	for _, v := range elements {
		list.Prepend(v)
	}
	return list
}

func LinkedListFrom[T any](collection Collection[T]) LinkedList[T] {
	var list = LinkedList[T]{ &struct {
		size int
		head *Node[T]
	}{0, nil} }
	ForEach[T](func(item T) { list.Prepend(item) }, collection)
	return list
}

type LinkedList[T any] struct {
	inner *struct {
		size int
		head *Node[T]
	}
}

func (a LinkedList[T]) Prepend(element T) {
	if a.inner.head == nil {
		a.inner.head = &Node[T]{element, nil}
	} else {
		a.inner.head = &Node[T]{element, a.inner.head}
	}
	a.inner.size++
}

func (a LinkedList[T]) Append(element T) {
	addNode(a.inner.head, element)
	a.inner.size++
}

func addNode[T any](n *Node[T], v T) {
	if n == nil {
		*n = Node[T]{v, nil}
	} else {
		addNode(n.Next, v)
	}
}

func (a LinkedList[T]) Insert(index int, element T) bool {
	if index < 0 || index > a.inner.size {
		return false
	}
	if index == 0 {
		a.Prepend(element)
	} else if index == a.inner.size {
		a.Append(element)
	} else {
		insertNode(a.inner.head.Next, a.inner.head, index - 1, element)
	}
	return true
}

func insertNode[T any](n *Node[T], pre *Node[T], i int, e T) {
	if i == 0 {
		pre.Next = &Node[T]{e, n}
	} else {
		insertNode(n.Next, n, i - 1, e)
	}
}

func (a LinkedList[T]) Remove(index int) (value T, ok bool) {
	if a.isOutOfBounds(index) {
		return
	}
	var item T
	if index == 0 {
		var temp = a.inner.head
		a.inner.head = a.inner.head.Next
		temp.Next = nil
		item = temp.Value
	} else {
		item = removeNode(a.inner.head.Next, a.inner.head, index - 1)
	}
	a.inner.size--
	return item, true
}

func removeNode[T any](n *Node[T], pre *Node[T], i int) T {
	if i == 0 {
		var item = n.Value
		pre.Next = n.Next
		n.Next = nil
		return item
	} else {
		return removeNode(n.Next, n, i - 1)
	}
}

func (a LinkedList[T]) Get(index int) (value T, ok bool) {
	if a.isOutOfBounds(index) {
		return
	}
	return getNode(a.inner.head, index), true
}

func (a LinkedList[T]) GetOrPanic(index int) T {
	if a.isOutOfBounds(index) {
		panic("out of bounds")
	}
	return getNode(a.inner.head, index)
}

func getNode[T any](n *Node[T], i int) T {
	if i == 0 {
		return n.Value
	} else {
		return getNode(n.Next, i - 1)
	}
}

func (a LinkedList[T]) Set(index int, newElement T) (value T, ok bool) {
	if a.isOutOfBounds(index) {
		return
	}
	return setNode(a.inner.head, index, newElement), true
}

func setNode[T any](n *Node[T], i int, v T) T {
	if i == 0 {
		var oldValue = n.Value
		n.Value = v
		return oldValue
	} else {
		return setNode(n.Next, i - 1, v)
	}
}

func (a LinkedList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.inner.size {
		return true
	}
	return false
}

func (a LinkedList[T]) Clean() {
	a.inner.head = nil
	a.inner.size = 0
}

func (a LinkedList[T]) Size() int {
	return a.inner.size
}

func (a LinkedList[T]) IsEmpty() bool {
	return a.inner.size == 0
}

func (a LinkedList[T]) Iter() Iterator[T] {
	return &LinkedListIterator[T]{ a.inner.head }
}

type LinkedListIterator[T any] struct {
	current *Node[T]
}

func (a *LinkedListIterator[T]) Next() (value T, ok bool) {
	if a.current != nil {
		var item = a.current.Value
		a.current = a.current.Next
		return item, true
	}
	return
}

func (a *LinkedListIterator[T]) Iter() Iterator[T] {
	return a
}