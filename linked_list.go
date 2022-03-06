package main

func LinkedListOf[T any](elements ...T) LinkedList[T] {
	var list = LinkedList[T]{&struct {
		size int
		head *Node[T]
	}{0, nil}}
	for _, v := range elements {
		list.Prepend(v)
	}
	return list
}

func LinkedListFrom[T any](collection Collection[T]) LinkedList[T] {
	var list = LinkedList[T]{&struct {
		size int
		head *Node[T]
	}{0, nil}}
	ForEach(list.Prepend, collection)
	return list
}

type LinkedList[T any] struct {
	inner *struct {
		size int
		head *Node[T]
	}
}

func (a LinkedList[T]) Prepend(element T) {
	PrependNode(a.inner.head, element)
	a.inner.size++
}

func (a LinkedList[T]) Append(element T) {
	AppendNode(a.inner.head, element)
	a.inner.size++
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
		InsertNode(a.inner.head.Next, a.inner.head, index-1, element)
	}
	return true
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
		item = RemoveNode(a.inner.head.Next, a.inner.head, index-1)
	}
	a.inner.size--
	return item, true
}

func (a LinkedList[T]) Get(index int) (value T, ok bool) {
	if a.isOutOfBounds(index) {
		return
	}
	return GetNode(a.inner.head, index), true
}

func (a LinkedList[T]) GetOrPanic(index int) T {
	if a.isOutOfBounds(index) {
		panic("out of bounds")
	}
	return GetNode(a.inner.head, index)
}

func (a LinkedList[T]) Set(index int, newElement T) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	return Some(SetNode(a.inner.head, index, newElement))
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
	return &LinkedListIterator[T]{a.inner.head}
}

type LinkedListIterator[T any] struct {
	current *Node[T]
}

func (a *LinkedListIterator[T]) Next() Option[T] {
	if a.current != nil {
		var item = a.current.Value
		a.current = a.current.Next
		return Some(item)
	}
	return None[T]()
}

func (a *LinkedListIterator[T]) Iter() Iterator[T] {
	return a
}
