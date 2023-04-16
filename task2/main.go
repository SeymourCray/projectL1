package main

import (
	"fmt"
	"math"
	"sync"
)

func calculateSquare(value float64) {
	res := math.Pow(value, 2)

	//выводим в stdout
	fmt.Println(res)
}

func main() {
	numbers := [5]float64{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, i := range numbers {
		//добавляем значение в счетчик
		wg.Add(1)

		//горутина
		go func(number float64) {
			// отложенно убавляем единичку из счетчика
			defer wg.Done()
			calculateSquare(number)
		}(i)
	}

	//ждем завершение работы горутин (счетчик равен 0)
	wg.Wait()
}
