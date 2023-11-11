package tree

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

func (b *BinaryTree[T]) DepthFirstSearch(value T) bool {
	return search(b.root, value, false)
}

func search[T constraints.Ordered](curr *bnode[T], value T, found bool) bool {
	if curr == nil {
		return found
	}
	if curr.value == value {
		return true
	}

	return search(curr.left, value, found) || search(curr.right, value, found)
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

func (t *BinaryTree[T]) BreadthFirstSearch(value T) bool {
	if t.root == nil {
		return false
	}
	curr := t.root
	q := list.Queue[*bnode[T]]{}
	q.Enqueue(curr)
	for q.Len > 0 {
		curr, err := q.Dequeue()
		if curr.value == value {
			return true
		}
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
	return false
}

func CompareTrees[T constraints.Ordered](a *BinaryTree[T], b *BinaryTree[T]) bool {
	return compare(a.root, b.root)
}

func compare[T constraints.Ordered](a *bnode[T], b *bnode[T]) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.value != b.value {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
}
