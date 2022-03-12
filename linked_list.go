package gollection

func LinkedListOf[T any](elements ...T) LinkedList[T] {
	var inner = &linkedList[T]{0, nil}
	var list = LinkedList[T]{inner}
	for _, v := range elements {
		list.Append(v)
	}
	return list
}

func LinkedListFrom[T any, I Collection[T]](collection I) LinkedList[T] {
	var inner = &linkedList[T]{0, nil}
	var list = LinkedList[T]{inner}
	ForEach(list.Append, collection)
	return list
}

type LinkedList[T any] struct {
	inner *linkedList[T]
}

type linkedList[T any] struct {
	size int
	head *Node[T]
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

func (a LinkedList[T]) Remove(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
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
	return Some(item)
}

func (a LinkedList[T]) Get(index int) T {
	if v, ok := a.TryGet(index).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedList[T]) Set(index int, newElement T) T {
	if v, ok := a.TrySet(index, newElement).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedList[T]) GetAndSet(index int, set func(oldElement T) T) Pair[T, T] {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	return GetAndSetNode(a.inner.head, index, set)
}

func (a LinkedList[T]) TryGet(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	return Some(GetNode(a.inner.head, index))
}

func (a LinkedList[T]) TrySet(index int, newElement T) Option[T] {
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

func (a LinkedList[T]) Clear() {
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

func (a LinkedList[T]) ToSlice() []T {
	var arr = make([]T, a.Size())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a)
	return arr
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
