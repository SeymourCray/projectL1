package main

import (
	"fmt"
)

func main() {
	numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	chOne := make(chan int64)
	chTwo := make(chan int64)

	//отправляем значения из массива в первый канал
	go func(numbers []int64) {
		for _, v := range numbers {
			chOne <- v
		}

		close(chOne)
	}(numbers)

	//получаем значения из первого канала, проводим вычисления и отправляем во второй канал
	go func() {
		for {
			for v := range chOne {
				chTwo <- v * v
			}

			close(chTwo)
		}
	}()

	//получаем значения из второго канала и выводим в stdout
	for v := range chTwo {
		fmt.Println(v)
	}
}
