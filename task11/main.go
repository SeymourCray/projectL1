package main

import "fmt"

func intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	res := make(map[int]struct{})

	for i := range set2 {
		if _, ok := set1[i]; ok {
			res[i] = struct{}{}
		}
	}

	return res
}

func main() {
	slice1 := []int{1, 7, 6, 5, 8, 9, 3}
	slice2 := []int{0, 6, 7, 4, 2}

	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	//struct{}{} занимают минимальное кол-во памяти, поэтому можно использовать мапу в качестве множества
	for i := 0; i < len(slice1); i++ {
		set1[slice1[i]] = struct{}{}
	}

	for i := 0; i < len(slice2); i++ {
		set2[slice2[i]] = struct{}{}
	}

	fmt.Println(intersection(set1, set2))
}
