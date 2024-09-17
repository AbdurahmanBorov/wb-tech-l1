package main

import (
	"fmt"

	"wb-tech-L1/24/point"
)

func main() {
	p := point.New(10, 20)

	fmt.Println(p.Distance())
}
