package gollection

// Constructing an ArrayList with variable-length parameters
func ArrayListOf[T any](elements ...T) *ArrayList[T] {
	var length = len(elements)
	var list = MakeArrayList[T](length)
	copy(list.elements, elements)
	list.length = length
	return list
}

// Constructing an empty ArrayList with capacity.
func MakeArrayList[T any](capacity int) *ArrayList[T] {
	if capacity < defaultElementsLength {
		capacity = defaultElementsLength
	}
	return &ArrayList[T]{make([]T, capacity), 0}
}

// Constructing an ArrayList from other Collection.
func ArrayListFrom[T any](collection Collection[T]) *ArrayList[T] {
	return &ArrayList[T]{collection.ToSlice(), collection.Count()}
}

// List implemented using Array.
// It has easier in-place modification than the built-in slice.
type ArrayList[T any] struct {
	elements []T
	length   int
}

func (a *ArrayList[T]) LastIndex() int {
	return a.length - 1
}

func (a *ArrayList[T]) GetFirst() T {
	return a.Get(0)
}

func (a *ArrayList[T]) TryGetFirst() Option[T] {
	return a.TryGet(0)
}

func (a *ArrayList[T]) GetLast() T {
	return a.Get(a.LastIndex())
}

func (a *ArrayList[T]) TryGetLast() Option[T] {
	return a.TryGet(a.LastIndex())
}

// Add element at the begin.
func (a *ArrayList[T]) Prepend(element T) {
	a.Insert(0, element)
}

// Add multiple elements at the begin.
func (a *ArrayList[T]) PrependAll(elements Collection[T]) {
	a.InsertAll(0, elements)
}

// Add element at the end.
func (a *ArrayList[T]) Append(element T) {
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	a.elements[a.length] = element
	a.length++
}

// Add multiple elements at the end.
func (a *ArrayList[T]) AppendAll(elements Collection[T]) {
	var additional = elements.Count()
	if growLength := a.length + additional; len(a.elements) < growLength {
		a.grow(growLength)
	}
	var i = a.length
	ForEach(func(item T) {
		a.elements[i] = item
		a.length++
		i++
	}, elements.Iter())
}

// Add element at the index.
func (a *ArrayList[T]) Insert(index int, element T) {
	if index < 0 || index > a.length {
		panic(OutOfBounds)
	}
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.elements[index+1:], a.elements[index:])
	a.elements[index] = element
}

// Add multiple elements at the index.
func (a *ArrayList[T]) InsertAll(index int, elements Collection[T]) {
	if index < 0 || index > a.length {
		panic(OutOfBounds)
	}
	var additional = elements.Count()
	if growLength := a.length + additional; len(a.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.elements[index+additional:], a.elements[index:])
	var i = index
	ForEach(func(item T) {
		a.elements[i] = item
		a.length++
		i++
	}, elements.Iter())
}

// Remove element at the index.
func (a *ArrayList[T]) RemoveAt(index int) T {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var removed = a.elements[index]
	copy(a.elements[:index], a.elements[index+1:])
	var emptyValue T
	a.elements[a.length-1] = emptyValue
	a.length--
	return removed
}

func (a *ArrayList[T]) RemoveRange(at Range[int]) {
	var begin, end = at.Get()
	if a.isOutOfBounds(begin) || a.isOutOfBounds(end-1) {
		panic(OutOfBounds)
	}
	if end == begin {
		return
	}
	copy(a.elements[begin:], a.elements[end:])
	var emptyValue T
	for i := len(a.elements) - (end - begin); i < len(a.elements); i++ {
		a.elements[i] = emptyValue
		a.length--
	}
}

// Ensure that list have enough space before expansion.
func (a *ArrayList[T]) Reserve(additional int) {
	if addable := len(a.elements) - a.length; addable < additional {
		a.grow(a.length + additional)
	}
}

// Return the element at the index.
// A panic is raised when a subscript is out of bounds.
func (a *ArrayList[T]) Get(index int) T {
	if v, ok := a.TryGet(index).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

// Set the element at the index, return the old element.
// A panic is raised when a subscript is out of bounds.
func (a *ArrayList[T]) Set(index int, newElement T) T {
	if v, ok := a.TrySet(index, newElement).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

// Return the element at the index.
// Return None when a subscript is out of bounds.
func (a *ArrayList[T]) TryGet(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	return Some(a.elements[index])
}

// Set the element at the index, return the old element.
// Return None when a subscript is out of bounds.
func (a *ArrayList[T]) TrySet(index int, newElement T) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	var oldElement = a.elements[index]
	a.elements[index] = newElement
	return Some(oldElement)
}

// Clears all elements, but does not reset the space.
func (a *ArrayList[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.length; i++ {
		a.elements[i] = emptyValue
	}
	a.length = 0
}

// Return the number of elements of list.
func (a *ArrayList[T]) Count() int {
	return a.length
}

// Return true when the number of elements of list is 0.
func (a *ArrayList[T]) IsEmpty() bool {
	return a.length == 0
}

// Return the Iterator of list.
func (a *ArrayList[T]) Iter() Iterator[T] {
	return &arrayListIterator[T]{-1, a}
}

// Return a new built-in slice that copies all elements.
func (a *ArrayList[T]) ToSlice() []T {
	var slice = make([]T, a.Count())
	copy(slice, a.elements)
	return slice
}

// Return a new list that copies all elements.
func (a *ArrayList[T]) Clone() *ArrayList[T] {
	var elements = make([]T, len(a.elements))
	copy(elements, a.elements)
	return &ArrayList[T]{
		elements: elements,
		length:   a.length,
	}
}

// Return the capacity of list.
func (a *ArrayList[T]) Capacity() int {
	return len(a.elements)
}

func (a *ArrayList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.length {
		return true
	}
	return false
}

func (a *ArrayList[T]) grow(minCapacity int) {
	var newLength = arrayGrow(a.length)
	if newLength < minCapacity {
		newLength = minCapacity
	}
	var newSource = make([]T, newLength)
	copy(newSource, a.elements)
	a.elements = newSource
}

type arrayListIterator[T any] struct {
	index  int
	source *ArrayList[T]
}

func (a *arrayListIterator[T]) Next() Option[T] {
	if a.index < a.source.Count()-1 {
		a.index++
		return Some(a.source.elements[a.index])
	}
	return None[T]()
}

func CollectToArrayList[T any](it Iterator[T]) *ArrayList[T] {
	var r = ArrayListOf[T]()
	for v, ok := it.Next().Get(); ok; v, ok = it.Next().Get() {
		r.Append(v)
	}
	return r
}
