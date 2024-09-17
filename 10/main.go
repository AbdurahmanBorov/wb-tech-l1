package main

import (
	"fmt"
)

func temp(temperature []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temperature {
		group := int(temp/10) * 10
		groups[group] = append(groups[group], temp)
	}
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := temp(temperatures)

	fmt.Println(groups)
}
