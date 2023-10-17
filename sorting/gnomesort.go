package main

func GnomeSort(arr []int) []int {
	i := 0
	l := len(arr)
	for {
		if i >= l {
			return arr
		}
		if i == 0 || arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}
}
