package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for v := range slice {
		set[slice[v]] = struct{}{}
	}

	fmt.Println(set)
}
