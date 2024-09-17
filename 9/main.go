package main

import (
	"fmt"
)

func producer(numbers []int, in chan<- int) {
	for _, number := range numbers {
		in <- number
	}
	close(in)
}

func multiplier(in <-chan int, out chan<- int) {
	for number := range in {
		out <- number * 2
	}
	close(out)
}

func consumer(out <-chan int) {
	for result := range out {
		fmt.Println(result)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	in := make(chan int)
	out := make(chan int)

	go producer(numbers, in)
	go multiplier(in, out)
	consumer(out)
}
