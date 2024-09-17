package main

import (
	"fmt"
	"sync"
)

func square(num int) int {
	return num * num
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	sum := 0

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for _, n := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			sum += square(n)
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
