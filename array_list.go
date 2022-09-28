package gollection

// Constructing an ArrayList with variable-length parameters
func ArrayListOf[T any](elements ...T) ArrayList[T] {
	var length = len(elements)
	var list = MakeArrayList[T](length)
	copy(list.inner.elements, elements)
	list.inner.length = length
	return list
}

// Constructing an empty ArrayList with capacity.
func MakeArrayList[T any](capacity int) ArrayList[T] {
	if capacity < defaultElementsLength {
		capacity = defaultElementsLength
	}
	var inner = &arrayList[T]{make([]T, capacity), 0}
	return ArrayList[T]{inner}
}

// Constructing an ArrayList from other Collection.
func ArrayListFrom[T any](collection Collection[T]) ArrayList[T] {
	var inner = &arrayList[T]{collection.ToSlice(), collection.Count()}
	return ArrayList[T]{inner}
}

// List implemented using Array.
// It has easier in-place modification than the built-in slice.
type ArrayList[T any] struct {
	inner *arrayList[T]
}

type arrayList[T any] struct {
	elements []T
	length   int
}

func (a ArrayList[T]) LastIndex() int {
	return a.inner.length - 1
}

func (a ArrayList[T]) GetFirst() T {
	return a.Get(0)
}

func (a ArrayList[T]) TryGetFirst() Option[T] {
	return a.TryGet(0)
}

func (a ArrayList[T]) GetLast() T {
	return a.Get(a.LastIndex())
}

func (a ArrayList[T]) TryGetLast() Option[T] {
	return a.TryGet(a.LastIndex())
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
	if growLength := a.inner.length + 1; len(a.inner.elements) < growLength {
		a.grow(growLength)
	}
	a.inner.elements[a.inner.length] = element
	a.inner.length++
}

// Add multiple elements at the end.
func (a ArrayList[T]) AppendAll(elements Collection[T]) {
	var additional = elements.Count()
	if growLength := a.inner.length + additional; len(a.inner.elements) < growLength {
		a.grow(growLength)
	}
	var i = a.inner.length
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.length++
		i++
	}, elements.Iter())
}

// Add element at the index.
func (a ArrayList[T]) Insert(index int, element T) {
	if index < 0 || index > a.inner.length {
		panic(OutOfBounds)
	}
	if growLength := a.inner.length + 1; len(a.inner.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.inner.elements[index+1:], a.inner.elements[index:])
	a.inner.elements[index] = element
}

// Add multiple elements at the index.
func (a ArrayList[T]) InsertAll(index int, elements Collection[T]) {
	if index < 0 || index > a.inner.length {
		panic(OutOfBounds)
	}
	var additional = elements.Count()
	if growLength := a.inner.length + additional; len(a.inner.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.inner.elements[index+additional:], a.inner.elements[index:])
	var i = index
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.length++
		i++
	}, elements.Iter())
}

// Remove element at the index.
func (a ArrayList[T]) RemoveAt(index int) T {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var removed = a.inner.elements[index]
	copy(a.inner.elements[:index], a.inner.elements[index+1:])
	var emptyValue T
	a.inner.elements[a.inner.length-1] = emptyValue
	a.inner.length--
	return removed
}

func (a ArrayList[T]) RemoveRange(at Range[int]) {
	var begin, end = at.Get()
	if a.isOutOfBounds(begin) || a.isOutOfBounds(end-1) {
		panic(OutOfBounds)
	}
	if end == begin {
		return
	}
	copy(a.inner.elements[begin:], a.inner.elements[end:])
	var emptyValue T
	for i := len(a.inner.elements) - (end - begin); i < len(a.inner.elements); i++ {
		a.inner.elements[i] = emptyValue
		a.inner.length--
	}
}

// Ensure that list have enough space before expansion.
func (a ArrayList[T]) Reserve(additional int) {
	if addable := len(a.inner.elements) - a.inner.length; addable < additional {
		a.grow(a.inner.length + additional)
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
	a.inner.elements[index] = update(a.inner.elements[index])
	return a.inner.elements[index]
}

func (a ArrayList[T]) UpdateAll(update func(oldElement T) T) {
	for i, v := range a.inner.elements {
		a.inner.elements[i] = update(v)
	}
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
	for i := 0; i < a.inner.length; i++ {
		a.inner.elements[i] = emptyValue
	}
	a.inner.length = 0
}

// Return the number of elements of list.
func (a ArrayList[T]) Count() int {
	return a.inner.length
}

// Return true when the number of elements of list is 0.
func (a ArrayList[T]) IsEmpty() bool {
	return a.inner.length == 0
}

// Return the Iterator of list.
func (a ArrayList[T]) Iter() Iterator[T] {
	return &arrayListIterator[T]{-1, a}
}

// Return a new built-in slice that copies all elements.
func (a ArrayList[T]) ToSlice() []T {
	var slice = make([]T, a.Count())
	copy(slice, a.inner.elements)
	return slice
}

// Return a new list that copies all elements.
func (a ArrayList[T]) Clone() ArrayList[T] {
	var elements = make([]T, len(a.inner.elements))
	copy(elements, a.inner.elements)
	var inner = &arrayList[T]{
		elements: elements,
		length:   a.inner.length,
	}
	return ArrayList[T]{inner}
}

// Return the capacity of list.
func (a ArrayList[T]) Capacity() int {
	return len(a.inner.elements)
}

func (a ArrayList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.inner.length {
		return true
	}
	return false
}

func (a ArrayList[T]) grow(minCapacity int) {
	var newLength = arrayGrow(a.inner.length)
	if newLength < minCapacity {
		newLength = minCapacity
	}
	var newSource = make([]T, newLength)
	copy(newSource, a.inner.elements)
	a.inner.elements = newSource
}

type arrayListIterator[T any] struct {
	index  int
	source ArrayList[T]
}

func (a *arrayListIterator[T]) Next() Option[T] {
	if a.index < a.source.Count()-1 {
		a.index++
		return Some(a.source.inner.elements[a.index])
	}
	return None[T]()
}
