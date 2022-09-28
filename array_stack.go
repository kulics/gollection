package gollection

// Constructing an ArrayStack with variable-length parameters
func ArrayStackOf[T any](elements ...T) ArrayStack[T] {
	var length = len(elements)
	var stack = MakeArrayStack[T](length)
	copy(stack.inner.elements, elements)
	stack.inner.length = length
	return stack
}

// Constructing an empty ArrayStack with capacity.
func MakeArrayStack[T any](capacity int) ArrayStack[T] {
	if capacity < defaultElementsLength {
		capacity = defaultElementsLength
	}
	var inner = &arrayStack[T]{make([]T, capacity), 0}
	return ArrayStack[T]{inner}
}

// Constructing an ArrayStack from other Collection.
func ArrayStackFrom[T any](collection Collection[T]) ArrayStack[T] {
	var inner = &arrayStack[T]{collection.ToSlice(), collection.Count()}
	return ArrayStack[T]{inner}
}

// Stack implemented using Array.
type ArrayStack[T any] struct {
	inner *arrayStack[T]
}

type arrayStack[T any] struct {
	elements []T
	length   int
}

// Return the number of elements of stack.
func (a ArrayStack[T]) Length() int {
	return a.inner.length
}

// Return true when the number of elements of stack is 0.
func (a ArrayStack[T]) IsEmpty() bool {
	return a.inner.length == 0
}

// Add an element to the top of the stack.
func (a ArrayStack[T]) Push(element T) {
	if growLength := a.inner.length + 1; len(a.inner.elements) < growLength {
		a.grow(growLength)
	}
	a.inner.elements[a.inner.length] = element
	a.inner.length++
}

// Remove an element from the top of the stack.
// A panic is raised when the stack is empty.
func (a ArrayStack[T]) Pop() T {
	if v, ok := a.TryPop().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

// Return an element at the top of the stack, but does not remove it.
// A panic is raised when the stack is empty.
func (a ArrayStack[T]) Peek() T {
	if v, ok := a.TryPeek().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

// Remove an element from the top of the stack.
// Return None when the stack is empty.
func (a ArrayStack[T]) TryPop() Option[T] {
	if a.IsEmpty() {
		return None[T]()
	}
	var index = a.inner.length - 1
	var item = a.inner.elements[index]
	var empty T
	a.inner.elements[index] = empty
	a.inner.length--
	return Some(item)
}

// Return an element at the top of the stack, but does not remove it.
// Return None when the stack is empty.
func (a ArrayStack[T]) TryPeek() Option[T] {
	if a.IsEmpty() {
		return None[T]()
	}
	return Some(a.inner.elements[a.inner.length-1])
}

// Return the Iterator of stack.
func (a ArrayStack[T]) Iter() Iterator[T] {
	return &arrayStackIterator[T]{a.Length(), a}
}

// Return a new built-in slice that copies all elements.
func (a ArrayStack[T]) ToSlice() []T {
	var arr = make([]T, a.Length())
	copy(arr, a.inner.elements)
	return arr
}

// Return a new stack that copies all elements.
func (a ArrayStack[T]) Clone() ArrayStack[T] {
	var elements = make([]T, len(a.inner.elements))
	copy(elements, a.inner.elements)
	var inner = &arrayStack[T]{
		elements: elements,
		length:   a.inner.length,
	}
	return ArrayStack[T]{inner}
}

// Ensure that stack have enough space before expansion.
func (a ArrayStack[T]) Reserve(additional int) {
	if addable := len(a.inner.elements) - a.inner.length; addable < additional {
		a.grow(a.inner.length + additional)
	}
}

// Return the capacity of stack.
func (a ArrayStack[T]) Capacity() int {
	return len(a.inner.elements)
}

// Clears all elements, but does not reset the space.
func (a ArrayStack[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.inner.length; i++ {
		a.inner.elements[i] = emptyValue
	}
	a.inner.length = 0
}

func (a ArrayStack[T]) grow(minCapacity int) {
	var newLength = arrayGrow(len(a.inner.elements))
	if newLength < minCapacity {
		newLength = minCapacity
	}
	var newSource = make([]T, newLength)
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
