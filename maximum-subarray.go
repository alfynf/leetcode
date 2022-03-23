package main

import "fmt"

func maxSubArray(nums []int) int {
	var max int = -9999999
	var curr_max int = 0
	for i := 0; i < len(nums); i++ {
		if curr_max < 0 {
			curr_max = 0
		}
		curr_max = curr_max + nums[i]
		if curr_max > max {
			max = curr_max
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
