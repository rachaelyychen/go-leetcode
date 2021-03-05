package main

import "fmt"

func main() {
fmt.Println(removeElement([]int{3,2,2,3}, 3))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if val != nums[j] {
			nums[i] = nums[j]
			i += 1
		}
	}
	return i
}
