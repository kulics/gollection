package gollection

func ArrayStackOf[T any](elements ...T) ArrayStack[T] {
	var array = make([]T, len(elements))
	copy(array, elements)
	var inner = &arrayStack[T]{array, len(elements)}
	return ArrayStack[T]{inner}
}

func MakeArrayStack[T any](capacity int) ArrayStack[T] {
	var inner = &arrayStack[T]{make([]T, capacity), 0}
	return ArrayStack[T]{inner}
}

func ArrayStackFrom[T any, I Collection[T]](collection I) ArrayStack[T] {
	var size = collection.Size()
	var array = make([]T, size)
	ForEach(func(item Pair[int, T]) {
		array[item.First] = item.Second
	}, WithIndex[T](collection))
	var inner = &arrayStack[T]{array, size}
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
	if len(a.inner.elements) < a.inner.size+1 {
		a.grow()
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

func (a ArrayStack[T]) grow() {
	var newSource = make([]T, int(float64(len(a.inner.elements))*1.5))
	copy(newSource, a.inner.elements)
	a.inner.elements = newSource
}
