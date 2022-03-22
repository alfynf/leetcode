package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	} else if target < nums[0] {
		return 0
	} else {
		for left < right {
			mid := left + ((right - left) / 2)
			if target == nums[mid] {
				return mid
			} else if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				right = mid - 1
			}
		}
	}
	return left + 1
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}
