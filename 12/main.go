package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, v := range arr {
		set[v] = struct{}{}
	}

	fmt.Println("result:")
	for k := range set {
		fmt.Println(k)
	}
}
