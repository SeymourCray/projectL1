package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	//получаем срез слов
	wordsList := strings.Split(s, " ")
	wordsListLength := len(wordsList)
	//reordering
	for i := 0; i < wordsListLength/2; i++ {
		wordsList[i], wordsList[wordsListLength-i-1] = wordsList[wordsListLength-i-1], wordsList[i]
	}

	return strings.Join(wordsList, " ")
}

func main() {
	example := "snow dog sun"

	fmt.Println(reverseWords(example))
}
