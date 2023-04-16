package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

//Поведение по умолчанию программы Go, получающей os.Interrupt, заключается в выходе.
//Вызов NotifyContext(parent, os.Interrupt) изменит поведение, чтобы отменить возвращенный контекст

const WorkersNumber = 4

func worker(number int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//получаем значения из канала пока он открыт
	for v := range ch {
		fmt.Printf("Worker №%d got %d\n", number, v)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	//создаем контекст для получение SIGINT
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	//создаем канал
	ch := make(chan int, 5)
	//чтобы горутины завершили всю работу
	var wg sync.WaitGroup
	//запускаем воркеры
	for i := 0; i < WorkersNumber; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	value := 1

	flag := true
	//отправляем значения в канал
	for flag {
		select {
		//если получаем сигнал
		case <-ctx.Done():
			//останавливаем получение уведомлений
			stop()
			//закрываем канал
			close(ch)
			flag = false
		//отправляем новое значение в канал
		case ch <- value:
			value += 1
			time.Sleep(time.Millisecond * 50)
		}
	}

	wg.Wait()
}
