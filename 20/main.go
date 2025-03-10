package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	words := strings.Split(str, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"

	fmt.Println(reverse(str))
}
