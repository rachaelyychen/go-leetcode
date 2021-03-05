package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			if i+1 < len(nums) && nums[i+1] > target || i+1 == len(nums) {
				return i + 1
			}
		}
	}
	return 0
}
