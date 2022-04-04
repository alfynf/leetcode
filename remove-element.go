package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 1 && nums[0] == val {
		return 0
	}
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == val {
			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
			nums = nums[:len(nums)-1]
			i = i - 1
		}
	}
	return len(nums)
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
