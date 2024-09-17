package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, dataChan <-chan int, doneChan <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				return
			}
			fmt.Printf("Worker %d received: %d\n", id, data)
		case <-doneChan:
			return
		}
	}
}

func main() {
	dataChan := make(chan int)
	doneChan := make(chan struct{})

	numWorkers := 3
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, doneChan, &wg)
	}

	go func() {
		defer close(dataChan)
		for i := 0; ; i++ {
			select {
			case <-doneChan:
				return
			default:
				dataChan <- i
				time.Sleep(1 * time.Second)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	close(doneChan)

	wg.Wait()
	fmt.Println("All workers have finished.")
}
