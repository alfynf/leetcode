package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var index int

	if len(nums) == 1 && nums[0] != target {
		return -1
	}

	if target > nums[len(nums)-1] || target < nums[0] {
		return -1
	}

	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			index = mid
			return index
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	// fmt.Println(1 / 2)
}
