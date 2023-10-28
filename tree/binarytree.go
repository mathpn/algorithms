package main

import (
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
