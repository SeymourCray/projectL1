package main

import "fmt"

func setBit(n int64, pos uint) int64 {
	//используем побитовое или
	n |= 1 << pos
	return n
}

func clearBit(n int64, pos uint) int64 {
	//используем побитовое "и" и "xor"
	n &= ^(1 << pos)
	return n
}

func main() {
	var number int64 = 100

	number = setBit(number, 10)

	fmt.Println(number)

	number = clearBit(number, 10)

	fmt.Println(number)
}
