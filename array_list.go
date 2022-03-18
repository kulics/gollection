package gollection

const defaultElementsSize = 10

func ArrayListOf[T any](elements ...T) ArrayList[T] {
	var size = len(elements)
	var list = MakeArrayList[T](size)
	copy(list.inner.elements, elements)
	list.inner.size = size
	return list
}

func MakeArrayList[T any](capacity int) ArrayList[T] {
	if capacity < defaultElementsSize {
		capacity = defaultElementsSize
	}
	var inner = &arrayList[T]{make([]T, capacity), 0}
	return ArrayList[T]{inner}
}

func ArrayListFrom[T any, I Collection[T]](collection I) ArrayList[T] {
	var inner = &arrayList[T]{collection.ToSlice(), collection.Size()}
	return ArrayList[T]{inner}
}

type ArrayList[T any] struct {
	inner *arrayList[T]
}

type arrayList[T any] struct {
	elements []T
	size     int
}

func (a ArrayList[T]) Prepend(element T) {
	if growSize := a.inner.size + 1; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	copy(a.inner.elements[1:], a.inner.elements[0:])
	a.inner.elements[0] = element
}

func (a ArrayList[T]) PrependAll(elements Collection[T]) {
	var additional = elements.Size()
	if growSize := a.inner.size + additional; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	copy(a.inner.elements[additional:], a.inner.elements[0:])
	var i = 0
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.size++
		i++
	}, elements.Iter())
}

func (a ArrayList[T]) Append(element T) {
	if growSize := a.inner.size + 1; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	a.inner.elements[a.inner.size] = element
	a.inner.size++
}

func (a ArrayList[T]) AppendAll(elements Collection[T]) {
	var additional = elements.Size()
	if growSize := a.inner.size + additional; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	var i = a.inner.size
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.size++
		i++
	}, elements.Iter())
}

func (a ArrayList[T]) Insert(index int, element T) {
	if index < 0 || index > a.inner.size {
		panic(OutOfBounds)
	}
	if growSize := a.inner.size + 1; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	copy(a.inner.elements[index+1:], a.inner.elements[index:])
	a.inner.elements[index] = element
}

func (a ArrayList[T]) InsertAll(index int, elements Collection[T]) {
	if index < 0 || index > a.inner.size {
		panic(OutOfBounds)
	}
	var additional = elements.Size()
	if growSize := a.inner.size + additional; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	copy(a.inner.elements[index+additional:], a.inner.elements[index:])
	var i = index
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.size++
		i++
	}, elements.Iter())
}

func (a ArrayList[T]) Remove(index int) T {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var removed = a.inner.elements[index]
	copy(a.inner.elements[:index], a.inner.elements[index+1:])
	var emptyValue T
	a.inner.elements[a.inner.size-1] = emptyValue
	a.inner.size--
	return removed
}

func (a ArrayList[T]) Reserve(additional int) {
	if addable := len(a.inner.elements) - a.inner.size; addable < additional {
		a.grow(a.inner.size + additional)
	}
}

func (a ArrayList[T]) Get(index int) T {
	if v, ok := a.TryGet(index).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a ArrayList[T]) Set(index int, newElement T) T {
	if v, ok := a.TrySet(index, newElement).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a ArrayList[T]) GetAndSet(index int, set func(oldElement T) T) Pair[T, T] {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var oldElement = a.inner.elements[index]
	var newElement = set(oldElement)
	return PairOf(newElement, oldElement)
}

func (a ArrayList[T]) TryGet(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	return Some(a.inner.elements[index])
}

func (a ArrayList[T]) TrySet(index int, newElement T) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	var oldElement = a.inner.elements[index]
	a.inner.elements[index] = newElement
	return Some(oldElement)
}

func (a ArrayList[T]) Clear() {
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

func (a ArrayList[T]) ToSlice() []T {
	var slice = make([]T, a.Size())
	copy(slice, a.inner.elements)
	return slice
}

func (a ArrayList[T]) Clone() ArrayList[T] {
	var elements = make([]T, len(a.inner.elements))
	copy(elements, a.inner.elements)
	var inner = &arrayList[T]{
		elements: elements,
		size:     a.inner.size,
	}
	return ArrayList[T]{inner}
}

func (a ArrayList[T]) Capacity() int {
	return len(a.inner.elements)
}

func (a ArrayList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.inner.size {
		return true
	}
	return false
}

func (a ArrayList[T]) grow(minCapacity int) {
	var newSize = arrayGrow(a.inner.size)
	if newSize < minCapacity {
		newSize = minCapacity
	}
	var newSource = make([]T, newSize)
	copy(newSource, a.inner.elements)
	a.inner.elements = newSource
}

func arrayGrow(size int) int {
	var newSize = size + (size >> 1)
	if newSize < defaultElementsSize {
		newSize = defaultElementsSize
	}
	return newSize
}

type arrayListIterator[T any] struct {
	index  int
	source ArrayList[T]
}

func (a *arrayListIterator[T]) Next() Option[T] {
	if a.index < a.source.Size()-1 {
		a.index++
		return Some(a.source.inner.elements[a.index])
	}
	return None[T]()
}
