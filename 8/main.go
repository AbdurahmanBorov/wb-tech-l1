package main

import (
	"fmt"
)

func setBit(n *int64, i int, bitValue int) {
	if bitValue == 1 {
		*n |= 1 << i
	} else {
		*n &= ^(1 << i)
	}
}

func main() {
	var n int64 = 3121

	fmt.Printf("Original value: %b\n", n)

	setBit(&n, 2, 1)
	fmt.Printf("After setting 2nd bit to 1: %b\n", n)

	setBit(&n, 1, 0)
	fmt.Printf("After setting 1st bit to 0: %b\n", n)
}
