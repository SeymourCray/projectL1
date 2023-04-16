package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	numbers := [5]int64{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	var sum int64

	for _, v := range numbers {
		wg.Add(1)
		//гарантирование взаимного исключения
		go func(v int64) {
			atomic.AddInt64(&sum, v*v)
			wg.Done()
		}(v)
	}

	wg.Wait()

	fmt.Println(sum)
}
