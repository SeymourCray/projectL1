package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for {
			v, ok := <-ch
			//если канал закрыт, завершаем выполнение функции/горутины
			if !ok {
				fmt.Println("Finish")
				return
			}
			fmt.Println(v)
		}
	}()

	ch <- "Hello"
	ch <- "World!"
	close(ch)

	time.Sleep(time.Second * 2)
}
