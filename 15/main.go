package main

import "fmt"

//Проблема была в утечке памяти, копируется только часть строки(100 символов)

var justString string

func createString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createString(1 << 10)
	justString = string(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
