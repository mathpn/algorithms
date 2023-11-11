package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	data []T
	Len  int
}

func NewMinHeap() *MinHeap[int] {
	return &MinHeap[int]{data: make([]int, 0), Len: 0}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return 2*idx + 1
}

func rightChild(idx int) int {
	return 2*idx + 2
}

func (h *MinHeap[T]) Insert(value T) error {
	h.data = append(h.data, value)
	h.heapifyUp(h.Len)
	h.Len++
	return nil
}

func (h *MinHeap[T]) Delete() (T, error) {
	var t T
	if h.Len == 0 {
		return t, fmt.Errorf("heap is empty")
	}
	t = h.data[0]
	h.Len--
	if h.Len == 0 {
		h.data = make([]T, 0)
		return t, nil
	}

	h.data[0] = h.data[h.Len]
	h.data = h.data[:h.Len+1]
	h.heapifyDown(0)
	return t, nil
}

func (h *MinHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	p := parent(idx)
	pv := h.data[p]
	v := h.data[idx]
	if pv > v {
		h.data[idx] = pv
		h.data[p] = v
		h.heapifyUp(parent(idx))
	}
}

func (h *MinHeap[T]) heapifyDown(idx int) {
	if idx >= h.Len {
		return
	}

	li := leftChild(idx)
	if li >= h.Len {
		return
	}
	ri := rightChild(idx)

	lv := h.data[li]
	rv := h.data[ri]
	v := h.data[idx]

	if lv > rv && v > rv {
		h.data[idx] = rv
		h.data[ri] = v
		h.heapifyDown(ri)
	} else if rv > lv && v > lv {
		h.data[idx] = lv
		h.data[li] = v
		h.heapifyDown(li)
	}
}
