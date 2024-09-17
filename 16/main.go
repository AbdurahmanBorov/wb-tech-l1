package main

import "fmt"

// первый способ
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	var left, right []int

	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

// второй способ
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort2(arr []int, low, high int) []int {
	if low < high {
		arr, i := partition(arr, low, high)

		quickSort2(arr, low, i-1)
		quickSort2(arr, i+1, high)
	}
	return arr
}

func main() {
	arr := []int{33, 10, 55, 71, 29, 15, 89, 77}
	fmt.Println("base array: ", arr)

	//первый способ
	sortedArr := quickSort(arr)
	fmt.Println("sorted array:", sortedArr)

	//второй
	fmt.Println(quickSort2(arr, 0, len(arr)-1))
}
