package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2)
	a.Exp(a, big.NewInt(20), nil)

	b := big.NewInt(3)
	b.Exp(b, big.NewInt(20), nil)

	mul := new(big.Int).Mul(a, b)
	fmt.Println(mul)

	div := new(big.Int).Div(a, b)
	fmt.Println(div)

	sum := new(big.Int).Add(a, b)
	fmt.Println(sum)

	sub := new(big.Int).Sub(a, b)
	fmt.Println(sub)
}
