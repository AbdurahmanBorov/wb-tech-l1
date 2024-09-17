package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	//sleep(1 * time.Second)
	//sleep2(1 * time.Second)

	fmt.Println(time.Since(start))
}

func sleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func sleep2(d time.Duration) {
	<-time.After(d)
}
