package list

import "fmt"

type Stack[T any] struct {
	head *node[T]
	Len  int
}

func NewStack() *Stack[int] {
	return &Stack[int]{Len: 0}
}

func (s *Stack[T]) Push(v T) {
	node := &node[T]{value: v}
	s.Len++
	if s.head == nil {
		s.head = node
		return
	}
	node.next = s.head
	s.head = node
}

func (s *Stack[T]) Pop() (T, error) {
	var t T
	if s.head == nil {
		return t, fmt.Errorf("stack is empty")
	}
	s.Len--
	h := s.head
	s.head = s.head.next
	h.next = nil
	return h.value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var t T
	if s.head == nil {
		return t, fmt.Errorf("stack is empty")
	}
	return s.head.value, nil
}
