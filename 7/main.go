package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			m[i] = i * i
		}()
	}

	wg.Wait()

	for i, el := range m {
		fmt.Println(i, el)
	}
}
