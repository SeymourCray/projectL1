package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const N = 5 * time.Second

func main() {
	//контекст закрывается после времени N
	ctx, cancel := context.WithTimeout(context.Background(), N)
	defer cancel()

	ch := make(chan int, 5)

	var wg sync.WaitGroup

	wg.Add(1)
	//создаем горутину, читающую значение из канала
	go func(ch chan int, wg *sync.WaitGroup) {
		for v := range ch {
			fmt.Println(v)
			time.Sleep(time.Second)
		}

		wg.Done()
	}(ch, &wg)

	var value int

	flag := true

	for flag {
		select {
		//когда пройдет N времени, канал закроется
		case <-ctx.Done():
			close(ch)
			flag = false
		case ch <- value:
			value += 1
		}
	}

	wg.Wait()
}
