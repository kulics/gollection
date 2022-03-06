package main

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func GetNode[T any](n *Node[T], i int) T {
	if i == 0 {
		return n.Value
	} else {
		return GetNode(n.Next, i-1)
	}
}

func SetNode[T any](n *Node[T], i int, v T) T {
	if i == 0 {
		var oldValue = n.Value
		n.Value = v
		return oldValue
	} else {
		return SetNode(n.Next, i-1, v)
	}
}

func PrependNode[T any](n *Node[T], v T) {
	if n == nil {
		*n = Node[T]{v, nil}
	} else {
		*n = Node[T]{v, n}
	}
}

func AppendNode[T any](n *Node[T], v T) {
	if n == nil {
		*n = Node[T]{v, nil}
	} else {
		AppendNode(n.Next, v)
	}
}

func InsertNode[T any](n *Node[T], pre *Node[T], i int, e T) {
	if i == 0 {
		pre.Next = &Node[T]{e, n}
	} else {
		InsertNode(n.Next, n, i-1, e)
	}
}

func RemoveNode[T any](n *Node[T], pre *Node[T], i int) T {
	if i == 0 {
		var item = n.Value
		pre.Next = n.Next
		n.Next = nil
		return item
	} else {
		return RemoveNode(n.Next, n, i-1)
	}
}
