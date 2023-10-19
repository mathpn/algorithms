package main

import (
	"fmt"
)

func BinarySearch(arr []int, value int) (int, error) {
	l := 0
	r := len(arr) - 1
	for {
		if l > r {
			return 0, fmt.Errorf("value %d not found in array", value)
		}
		m := (l + r) / 2
		if arr[m] < value {
			l = m + 1
		} else if arr[m] > value {
			r = m - 1
		} else {
			return m, nil
		}
	}
}
