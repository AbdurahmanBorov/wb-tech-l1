package main

import "fmt"

func main() {
	a := 1
	b := 2

	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)
}
