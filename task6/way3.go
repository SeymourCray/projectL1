package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	//добавляем значение в канал ch и ждем, пока не получим значение из канала done, чтобы завершить работу горутины
	go func() {
		count := 0

		for {
			select {
			case ch <- count:
			case <-done:
				close(ch)
				return
			}

			count += 1

			time.Sleep(200 * time.Millisecond)
		}
	}()

	//ждем 3 секунды и добавляем значение(семафору) в канал done
	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	//получаем значения из канала, пока он не будет закрыт
	for i := range ch {
		fmt.Printf("Received: %d\n", i)
	}

	fmt.Println("Finish")
}
