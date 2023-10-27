package main

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("enqueue/deque 1", func(t *testing.T) {
		q := NewQueue()
		v := 42
		q.Enqueue(v)
		o, err := q.Dequeue()
		if err != nil {
			t.Fatal(err)
		}
		if o != v {
			t.Errorf("expected %d, got %d", v, o)
		}
	})

	t.Run("enqueue/dequeue 3", func(t *testing.T) {
		q := NewQueue()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			q.Enqueue(v)
		}
		for _, v := range values {
			o, err := q.Dequeue()
			if err != nil {
				t.Fatal(err)
			}
			if o != v {
				t.Errorf("expected %d, got %d", v, o)
			}
		}
	})

	t.Run("enqueue and peek", func(t *testing.T) {
		q := NewQueue()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			q.Enqueue(v)
			o, err := q.Peek()
			if err != nil {
				t.Fatal(err)
			}
			if values[0] != o {
				t.Errorf("expected %d, got %d", values[0], o)
			}
		}
	})

	t.Run("deque and peek", func(t *testing.T) {
		q := NewQueue()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			q.Enqueue(v)
		}
		for _, v := range values {
			o, err := q.Peek()
			if err != nil {
				t.Fatal(err)
			}
			if v != o {
				t.Errorf("expected %d, got %d", v, o)
			}
			q.Dequeue()
		}
	})

	t.Run("enqueue and dequeue 3", func(t *testing.T) {
		q := NewQueue()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			q.Enqueue(v)
			o, err := q.Dequeue()
			if err != nil {
				t.Fatal(err)
			}
			if o != v {
				t.Errorf("expected %d, got %d", v, o)
			}
		}
	})

	t.Run("enqueue check len", func(t *testing.T) {
		q := NewQueue()
		values := [3]int{1, 2, 3}
		for i, v := range values {
			q.Enqueue(v)
			if i+1 != q.Len {
				t.Errorf("expected %d, got %d", i+1, q.Len)
			}
		}
	})

	t.Run("deque check len", func(t *testing.T) {
		q := NewQueue()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			q.Enqueue(v)
		}

		for i := len(values); i > 0; i-- {
			if i != q.Len {
				t.Errorf("expected %d, got %d", i, q.Len)
			}
			q.Dequeue()
		}
	})
}

func TestStack(t *testing.T) {
	t.Run("push/pop 1", func(t *testing.T) {
		q := NewStack()
		v := 42
		q.Push(v)
		o, err := q.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if o != v {
			t.Errorf("expected %d, got %d", v, o)
		}
	})

	t.Run("push/pop 3", func(t *testing.T) {
		s := NewStack()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			s.Push(v)
		}
		for i := len(values) - 1; i >= 0; i-- {
			o, err := s.Pop()
			if err != nil {
				t.Fatal(err)
			}
			if o != values[i] {
				t.Errorf("expected %d, got %d", values[i], o)
			}
		}
	})

	t.Run("push and peek", func(t *testing.T) {
		s := NewStack()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			s.Push(v)
			o, err := s.Peek()
			if err != nil {
				t.Fatal(err)
			}
			if v != o {
				t.Errorf("expected %d, got %d", v, o)
			}
		}
	})

	t.Run("pop and peek", func(t *testing.T) {
		s := NewStack()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			s.Push(v)
		}
		for i := len(values) - 1; i > 0; i-- {
			o, err := s.Peek()
			if err != nil {
				t.Fatal(err)
			}
			if values[i] != o {
				t.Errorf("expected %d, got %d", values[i], o)
			}
			s.Pop()
		}
	})

	t.Run("push and pop 3", func(t *testing.T) {
		s := NewStack()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			s.Push(v)
			o, err := s.Pop()
			if err != nil {
				t.Fatal(err)
			}
			if o != v {
				t.Errorf("expected %d, got %d", v, o)
			}
		}
	})

	t.Run("push check len", func(t *testing.T) {
		s := NewStack()
		values := [3]int{1, 2, 3}
		for i, v := range values {
			s.Push(v)
			if i+1 != s.Len {
				t.Errorf("expected %d, got %d", i+1, s.Len)
			}
		}
	})

	t.Run("pop check len", func(t *testing.T) {
		s := NewStack()
		values := [3]int{1, 2, 3}
		for _, v := range values {
			s.Push(v)
		}

		for i := len(values); i > 0; i-- {
			if i != s.Len {
				t.Errorf("expected %d, got %d", i, s.Len)
			}
			s.Pop()
		}
	})
}

