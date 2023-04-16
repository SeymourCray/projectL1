package main

import (
	"fmt"
)

func binarySearch(slice []int, v int) int {
	start := 0
	end := len(slice) - 1

	var index int

	for start <= end {
		//при делении типа int остается целая часть
		middle := (start + end) / 2

		if slice[middle] > v {
			end = middle - 1
		} else if slice[middle] < v {
			start = middle + 1
		} else {
			index = middle
			break
		}
	}

	return index
}

func main() {
	items := []int{1, 2, 3, 9, 20, 31, 45, 1, 70, 100}
	fmt.Println(binarySearch(items, 63))
}
