package main

import "fmt"

func main() {
fmt.Println(containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}))
}

func containsDuplicate(nums []int) bool {
	numsCntMap := make(map[int]int, 0)
	for i := range nums {
		if _, exists := numsCntMap[nums[i]]; exists {
			return true
		} else {
			numsCntMap[nums[i]] = 1
		}
	}
	return false
}
