package gollection

func LinkedStackOf[T any](elements ...T) *LinkedStack[T] {
	var stack = &LinkedStack[T]{0, nil}
	for _, v := range elements {
		stack.Push(v)
	}
	return stack
}

func LinkedStackFrom[T any](collection Collection[T]) *LinkedStack[T] {
	var stack = LinkedStackOf[T]()
	ForEach(stack.Push, collection.Iter())
	return stack
}

type LinkedStack[T any] struct {
	length int
	first  *oneWayNode[T]
}

type oneWayNode[T any] struct {
	value T
	next  *oneWayNode[T]
}

func (a *LinkedStack[T]) Count() int {
	return a.length
}

func (a *LinkedStack[T]) IsEmpty() bool {
	return a.length == 0
}

func (a *LinkedStack[T]) Push(element T) {
	if a.first == nil {
		a.first = &oneWayNode[T]{element, nil}
	} else {
		a.first = &oneWayNode[T]{element, a.first}
	}
	a.length++
}

// Add multiple elements to the top of the stack.
func (a *LinkedStack[T]) PushAll(elements Collection[T]) {
	ForEach(func(i T) {
		a.Push(i)
	}, elements.Iter())
}

func (a *LinkedStack[T]) Pop() T {
	if v, ok := a.TryPop().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a *LinkedStack[T]) Peek() T {
	if v, ok := a.TryPeek().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a *LinkedStack[T]) TryPop() Option[T] {
	if a.first == nil {
		return None[T]()
	}
	a.length--
	var item = a.first.value
	a.first = a.first.next
	return Some(item)
}

func (a *LinkedStack[T]) TryPeek() Option[T] {
	if a.first == nil {
		return None[T]()
	}
	return Some(a.first.value)
}

func (a *LinkedStack[T]) Iter() Iterator[T] {
	return &linkedStackIterator[T]{a.first}
}

func (a *LinkedStack[T]) ToSlice() []T {
	var arr = make([]T, a.Count())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a.Iter())
	return arr
}

func (a *LinkedStack[T]) Clone() *LinkedStack[T] {
	return LinkedStackFrom[T](a)
}

func (a *LinkedStack[T]) Clear() {
	a.length = 0
	a.first = nil
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

func CollectToLinkedStack[T any](it Iterator[T]) *LinkedStack[T] {
	var r = LinkedStackOf[T]()
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		r.Push(v)
	}
	return r
}
