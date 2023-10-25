package main

import "fmt"

type llnode struct {
	value interface{}
	next  *llnode
	prev  *llnode
}

type DoublyLinkedList struct {
	head *llnode
	tail *llnode
	Len  int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{Len: 0}
}

func (l *DoublyLinkedList) Prepend(value interface{}) {
	node := &llnode{value: value}
	if l.Len == 0 {
		l.Append(value)
		return
	}
	l.Len++
	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *DoublyLinkedList) InsertAt(value interface{}, idx int) error {
	if idx > l.Len {
		return fmt.Errorf("idx %d out of bounds of list with length %d", idx, l.Len)
	} else if idx == l.Len {
		l.Append(value)
	} else if idx == 0 {
		l.Prepend(value)
	}
	l.Len++
	node := &llnode{value: value}
	curr := l.getAt(idx)
	node.next = curr
	node.prev = curr.prev
	curr.prev = node
	node.prev.next = node
	return nil
}

func (l *DoublyLinkedList) Append(value interface{}) {
	node := &llnode{value: value}
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

func (l *DoublyLinkedList) Remove(value interface{}) error {
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

func (l *DoublyLinkedList) Get(idx int) (interface{}, error) {
	if idx > (l.Len - 1) {
		return nil, fmt.Errorf("idx %d out of bounds of list with length %d", idx, l.Len)
	}
	n := l.getAt(idx)
	return n.value, nil
}

func (l *DoublyLinkedList) RemoveAt(idx int) (interface{}, error) {
	node := l.getAt(idx)
	if node == nil {
		return nil, fmt.Errorf("idx %d is out of bounds of list with length %d", idx, l.Len)
	}
	l.remove(node)
	return node.value, nil
}

func (l *DoublyLinkedList) remove(node *llnode) {
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

func (l *DoublyLinkedList) getAt(idx int) *llnode {
	curr := l.head
	for i := 0; i < idx; i++ {
		if curr.next == nil {
			break
		}
		curr = curr.next
	}
	return curr
}
