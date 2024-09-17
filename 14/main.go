package main

import (
	"fmt"
	"reflect"
)

// приведение типов, использование type switch
func typeAssert(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("variable type: int, value:", v)
	case string:
		fmt.Println("variable type: string, value:")
	case bool:
		fmt.Println("variable type: bool, value:", v)
	case chan int:
		fmt.Println("variable type: chan int, value:", v)
	case chan string:
		fmt.Println("variable type: chan string, value:", v)
	default:
		fmt.Println("unknown", v)
	}
}

func main() {
	num := 10
	str := "hello"
	z := false
	ch := make(chan int)

	typeAssert(num)
	typeAssert(str)
	typeAssert(z)
	typeAssert(ch)

	//использование reflect.ValueOf
	fmt.Println(reflect.TypeOf(ch))
	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(reflect.TypeOf(str))
}
