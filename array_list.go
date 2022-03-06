package main

func ArrayListOf[T any](elements ...T) ArrayList[T] {
	var array = make([]T, len(elements))
	copy(array, elements)
	return ArrayList[T]{ &struct {
		elements []T
		size int
	}{ array, len(elements) }}
}

func MakeArrayList[T any](capacity int) ArrayList[T] {
	return ArrayList[T]{ &struct {
		elements []T
		size int
	}{ make([]T, capacity), 0 }}
}

func ArrayListFrom[T any](collection Collection[T]) ArrayList[T] {
	var size = collection.Size()
	var array = make([]T, size)
	ForEach[Pair[int, T]](func(item Pair[int, T]) {
		array[item.First] = item.Second
	}, WithIndex[T](collection))
	return ArrayList[T]{ &struct {
		elements []T
		size int
	}{ array, size }}
}

type ArrayList[T any] struct {
	inner *struct {
		elements []T
		size int
	}
}

func (a ArrayList[T]) Prepend(element T) {
	if len(a.inner.elements) < a.inner.size + 1 {
		a.grow(1)
	}
	copy(a.inner.elements[1:], a.inner.elements[0:])
	a.inner.elements[0] = element
}

func (a ArrayList[T]) PrependAll(elements Collection[T]) {
	var additional = elements.Size()
	if len(a.inner.elements) < a.inner.size + additional {
		a.grow(additional)
	}
	copy(a.inner.elements[additional:], a.inner.elements[0:])
	var iter = elements.Iter()
	var i = 0
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		a.inner.elements[i] = v
		a.inner.size++
		i++
	}
}

func (a ArrayList[T]) Append(element T) {
	if len(a.inner.elements) < a.inner.size + 1 {
		a.grow(1)
	}
	a.inner.elements[a.inner.size] = element
	a.inner.size++
}

func (a ArrayList[T]) AppendAll(elements Collection[T]) {
	var additional = elements.Size()
	if len(a.inner.elements) < a.inner.size + additional {
		a.grow(additional)
	}
	var iter = elements.Iter()
	var i = a.inner.size
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		a.inner.elements[i] = v
		a.inner.size++
		i++
	}
}

func (a ArrayList[T]) Insert(index int, element T) bool {
	if index < 0 || index > a.inner.size {
		return false
	}
	if len(a.inner.elements) < a.inner.size + 1 {
		a.grow(1)
	}
	copy(a.inner.elements[index + 1:], a.inner.elements[index:])
	a.inner.elements[index] = element
	return true
}

func (a ArrayList[T]) InsertAll(index int, elements Collection[T]) bool {
	if index < 0 || index > a.inner.size {
		return false
	}
	var additional = elements.Size()
	if len(a.inner.elements) < a.inner.size + additional {
		a.grow(additional)
	}
	copy(a.inner.elements[index + additional:], a.inner.elements[index:])
	var iter = elements.Iter()
	var i = index
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		a.inner.elements[i] = v
		a.inner.size++
		i++
	}
	return true
}

func (a ArrayList[T]) Remove(index int) (value T, ok bool) {
	if a.isOutOfBounds(index) {
		return
	}
	var removed = a.inner.elements[index]
	copy(a.inner.elements[:index], a.inner.elements[index+1:])
	var emptyValue T
	a.inner.elements[a.inner.size - 1] = emptyValue
	a.inner.size--
	return removed, true
}

func (a ArrayList[T]) Reserve(additional int) {
	var addable = len(a.inner.elements) - a.inner.size
	if addable < additional {
		a.grow(additional - addable)
	}
}

func (a ArrayList[T]) Get(index int) (value T, ok bool) {
	if a.isOutOfBounds(index) {
		return
	}
	return a.inner.elements[index], true
}

func (a ArrayList[T]) GetOrPanic(index int) T {
	if a.isOutOfBounds(index) {
		panic("out of bounds")
	}
	return a.inner.elements[index]
}

func (a ArrayList[T]) Set(index int, newElement T) (value T, ok bool) {
	if a.isOutOfBounds(index) {
		return
	}
	var oldElement = a.inner.elements[index]
	a.inner.elements[index] = newElement
	return oldElement, true
}

func (a ArrayList[T]) Clean() {
	var emptyValue T
	for i := 0; i < a.inner.size; i++ {
		a.inner.elements[i] = emptyValue
	}
	a.inner.size = 0
}

func (a ArrayList[T]) Size() int {
	return a.inner.size
}

func (a ArrayList[T]) IsEmpty() bool {
	return a.inner.size == 0
}

func (a ArrayList[T]) Iter() Iterator[T] {
	return &arrayListIterator[T]{-1, a}
}

func (a ArrayList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.inner.size {
		return true
	}
	return false
}

func (a ArrayList[T]) grow(minCapacity int) {
	var newSize = int(float64(len(a.inner.elements)) * 1.5)
	if newSize < minCapacity {
		newSize = minCapacity
	}
	var newSource = make([]T, newSize)
	copy(newSource, a.inner.elements)
	a.inner.elements = newSource
}

type arrayListIterator[T any] struct {
	index int
	source ArrayList[T]
}

func (a *arrayListIterator[T]) Next() (value T, ok bool) {
	if a.index < a.source.Size() - 1 {
		a.index++
		return a.source.Get(a.index)
	}
	return
}

func (a *arrayListIterator[T]) Iter() Iterator[T] {
	return a
}
