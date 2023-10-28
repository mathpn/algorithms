package main

import "fmt"

type llnode[T comparable] struct {
	value T
	next  *llnode[T]
	prev  *llnode[T]
}

type DoublyLinkedList[T comparable] struct {
	head *llnode[T]
	tail *llnode[T]
	Len  int
}

func NewDoublyLinkedList() *DoublyLinkedList[int] {
	return &DoublyLinkedList[int]{Len: 0}
}

func (l *DoublyLinkedList[T]) Prepend(value T) {
	node := &llnode[T]{value: value}
	if l.Len == 0 {
		l.Append(value)
		return
	}
	l.Len++
	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *DoublyLinkedList[T]) InsertAt(value T, idx int) error {
	if idx > l.Len {
		return fmt.Errorf("idx %d out of bounds of list with length %d", idx, l.Len)
	} else if idx == l.Len {
		l.Append(value)
		return nil
	} else if idx == 0 {
		l.Prepend(value)
		return nil
	}
	l.Len++
	node := &llnode[T]{value: value}
	curr := l.getAt(idx)
	node.next = curr
	node.prev = curr.prev
	curr.prev = node
	node.prev.next = node
	return nil
}

func (l *DoublyLinkedList[T]) Append(value T) {
	node := &llnode[T]{value: value}
	l.Len++
	if l.tail == nil {
		l.head = node
		l.tail = node
		return
	}
	node.prev = l.tail
	l.tail.next = node
	l.tail = node
}

func (l *DoublyLinkedList[T]) Remove(value T) error {
	curr := l.head
	for i := 0; i < l.Len; i++ {
		if curr.value == value {
			break
		}
		curr = curr.next
	}
	if curr == nil || curr.value != value {
		return fmt.Errorf("value not found in list")
	}
	l.remove(curr)
	return nil
}

func (l *DoublyLinkedList[T]) Get(idx int) (T, error) {
	var t T
	if idx > (l.Len - 1) {
		return t, fmt.Errorf("idx %d out of bounds of list with length %d", idx, l.Len)
	}
	n := l.getAt(idx)
	return n.value, nil
}

func (l *DoublyLinkedList[T]) RemoveAt(idx int) (T, error) {
	var t T
	node := l.getAt(idx)
	if node == nil {
		return t, fmt.Errorf("idx %d is out of bounds of list with length %d", idx, l.Len)
	}
	l.remove(node)
	return node.value, nil
}

func (l *DoublyLinkedList[T]) remove(node *llnode[T]) {
	l.Len--
	if l.Len == 0 {
		l.head = nil
		l.tail = nil
		return
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node == l.head {
		l.head = node.next
	}
	if node == l.tail {
		l.tail = node.prev
	}
	node.prev = nil
	node.next = nil
}

func (l *DoublyLinkedList[T]) getAt(idx int) *llnode[T] {
	curr := l.head
	for i := 0; i < idx; i++ {
		if curr.next == nil {
			break
		}
		curr = curr.next
	}
	return curr
}
