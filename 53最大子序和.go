package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0{
			sum += num
		} else {
			sum = num
		}
		if res < sum {
			res = sum
		}
	}
	return res
}
