package main

import "fmt"

func LinearSearch(arr []int, value int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i, nil
		}
	}
	return 0, fmt.Errorf("value %d not found in array", value)
}
