package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type sortParams struct {
	input    []int
	value    int
	expected int
}

func createParams() []sortParams {
	return []sortParams{
		{input: []int{1, 2, 3}, value: 1, expected: 0},
		{input: []int{1, 2, 3, 4}, value: 2, expected: 1},
		{input: []int{1, 2, 3}, value: 4, expected: -1},
		{input: []int{10, 20, 30}, value: 10, expected: 0},
		{input: []int{10, 20, 30, 40}, value: 40, expected: 3},
	}
}

func TestLinearSearch(t *testing.T) {
	params := createParams()
	for _, param := range params {
		idx, err := LinearSearch(param.input, param.value)
		if param.expected >= 0 && err != nil {
			t.Fatal(err)
		}
		if param.expected == -1 && err == nil {
			t.Errorf("binary search failed: %v (value: %d) - expected error, got %d", param.input, param.value, idx)
		}
		if param.expected >= 0 && idx != param.expected {
			t.Fatalf("binary search failed: %v (value: %d) - expected %d, got %d", param.input, param.value, param.expected, idx)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	params := createParams()
	for _, param := range params {
		idx, err := BinarySearch(param.input, param.value)
		if param.expected >= 0 && err != nil {
			t.Fatal(err)
		}
		if param.expected == -1 && err == nil {
			t.Errorf("binary search failed: %v (value: %d) - expected error, got %d", param.input, param.value, idx)
		}
		if param.expected >= 0 && idx != param.expected {
			t.Fatalf("binary search failed: %v (value: %d) - expected %d, got %d", param.input, param.value, param.expected, idx)
		}
	}
}

func createSlices() [][]int {
	lens := [5]int{100000, 1000000, 10000000, 100000000, 1000000000}
	arrays := make([][]int, 0)
	for _, l := range lens {
		arr := make([]int, l)
		for i := 0; i < l; i++ {
			arr[i] = i
		}
		arrays = append(arrays, arr)
	}
	return arrays
}

func BenchmarkLinearSearch(b *testing.B) {
	for _, slice := range createSlices() {
		b.Run(fmt.Sprintf("linear search %d", len(slice)), func(_ *testing.B) {
			v := rand.Intn(len(slice))
			b.StartTimer()
			_, err := LinearSearch(slice, v)
			if err != nil {
				panic(err)
			}
			b.StopTimer()
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for _, slice := range createSlices() {
		b.Run(fmt.Sprintf("binary search %d", len(slice)), func(_ *testing.B) {
			v := rand.Intn(len(slice))
			b.StartTimer()
			_, err := BinarySearch(slice, v)
			if err != nil {
				panic(err)
			}
			b.StopTimer()
		})
	}
}
