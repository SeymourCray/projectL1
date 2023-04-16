package main

import (
	"fmt"
	"strings"
)

func main() {
	a := `abCdefA`

	res := checkCharacters(a)

	fmt.Println(res)
}

func checkCharacters(s string) bool {
	visited := make(map[rune]bool, 0)
	//к нижнему регистру
	s = strings.ToLower(s)
	//из строк в срез рун
	sl := []rune(s)

	for _, v := range sl {
		_, b := visited[v]
		//если в мапе по такому ключу уже есть значение, возвращаем false
		if b {
			return false
		}
		visited[v] = true
	}
	return true
}
