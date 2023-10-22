package main

import "fmt"

type node struct {
	next  *node
	value interface{}
}

type Queue struct {
	head *node
	tail *node
	Len  int
}

func NewQueue() *Queue {
	return &Queue{Len: 0}
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.head == nil {
		return nil, fmt.Errorf("queue is empty")
	}
	q.Len--
	h := q.head
	q.head = q.head.next
	h.next = nil

	if q.Len == 0 {
		q.tail = nil
	}

	return h.value, nil
}

func (q *Queue) Enqueue(v interface{}) {
	node := &node{value: v}
	q.Len++
	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue) Peek() (interface{}, error) {
	if q.head == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.head.value, nil
}
