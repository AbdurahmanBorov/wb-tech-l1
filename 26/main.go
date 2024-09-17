package main

import (
	"fmt"
)

func unique(str string) bool {
	r := []rune(str)
	length := len(r)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if r[i] == r[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(unique(str1))
	fmt.Println(unique(str2))
	fmt.Println(unique(str3))
}
