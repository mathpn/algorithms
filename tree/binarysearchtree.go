package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BinarySearchTree[K constraints.Ordered, V any] struct {
	root *searchNode[K, V]
}

func NewBinarySearchTree() *BinarySearchTree[int, string] {
	return &BinarySearchTree[int, string]{}
}

func (t *BinarySearchTree[K, V]) Insert(key K, value V) {
	if t.root == nil {
		t.root = &searchNode[K, V]{}
	}
	t.root.insert(key, value)
}

func (t *BinarySearchTree[K, V]) Search(key K) (V, error) {
	var v V
	if t.root == nil {
		return v, fmt.Errorf("tree is empty")
	}
	return t.root.search(key)
}

type searchNode[K constraints.Ordered, V any] struct {
	key   K
	value V
	left  *searchNode[K, V]
	right *searchNode[K, V]
}

func (n *searchNode[K, V]) insert(key K, value V) {
	if key < n.key {
		if n.left == nil {
			n.left = &searchNode[K, V]{key: key, value: value}
		} else {
			n.left.insert(key, value)
		}
	} else {
		if n.right == nil {
			n.right = &searchNode[K, V]{key: key, value: value}
		} else {
			n.right.insert(key, value)
		}
	}
}

func (n *searchNode[K, V]) search(key K) (V, error) {
	if n.key == key {
		return n.value, nil
	}
	if n.left != nil && key < n.key {
		return n.left.search(key)
	}
	if n.right != nil {
		return n.right.search(key)
	}
	var v V
	return v, fmt.Errorf("key %v not found", key)
}
