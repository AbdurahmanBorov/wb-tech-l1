package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	t := time.After(5 * time.Second)

	go func() {
		i := 0
		for {
			select {
			case <-t:
				close(ch)
				return
			default:
				ch <- i
				fmt.Println("Отправлено:", i)
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for val := range ch {
		fmt.Println("Получено:", val)
		time.Sleep(1 * time.Second)
	}
}
