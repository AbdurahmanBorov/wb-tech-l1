package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	wg := sync.WaitGroup{}
	counter := Counter{}

	numGoroutines := 10
	iterations := 100

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				counter.Inc()
			}
		}()
	}
	wg.Wait()

	fmt.Println(counter.Value())
}
