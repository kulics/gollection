package main

func ArrayStackOf[T any](elements ...T) ArrayStack[T] {
	var array = make([]T, len(elements))
	copy(array, elements)
	return ArrayStack[T]{&struct {
		elements []T
		size     int
	}{array, len(elements)}}
}

func MakeArrayStack[T any](capacity int) ArrayStack[T] {
	return ArrayStack[T]{&struct {
		elements []T
		size     int
	}{make([]T, capacity), 0}}
}

func ArrayStackFrom[T any](collection Collection[T]) ArrayStack[T] {
	var size = collection.Size()
	var array = make([]T, size)
	ForEach(func(item Pair[int, T]) {
		array[item.First] = item.Second
	}, WithIndex[T](collection))
	return ArrayStack[T]{&struct {
		elements []T
		size     int
	}{array, size}}
}

type ArrayStack[T any] struct {
	inner *struct {
		elements []T
		size     int
	}
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

func (a ArrayStack[T]) Pop() (value T, ok bool) {
	if a.IsEmpty() {
		return
	}
	var index = a.inner.size - 1
	var item = a.inner.elements[index]
	var empty T
	a.inner.elements[index] = empty
	a.inner.size--
	return item, true
}

func (a ArrayStack[T]) Peek() (value T, ok bool) {
	if a.IsEmpty() {
		return
	}
	return a.inner.elements[a.inner.size-1], true
}

func (a ArrayStack[T]) grow() {
	var newSource = make([]T, int(float64(len(a.inner.elements))*1.5))
	copy(newSource, a.inner.elements)
	a.inner.elements = newSource
}
