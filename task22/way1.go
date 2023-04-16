package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := math.Pow(3, 30)
	b := math.Pow(2, 29)

	fmt.Printf("Вычитание: %v\n", a-b)
	fmt.Printf("Сумма: %v\n", a+b)
	fmt.Printf("Умножение: %v\n", a*b)
	fmt.Printf("Деление: %v\n", a/b)

	c := big.NewFloat(math.Pow(3, 40))
	d := big.NewFloat(math.Pow(2, 50))

	var res big.Float
	fmt.Println(res.Sub(c, d))
	fmt.Println(res.Add(c, d))
	fmt.Println(res.Mul(c, d))
	fmt.Println(res.Quo(c, d))
}
