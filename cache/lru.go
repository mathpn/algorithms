package cache

import "fmt"

type lruNode[T any] struct {
	value T
	next  *lruNode[T]
	prev  *lruNode[T]
}

type LRU[K comparable, V any] struct {
	head      *lruNode[V]
	tail      *lruNode[V]
	lookup    map[K]*lruNode[V]
	revLookup map[*lruNode[V]]K
	len       int
	capacity  int
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		len:       0,
		lookup:    make(map[K]*lruNode[V], capacity),
		capacity:  capacity,
		revLookup: make(map[*lruNode[V]]K, capacity),
	}
}

func (l *LRU[K, V]) Update(key K, value V) {
	node, ok := l.lookup[key]
	if !ok {
		node = &lruNode[V]{value: value}
		l.len++
		l.prepend(node)
		l.evict()

		l.lookup[key] = node
		l.revLookup[node] = key
		return
	}

	l.detach(node)
	l.prepend(node)
	node.value = value
}

func (l *LRU[K, V]) Get(key K) (V, error) {
	var value V
	node, ok := l.lookup[key]
	if !ok {
		return value, fmt.Errorf("key %v not found", key)
	}

	l.detach(node)
	l.prepend(node)

	return node.value, nil
}

func (l *LRU[K, V]) detach(node *lruNode[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = node.next
	}
	if l.tail == node {
		l.tail = node.prev
	}

	node.next = nil
	node.prev = nil
}

func (l *LRU[K, V]) prepend(node *lruNode[V]) {
	if l.head == nil {
		l.head = node
		l.tail = l.head
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *LRU[K, V]) evict() {
	if l.len <= l.capacity {
		return
	}
	tail := l.tail
	l.detach(tail)

	key := l.revLookup[tail]
	delete(l.lookup, key)
	delete(l.revLookup, tail)
	l.len--
}
