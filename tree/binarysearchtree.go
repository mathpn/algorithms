package tree

import (
	"fmt"
	"main/list"
	"strings"

	"golang.org/x/exp/constraints"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type BinarySearchable[K constraints.Ordered, V any] interface {
	Insert(key K, value V)
	Search(key K) (V, error)
	Dump() (string, error)
}

type BinarySearchTree[K constraints.Ordered, V any] struct {
	root *simpleNode[K, V]
}

func NewBinarySearchTree[K constraints.Ordered, V any]() BinarySearchable[K, V] {
	return &BinarySearchTree[K, V]{}
}

func (t *BinarySearchTree[K, V]) Insert(key K, value V) {
	if t.root == nil {
		t.root = &simpleNode[K, V]{}
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

func (t *BinarySearchTree[K, V]) Dump() (string, error) {
	pq := &list.Queue[string]{}
	t.root.dump(pq, 0, "")
	out := ""
	var line string
	var err error
	for pq.Len > 0 {
		line, err = pq.Dequeue()
		if err != nil {
			return "", err
		}
		out += line
	}
	return out, nil
}

type simpleNode[K constraints.Ordered, V any] struct {
	key         K
	value       V
	left, right *simpleNode[K, V]
}

func (n *simpleNode[K, V]) insert(key K, value V) {
	if key == n.key {
		n.value = value
		return
	}

	if key < n.key {
		if n.left == nil {
			n.left = &simpleNode[K, V]{key: key, value: value}
		} else {
			n.left.insert(key, value)
		}
	} else {
		if n.right == nil {
			n.right = &simpleNode[K, V]{key: key, value: value}
		} else {
			n.right.insert(key, value)
		}
	}
}

func (n *simpleNode[K, V]) search(key K) (V, error) {
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

func (n *simpleNode[K, V]) dump(pq *list.Queue[string], i int, lr string) {
	if n == nil {
		return
	}
	indent := ""
	if i > 0 {
		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}
	pq.Enqueue(fmt.Sprintf("%s{key: %v, value: %v}\n", indent, n.key, n.value))
	n.left.dump(pq, i+1, "L")
	n.right.dump(pq, i+1, "R")
}

type AVLTree[K constraints.Ordered, V any] struct {
	root *avlNode[K, V]
}

func NewAVLTree[K constraints.Ordered, V any]() BinarySearchable[K, V] {
	return &AVLTree[K, V]{}
}

func (t *AVLTree[K, V]) Insert(key K, value V) {
	t.root = t.root.insert(key, value)
	if t.root != nil && (t.root.bal() < -1 || t.root.bal() > 1) {
		t.root.rebalance()
	}
}

func (t *AVLTree[K, V]) Search(key K) (V, error) {
	var v V
	if t.root == nil {
		return v, fmt.Errorf("tree is empty")
	}
	return t.root.search(key)
}

func (t *AVLTree[K, V]) Dump() (string, error) {
	pq := &list.Queue[string]{}
	t.root.dump(pq, 0, "")
	out := ""
	var line string
	var err error
	for pq.Len > 0 {
		line, err = pq.Dequeue()
		if err != nil {
			return "", err
		}
		out += line
	}
	return out, nil
}

type avlNode[K constraints.Ordered, V any] struct {
	left   *avlNode[K, V]
	right  *avlNode[K, V]
	key    K
	value  V
	height int
}

func (n *avlNode[K, V]) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *avlNode[K, V]) bal() int {
	return n.right.Height() - n.left.Height()
}

func (n *avlNode[K, V]) insert(key K, value V) *avlNode[K, V] {
	if n == nil {
		return &avlNode[K, V]{key: key, value: value}
	}
	if key == n.key {
		n.value = value
		return n
	}

	if key < n.key {
		n.left = n.left.insert(key, value)
	} else {
		n.right = n.right.insert(key, value)
	}
	n.height = max(n.left.Height(), n.right.Height()) + 1
	return n.rebalance()
}

func (n *avlNode[K, V]) search(key K) (V, error) {
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

func (n *avlNode[K, V]) updateHeight() {
	n.height = max(n.left.Height(), n.right.Height()) + 1
}

func (n *avlNode[K, V]) rotateLeft() *avlNode[K, V] {
	r := n.right
	n.right = r.left
	r.left = n

	n.updateHeight()
	r.updateHeight()
	return r
}

func (n *avlNode[K, V]) rotateRight() *avlNode[K, V] {
	l := n.left
	n.left = l.right
	l.right = n

	n.updateHeight()
	l.updateHeight()
	return l
}

func (n *avlNode[K, V]) rotateRightLeft() *avlNode[K, V] {
	n.right = n.right.rotateRight()
	n = n.rotateLeft()

	n.updateHeight()
	return n
}

func (n *avlNode[K, V]) rotateLeftRight() *avlNode[K, V] {
	n.left = n.left.rotateLeft()
	n = n.rotateRight()

	n.updateHeight()
	return n
}

func (n *avlNode[K, V]) rebalance() *avlNode[K, V] {
	switch {
	case n.bal() < -1 && n.left.bal() == -1:
		return n.rotateRight()
	case n.bal() > 1 && n.right.bal() == 1:
		return n.rotateLeft()
	case n.bal() < -1 && n.left.bal() == 1:
		return n.rotateLeftRight()
	case n.bal() > 1 && n.right.bal() == -1:
		return n.rotateRightLeft()
	}
	return n
}

func (n *avlNode[K, V]) dump(pq *list.Queue[string], i int, lr string) {
	if n == nil {
		return
	}
	indent := ""
	if i > 0 {
		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}
	pq.Enqueue(fmt.Sprintf("%s{key: %v, value: %v}[%d,%d]\n", indent, n.key, n.value, n.bal(), n.Height()))
	n.left.dump(pq, i+1, "L")
	n.right.dump(pq, i+1, "R")
}
