package main

import (
	"main/list"

	"golang.org/x/exp/constraints"
)

type bnode[T constraints.Ordered] struct {
	value T
	right *bnode[T]
	left  *bnode[T]
}

type BinaryTree[T constraints.Ordered] struct {
	root *bnode[T]
}

func NewBinaryTree() *BinaryTree[int] {
	return &BinaryTree[int]{}
}

func (b *BinaryTree[T]) PreOrderTraversal() []T {
	path := []T{}
	processNode := func(node *bnode[T]) {
		path = append(path, node.value)
	}
	walkPre(b.root, processNode)
	return path
}

func (b *BinaryTree[T]) InOrderTraversal() []T {
	path := []T{}
	processNode := func(node *bnode[T]) {
		path = append(path, node.value)
	}
	walkIn(b.root, processNode)
	return path
}

func (b *BinaryTree[T]) PostOrderTraversal() []T {
	path := []T{}
	processNode := func(node *bnode[T]) {
		path = append(path, node.value)
	}
	walkPost(b.root, processNode)
	return path
}

func walkPre[T constraints.Ordered](curr *bnode[T], processNode func(*bnode[T])) {
	if curr == nil {
		return
	}
	processNode(curr)
	walkPre(curr.left, processNode)
	walkPre(curr.right, processNode)
}

func walkIn[T constraints.Ordered](curr *bnode[T], processNode func(*bnode[T])) {
	if curr == nil {
		return
	}
	walkIn(curr.left, processNode)
	processNode(curr)
	walkIn(curr.right, processNode)
}

func walkPost[T constraints.Ordered](curr *bnode[T], processNode func(*bnode[T])) {
	if curr == nil {
		return
	}
	walkPost(curr.left, processNode)
	walkPost(curr.right, processNode)
	processNode(curr)
}

func (t *BinaryTree[T]) BreadthFirstTraversal() []T {
	out := make([]T, 0)
	if t.root == nil {
		return out
	}
	curr := t.root
	q := list.Queue[*bnode[T]]{}
	q.Enqueue(curr)
	for q.Len > 0 {
		curr, err := q.Dequeue()
		out = append(out, curr.value)
		if err != nil {
			panic(err) // TODO improve
		}
		if curr.left != nil {
			q.Enqueue(curr.left)
		}
		if curr.right != nil {
			q.Enqueue(curr.right)
		}
	}
	return out
}
