package main

import "fmt"

type Stack struct {
	head *node
	Len  int
}

func NewStack() *Stack {
	return &Stack{Len: 0}
}

func (s *Stack) Push(v interface{}) {
	node := &node{value: v}
	s.Len++
	if s.head == nil {
		s.head = node
		return
	}
	node.next = s.head
	s.head = node
}

func (s *Stack) Pop() (interface{}, error) {
	if s.head == nil {
		return nil, fmt.Errorf("stack is empty")
	}
	s.Len--
	h := s.head
	s.head = s.head.next
	h.next = nil
	return h.value, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.head == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	return s.head.value, nil
}
