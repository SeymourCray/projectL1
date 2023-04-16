package main

import (
	"fmt"
	"math"
)

func calculateSquare(value float64, c chan<- float64) {
	res := math.Pow(value, 2)

	//добавляем в канал
	c <- res
}

func main() {
	numbers := [5]float64{2, 4, 6, 8, 10}

	c := make(chan float64)

	for _, i := range numbers {
		//запускаем горутину для каждого значения
		go calculateSquare(i, c)
	}

	var sum float64

	//ожидаем, пока получим все 5 значений из канала и прибавляем
	//Share memory by communicating; don't communicate by sharing memory
	for _ = range numbers {
		sum += <-c
	}

	close(c)

	fmt.Println(sum)
}
