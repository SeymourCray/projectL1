package main

import (
	"fmt"
	"math/rand"
)

// reordering
func quicksort(slice []int, start, end int, r int) {
	if start >= end {
		return
	}

	i := start
	j := end
	q := slice[rand.Intn(end-start)+start]

	for i < j {
		for slice[i] < q {
			i++
		}

		for slice[j] > q {
			j--
		}

		if i < j {
			slice[i], slice[j] = slice[j], slice[i]
			i++
			j--
		}
	}

	if end-start > 1 {
		quicksort(slice, start, i, r+1)
		quicksort(slice, i, end, r+1)
	}
}

// большее потребление памяти
func quicksort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left []int
	var right []int

	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	left = quicksort2(left)
	right = quicksort2(right)

	return append(append(left, pivot), right...)
}

func main() {
	slice := []int{5, 7, 5, 4, 3, 6}

	a := quicksort2(slice)

	fmt.Println(a)

	quicksort(slice, 0, len(slice)-1, 0)

	fmt.Println(slice)
}
