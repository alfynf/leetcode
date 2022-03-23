package main

import "fmt"

func containsDuplicate(nums []int) bool {
	data := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		data[nums[i]]++
		if data[nums[i]] == 2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
