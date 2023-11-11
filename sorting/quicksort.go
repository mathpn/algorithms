package sorting

func QuickSort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	quicksort(arr, lo, pivotIdx-1)
	quicksort(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	idx++
	arr[idx], arr[hi] = pivot, arr[idx]
	return idx
}
