package main

import "fmt"

func main() {
	a := 0
	b := 4

	a = a - b
	b = b + a
	a = b - a

	fmt.Println(a, b)

	c := 1.1
	d := 2.2

	c = c * d
	d = c / d
	c = c / d

	fmt.Println(c, d)
}
