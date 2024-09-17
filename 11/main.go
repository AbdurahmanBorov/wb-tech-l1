package main

import "fmt"

func intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for key := range set1 {
		if _, found := set2[key]; found {
			result[key] = struct{}{}
		}
	}

	return result
}

func main() {
	set1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}

	set2 := map[int]struct{}{
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
	}

	result := intersection(set1, set2)

	for key := range result {
		fmt.Println(key)
	}
}
