package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//используем контекст для остановки горутины
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			//если возвращается закрытый канал, завершаем работу горутины
			case <-ctx.Done():
				return
			default:
				fmt.Println("do some work")
			}

			time.Sleep(200 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		//закрываем канал
		cancel()
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Finish")
}
