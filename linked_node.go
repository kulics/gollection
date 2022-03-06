package main

type Node[T any] struct {
	Value T
	Next *Node[T]
}