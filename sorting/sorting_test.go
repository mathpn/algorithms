package sorting

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
		BubbleSort(test.input)
		if !slices.IsSorted(test.input) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, test.input)
		}
	}
}

func TestCocktailSort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		CocktailSort(test.input)
		if !slices.IsSorted(test.input) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, test.input)
		}
	}
}

func TestGnomeSort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		GnomeSort(test.input)
		if !slices.IsSorted(test.input) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, test.input)
		}
	}
}

func TestOddEvenSort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		OddEvenSort(test.input)
		if !slices.IsSorted(test.input) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, test.input)
		}
	}
}

func TestQuicksort(t *testing.T) {
	tests := createParams()
	for _, test := range tests {
		QuickSort(test.input)
		if !slices.IsSorted(test.input) {
			t.Errorf("array is not sorted. Expected %v, got %v", test.expected, test.input)
		}
	}
}

func createRandomSlices() [][]int {
	return [][]int{
		rand.Perm(10),
		rand.Perm(100),
		rand.Perm(1000),
		rand.Perm(10000),
		rand.Perm(100000),
	}
}

func createSortedSlices() [][]int {
	lens := [5]int{10, 100, 1000, 10000, 100000}
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

func createRevSortedSlices() [][]int {
	lens := [5]int{10, 100, 1000, 10000, 100000}
	arrays := make([][]int, 0)
	for _, l := range lens {
		arr := make([]int, l)
		for i := 0; i < l; i++ {
			arr[i] = l - i
		}
		arrays = append(arrays, arr)
	}
	return arrays
}

func BenchmarkBubbleSort(b *testing.B) {
	for _, slice := range createRandomSlices() {
		b.Run(fmt.Sprintf("random %d", len(slice)), func(_ *testing.B) {
			BubbleSort(slice)
		})
	}
	for _, slice := range createSortedSlices() {
		b.Run(fmt.Sprintf("sorted %d", len(slice)), func(_ *testing.B) {
			BubbleSort(slice)
		})
	}
	for _, slice := range createRevSortedSlices() {
		b.Run(fmt.Sprintf("rev sorted %d", len(slice)), func(_ *testing.B) {
			BubbleSort(slice)
		})
	}
}

func BenchmarkCocktailSort(b *testing.B) {
	for _, slice := range createRandomSlices() {
		b.Run(fmt.Sprintf("random %d", len(slice)), func(_ *testing.B) {
			CocktailSort(slice)
		})
	}
	for _, slice := range createSortedSlices() {
		b.Run(fmt.Sprintf("sorted %d", len(slice)), func(_ *testing.B) {
			CocktailSort(slice)
		})
	}
	for _, slice := range createRevSortedSlices() {
		b.Run(fmt.Sprintf("rev sorted %d", len(slice)), func(_ *testing.B) {
			CocktailSort(slice)
		})
	}
}

func BenchmarkGnomeSort(b *testing.B) {
	for _, slice := range createRandomSlices() {
		b.Run(fmt.Sprintf("random %d", len(slice)), func(_ *testing.B) {
			GnomeSort(slice)
		})
	}
	for _, slice := range createSortedSlices() {
		b.Run(fmt.Sprintf("sorted %d", len(slice)), func(_ *testing.B) {
			GnomeSort(slice)
		})
	}
	for _, slice := range createRevSortedSlices() {
		b.Run(fmt.Sprintf("rev sorted %d", len(slice)), func(_ *testing.B) {
			GnomeSort(slice)
		})
	}

}

func BenchmarkOddEvenSort(b *testing.B) {
	for _, slice := range createRandomSlices() {
		b.Run(fmt.Sprintf("sort random %d", len(slice)), func(_ *testing.B) {
			OddEvenSort(slice)
		})
	}
	for _, slice := range createSortedSlices() {
		b.Run(fmt.Sprintf("sort sorted %d", len(slice)), func(_ *testing.B) {
			OddEvenSort(slice)
		})
	}
	for _, slice := range createRevSortedSlices() {
		b.Run(fmt.Sprintf("sort rev sorted %d", len(slice)), func(_ *testing.B) {
			OddEvenSort(slice)
		})
	}
}

func BenchmarkQuicksort(b *testing.B) {
	for _, slice := range createRandomSlices() {
		b.Run(fmt.Sprintf("sort random %d", len(slice)), func(_ *testing.B) {
			QuickSort(slice)
		})
	}
	for _, slice := range createSortedSlices() {
		b.Run(fmt.Sprintf("sort sorted %d", len(slice)), func(_ *testing.B) {
			QuickSort(slice)
		})
	}
	for _, slice := range createRevSortedSlices() {
		b.Run(fmt.Sprintf("sort rev sorted %d", len(slice)), func(_ *testing.B) {
			QuickSort(slice)
		})
	}
}
