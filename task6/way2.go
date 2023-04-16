package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		//цикл работает до тех пор, пока канал не будет закрыт
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("Finish")
	}()

	ch <- "Hello"
	ch <- "World"
	close(ch)

	time.Sleep(time.Second * 2)
}
