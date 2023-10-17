package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

type sortParams struct {
	input    []int
	expected []int
}

func createParams() []sortParams {
	r := rand.Perm(100)
	sorted := make([]int, 100)
	copy(sorted, r)
	sort.Ints(sorted)
	return []sortParams{
		{input: []int{3, 4, 1, 2}, expected: []int{1, 2, 3, 4}},
		{input: []int{10, 30, 20, 40}, expected: []int{10, 20, 30, 40}},
		{input: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
		{input: []int{4, 3, 2, 1}, expected: []int{1, 2, 3, 4}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: r, expected: sorted},
	}
}

func TestBubbleSort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		output := BubbleSort(test.input)
		if !slices.IsSorted(output) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, output)
		}
	}
}

func TestCocktailSort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		output := CocktailSort(test.input)
		if !slices.IsSorted(output) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, output)
		}
	}
}

func TestGnomeSort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		output := GnomeSort(test.input)
		if !slices.IsSorted(output) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, output)
		}
	}
}

func TestOddEvenSort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		output := OddEvenSort(test.input)
		if !slices.IsSorted(output) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, output)
		}
	}
}

func createSlices() [][]int {
	return [][]int{
		rand.Perm(10),
		rand.Perm(100),
		rand.Perm(1000),
		rand.Perm(10000),
		rand.Perm(100000),
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for _, slice := range createSlices() {
		b.Run(fmt.Sprintf("bubble sort %d", len(slice)), func(_ *testing.B) {
			BubbleSort(slice)
		})
	}
}

func BenchmarkCocktailSort(b *testing.B) {
	for _, slice := range createSlices() {
		b.Run(fmt.Sprintf("cocktail sort %d", len(slice)), func(_ *testing.B) {
			CocktailSort(slice)
		})
	}
}

func BenchmarkGnomeSort(b *testing.B) {
	for _, slice := range createSlices() {
		b.Run(fmt.Sprintf("gnome sort %d", len(slice)), func(_ *testing.B) {
			GnomeSort(slice)
		})
	}
}

func BenchmarkOddEvenSort(b *testing.B) {
	for _, slice := range createSlices() {
		b.Run(fmt.Sprintf("odd even sort %d", len(slice)), func(_ *testing.B) {
			OddEvenSort(slice)
		})
	}
}
