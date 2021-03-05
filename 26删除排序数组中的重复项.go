package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}
