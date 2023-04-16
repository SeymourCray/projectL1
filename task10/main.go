package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	//map для распределения температур с шагом в 10
	tempMap := make(map[int][]float64)

	for n := range temperatures {
		//получаем значение шага
		k := int(math.Trunc(temperatures[n]/10) * 10)
		//добавляем в map
		tempMap[k] = append(tempMap[k], temperatures[n])
	}

	fmt.Println(tempMap)
}
