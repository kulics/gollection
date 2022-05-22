package gollection

const defaultElementsSize = 10

// Constructing an ArrayList with variable-length parameters
func ArrayListOf[T any](elements ...T) ArrayList[T] {
	var size = len(elements)
	var list = MakeArrayList[T](size)
	copy(list.inner.elements, elements)
	list.inner.size = size
	return list
}

// Constructing an empty ArrayList with capacity.
func MakeArrayList[T any](capacity int) ArrayList[T] {
	if capacity < defaultElementsSize {
		capacity = defaultElementsSize
	}
	var inner = &arrayList[T]{make([]T, capacity), 0}
	return ArrayList[T]{inner}
}

// Constructing an ArrayList from other Collection.
func ArrayListFrom[T any](collection Collection[T]) ArrayList[T] {
	var inner = &arrayList[T]{collection.ToSlice(), collection.Size()}
	return ArrayList[T]{inner}
}

// List implemented using Array.
// It has easier in-place modification than the built-in slice.
type ArrayList[T any] struct {
	inner *arrayList[T]
}

type arrayList[T any] struct {
	elements []T
	size     int
}

// Add element at the begin.
func (a ArrayList[T]) Prepend(element T) {
	a.Insert(0, element)
}

// Add multiple elements at the begin.
func (a ArrayList[T]) PrependAll(elements Collection[T]) {
	a.InsertAll(0, elements)
}

// Add element at the end.
func (a ArrayList[T]) Append(element T) {
	if growSize := a.inner.size + 1; len(a.inner.elements) < growSize {
		a.grow(growSize)
	}
	a.inner.elements[a.inner.size] = element
	a.inner.size++
}

// Add multiple elements at the end.
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

// Add element at the index.
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

// Add multiple elements at the index.
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

// Remove element at the index.
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

// Ensure that list have enough space before expansion.
func (a ArrayList[T]) Reserve(additional int) {
	if addable := len(a.inner.elements) - a.inner.size; addable < additional {
		a.grow(a.inner.size + additional)
	}
}

// Return the element at the index.
// A panic is raised when a subscript is out of bounds.
func (a ArrayList[T]) Get(index int) T {
	if v, ok := a.TryGet(index).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

// Set the element at the index, return the old element.
// A panic is raised when a subscript is out of bounds.
func (a ArrayList[T]) Set(index int, newElement T) T {
	if v, ok := a.TrySet(index, newElement).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a ArrayList[T]) Update(index int, update func(oldElement T) T) T {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var oldElement = a.inner.elements[index]
	var newElement = update(oldElement)
	return newElement
}

// Return the element at the index.
// Return None when a subscript is out of bounds.
func (a ArrayList[T]) TryGet(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	return Some(a.inner.elements[index])
}

// Set the element at the index, return the old element.
// Return None when a subscript is out of bounds.
func (a ArrayList[T]) TrySet(index int, newElement T) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	var oldElement = a.inner.elements[index]
	a.inner.elements[index] = newElement
	return Some(oldElement)
}

// Clears all elements, but does not reset the space.
func (a ArrayList[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.inner.size; i++ {
		a.inner.elements[i] = emptyValue
	}
	a.inner.size = 0
}

// Return the size of list.
func (a ArrayList[T]) Size() int {
	return a.inner.size
}

// Return true when the size of list is 0.
func (a ArrayList[T]) IsEmpty() bool {
	return a.inner.size == 0
}

// Return the Iterator of list.
func (a ArrayList[T]) Iter() Iterator[T] {
	return &arrayListIterator[T]{-1, a}
}

// Return a new built-in slice that copies all elements.
func (a ArrayList[T]) ToSlice() []T {
	var slice = make([]T, a.Size())
	copy(slice, a.inner.elements)
	return slice
}

// Return a new list that copies all elements.
func (a ArrayList[T]) Clone() ArrayList[T] {
	var elements = make([]T, len(a.inner.elements))
	copy(elements, a.inner.elements)
	var inner = &arrayList[T]{
		elements: elements,
		size:     a.inner.size,
	}
	return ArrayList[T]{inner}
}

// Return the capacity of list.
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
