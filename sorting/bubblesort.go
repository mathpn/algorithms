package main

import "fmt"

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 10, 3, 2, 12, 5, 2, 8, 5}
	arr = BubbleSort(arr)
	fmt.Printf("%v\n", arr)
}
