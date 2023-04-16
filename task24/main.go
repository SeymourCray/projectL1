package main

import (
	"fmt"
	"projectL1/task24/point"
)

func main() {
	a := point.NewPoint(1, 2)
	b := point.NewPoint(5, 4)

	fmt.Println(point.FindDistance(a, b))
}
