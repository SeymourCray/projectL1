package main

import (
	"fmt"
	"strings"
)

// использование глобальных переменных должно быть оправданным

func createHugeString(v int) string {
	// строки в го иммутабельны, соответственно используем strings.Builder (аналог bytes.Buffer)
	var builder strings.Builder
	builder.Grow(v)
	for i := 0; i < v; i++ {
		builder.WriteString("a")
	}

	return builder.String()
}

func someFunc(s *[]rune) {
	v := createHugeString(1 << 10)
	*s = []rune(v[:100])
}

func main() {
	var justString []rune

	someFunc(&justString)

	fmt.Println(string(justString))
}
