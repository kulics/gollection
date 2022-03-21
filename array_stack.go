package gollection

func ArrayStackOf[T any](elements ...T) ArrayStack[T] {
	var size = len(elements)
	var stack = MakeArrayStack[T](size)
	copy(stack.inner.elements, elements)
	stack.inner.size = size
	return stack
}

func MakeArrayStack[T any](capacity int) ArrayStack[T] {
	if capacity < defaultElementsSize {
		capacity = defaultElementsSize
	}
	var inner = &arrayStack[T]{make([]T, capacity), 0}
	return ArrayStack[T]{inner}
}

func ArrayStackFrom[T any](collection Collection[T]) ArrayStack[T] {
	var inner = &arrayStack[T]{collection.ToSlice(), collection.Size()}
	return ArrayStack[T]{inner}
}

type ArrayStack[T any] struct {
	inner *arrayStack[T]
}

type arrayStack[T any] struct {
	elements []T
	size     int
}

func (a ArrayStack[T]) Size() int {
	return a.inner.size
}

func (a ArrayStack[T]) IsEmpty() bool {
	return a.inner.size == 0
}

func (a ArrayStack[T]) Push(element T) {
	if growSize := a.inner.size + 1; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	a.inner.elements[a.inner.size] = element
	a.inner.size++
}

func (a ArrayStack[T]) Pop() T {
	if v, ok := a.TryPop().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a ArrayStack[T]) Peek() T {
	if v, ok := a.TryPeek().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a ArrayStack[T]) TryPop() Option[T] {
	if a.IsEmpty() {
		return None[T]()
	}
	var index = a.inner.size - 1
	var item = a.inner.elements[index]
	var empty T
	a.inner.elements[index] = empty
	a.inner.size--
	return Some(item)
}

func (a ArrayStack[T]) TryPeek() Option[T] {
	if a.IsEmpty() {
		return None[T]()
	}
	return Some(a.inner.elements[a.inner.size-1])
}

func (a ArrayStack[T]) Iter() Iterator[T] {
	return &arrayStackIterator[T]{a.Size(), a}
}

func (a ArrayStack[T]) ToSlice() []T {
	var arr = make([]T, a.Size())
	copy(arr, a.inner.elements)
	return arr
}

func (a ArrayStack[T]) Clone() ArrayStack[T] {
	var elements = make([]T, len(a.inner.elements))
	copy(elements, a.inner.elements)
	var inner = &arrayStack[T]{
		elements: elements,
		size:     a.inner.size,
	}
	return ArrayStack[T]{inner}
}

func (a ArrayStack[T]) Reserve(additional int) {
	if addable := len(a.inner.elements) - a.inner.size; addable < additional {
		a.grow(a.inner.size + additional)
	}
}

func (a ArrayStack[T]) Capacity() int {
	return len(a.inner.elements)
}

func (a ArrayStack[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.inner.size; i++ {
		a.inner.elements[i] = emptyValue
	}
	a.inner.size = 0
}

func (a ArrayStack[T]) grow(minCapacity int) {
	var newSize = arrayGrow(len(a.inner.elements))
	if newSize < minCapacity {
		newSize = minCapacity
	}
	var newSource = make([]T, newSize)
	copy(newSource, a.inner.elements)
	a.inner.elements = newSource
}

type arrayStackIterator[T any] struct {
	index  int
	source ArrayStack[T]
}

func (a *arrayStackIterator[T]) Next() Option[T] {
	if a.index > 0 {
		a.index--
		return Some(a.source.inner.elements[a.index])
	}
	return None[T]()
}

func (a *arrayStackIterator[T]) Iter() Iterator[T] {
	return a
}