func TestDoublyLinkedList(t *testing.T) {
	type dblTest struct {
		values []int
	}
	tests := []dblTest{
		{values: []int{1, 2, 3}},
		{values: []int{3, 2, 1}},
		{values: []int{-10, 20, 30, 40}},
	}
	t.Run("append and get", func(t *testing.T) {
		l := NewDoublyLinkedList()
		for _, values := range tests {
			for _, v := range values.values {
				l.Append(v)
				o, err := l.Get(l.Len - 1)
				if err != nil {
					t.Fatal(err)
				}
				if v != o {
					t.Errorf("expected %d, got %d", v, o)
				}
			}
		}
	})
	t.Run("prepend and get", func(t *testing.T) {
		l := NewDoublyLinkedList()
		for _, values := range tests {
			for _, v := range values.values {
				l.Prepend(v)
				o, err := l.Get(0)
				if err != nil {
					t.Fatal(err)
				}
				if v != o {
					t.Errorf("expected %d, got %d", v, o)
				}
			}
		}
	})
	t.Run("insertAt and get", func(t *testing.T) {
		l := NewDoublyLinkedList()
		l.Append(-1)
		for _, values := range tests {
			for _, v := range values.values {
				l.InsertAt(v, 1)
				o, err := l.Get(1)
				if err != nil {
					t.Fatal(err)
				}
				if v != o {
					t.Errorf("expected %d, got %d", v, o)
				}
			}
		}
	})
	t.Run("remove", func(t *testing.T) {
		for _, values := range tests {
			l := NewDoublyLinkedList()
			for _, v := range values.values {
				l.Append(v)
			}
			err := l.Remove(values.values[0])
			if err != nil {
				t.Error(err)
			}
			err = l.Remove(values.values[0])
			if err == nil {
				t.Errorf("expected error, got nil when removing %d", values.values[0])
			}
		}
	})
	t.Run("removeAt", func(t *testing.T) {
		for _, values := range tests {
			idx := []int{0, len(values.values) - 1, len(values.values) / 2}
			l := NewDoublyLinkedList()
			for _, v := range values.values {
				l.Append(v)
			}
			for _, idx := range idx {
				v := values.values[idx]
				o, err := l.RemoveAt(idx)
				if err != nil {
					t.Error(err)
				}
				if o != v {
					t.Errorf("expected %d, got %d", v, o)
				}
			}
		}
	})
}

func BenchmarkQueue(b *testing.B) {
	lens := [4]int{1000, 10000, 100000, 1000000}
	for _, l := range lens {
		q := NewQueue()
		b.Run(fmt.Sprintf("enqueue %d", l), func(_ *testing.B) {
			for i := 0; i < l; i++ {
				q.Enqueue(i)
			}
		})
	}
	for _, l := range lens {
		q := NewQueue()
		b.Run(fmt.Sprintf("deque %d", l), func(b *testing.B) {
			for i := 0; i < l; i++ {
				q.Enqueue(i)
			}
			b.StartTimer()
			for i := 0; i < l; i++ {
				q.Dequeue()
			}
			b.StopTimer()
		})
	}
	for _, l := range lens {
		q := NewQueue()
		b.Run(fmt.Sprintf("enqueue deque %d", l), func(_ *testing.B) {
			for i := 0; i < l; i++ {
				q.Enqueue(i)
				q.Dequeue()
			}
		})
	}
}

func BenchmarkStack(b *testing.B) {
	lens := [4]int{1000, 10000, 100000, 1000000}
	for _, l := range lens {
		s := NewStack()
		b.Run(fmt.Sprintf("push %d", l), func(_ *testing.B) {
			for i := 0; i < l; i++ {
				s.Push(i)
			}
		})
	}
	for _, l := range lens {
		s := NewStack()
		b.Run(fmt.Sprintf("pop %d", l), func(b *testing.B) {
			for i := 0; i < l; i++ {
				s.Push(i)
			}
			b.StartTimer()
			for i := 0; i < l; i++ {
				s.Pop()
			}
			b.StopTimer()
		})
	}
	for _, l := range lens {
		s := NewStack()
		b.Run(fmt.Sprintf("push pop %d", l), func(_ *testing.B) {
			for i := 0; i < l; i++ {
				s.Push(i)
				s.Pop()
			}
		})
	}
}
