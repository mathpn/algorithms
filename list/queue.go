package list

import "fmt"

type node[T any] struct {
	next  *node[T]
	value T
}

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	Len  int
}

func NewQueue() *Queue[int] {
	return &Queue[int]{Len: 0}
}

func (q *Queue[T]) Dequeue() (T, error) {
	var t T
	if q.head == nil {
		return t, fmt.Errorf("queue is empty")
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

func (q *Queue[T]) Enqueue(v T) {
	node := &node[T]{value: v}
	q.Len++
	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Peek() (T, error) {
	var t T
	if q.head == nil {
		return t, fmt.Errorf("queue is empty")
	}
	return q.head.value, nil
}
