package gollection

func LinkedStackOf[T any](elements ...T) LinkedStack[T] {
	var inner = &linkedStack[T]{0, nil}
	var stack = LinkedStack[T]{inner}
	for _, v := range elements {
		stack.Push(v)
	}
	return stack
}

func LinkedStackFrom[T any](collection Collection[T]) LinkedStack[T] {
	var stack = LinkedStackOf[T]()
	ForEach(stack.Push, collection.Iter())
	return stack
}

type LinkedStack[T any] struct {
	inner *linkedStack[T]
}

type linkedStack[T any] struct {
	length int
	first  *oneWayNode[T]
}

type oneWayNode[T any] struct {
	value T
	next  *oneWayNode[T]
}

func (a LinkedStack[T]) Count() int {
	return a.inner.length
}

func (a LinkedStack[T]) IsEmpty() bool {
	return a.inner.length == 0
}

func (a LinkedStack[T]) Push(element T) {
	if a.inner.first == nil {
		a.inner.first = &oneWayNode[T]{element, nil}
	} else {
		a.inner.first = &oneWayNode[T]{element, a.inner.first}
	}
	a.inner.length++
}

func (a LinkedStack[T]) Pop() T {
	if v, ok := a.TryPop().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedStack[T]) Peek() T {
	if v, ok := a.TryPeek().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedStack[T]) TryPop() Option[T] {
	if a.inner.first == nil {
		return None[T]()
	}
	a.inner.length--
	var item = a.inner.first.value
	a.inner.first = a.inner.first.next
	return Some(item)
}

func (a LinkedStack[T]) TryPeek() Option[T] {
	if a.inner.first == nil {
		return None[T]()
	}
	return Some(a.inner.first.value)
}

func (a LinkedStack[T]) Iter() Iterator[T] {
	return &linkedStackIterator[T]{a.inner.first}
}

func (a LinkedStack[T]) ToSlice() []T {
	var arr = make([]T, a.Count())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a.Iter())
	return arr
}

func (a LinkedStack[T]) Clone() LinkedStack[T] {
	return LinkedStackFrom[T](a)
}

func (a LinkedStack[T]) Clear() {
	a.inner.length = 0
	a.inner.first = nil
}

type linkedStackIterator[T any] struct {
	current *oneWayNode[T]
}

func (a *linkedStackIterator[T]) Next() Option[T] {
	if a.current != nil {
		var item = a.current.value
		a.current = a.current.next
		return Some(item)
	}
	return None[T]()
}

func CollectToLinkedStack[T any](it Iterator[T]) LinkedStack[T] {
	var r = LinkedStackOf[T]()
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		r.Push(v)
	}
	return r
}
