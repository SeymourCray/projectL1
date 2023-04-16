package main

import "fmt"

func reverseRow(s string) string {
	//из строки получаем срез рун
	runes := []rune(s)
	end := len(runes)
	//reordering
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[end-1-i] = runes[end-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	example := "simple string"

	example = reverseRow(example)

	fmt.Println(example)
}
