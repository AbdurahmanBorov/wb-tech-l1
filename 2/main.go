package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{}

	for _, n := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(n * n)
		}()
	}

	wg.Wait()
}
