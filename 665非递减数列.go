package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{3,4,2,3}))
}

func checkPossibility(nums []int) bool {
	// 4 3 2
	// 4 2 3
	// 3 4 2 3
	var cnt int
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] > nums[i+1] {
			cnt += 1
			if cnt > 1 {
				return false
			}
			if i > 0 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return cnt < 2
}
