package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Duration(10) * time.Second
	mySleep(d)
	fmt.Println("sleep is over")
}

func mySleep(d time.Duration) {
	//ждет определенное время и возвращает канал
	<-time.After(d)
}
