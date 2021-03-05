package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{0, 0, 0, 0, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	var res int
	i := 0
	for i < len(nums) {
		if nums[i] != 1 {
			i += 1
			continue
		}
		t := 1
		for j := i + 1; j <= len(nums); j++ {
			if j == len(nums) || nums[j] != nums[i] {
				i = j
				break
			}
			t += 1
		}
		if t > res {
			res = t
		}
	}
	return res
}
