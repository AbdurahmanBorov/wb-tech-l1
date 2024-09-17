package main

import "fmt"

// с сохранением порядка
func removeElem(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

// без сохранения
func removeElem2(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(removeElem(arr, 3))

	arr2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(removeElem2(arr2, 3))
}
