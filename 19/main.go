package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "главрыба"

	fmt.Println(reverse(a))
	fmt.Println(reverse2(a))
}

func reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func reverse2(s string) string {
	var builder strings.Builder
	r := []rune(s)

	for i := len(r) - 1; i >= 0; i-- {
		builder.WriteRune(r[i])
	}

	return builder.String()
}
