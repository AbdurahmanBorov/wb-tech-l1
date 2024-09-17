package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 10, 22, 31, 34, 93}

	fmt.Println(binarySearch(arr, 22))
}
