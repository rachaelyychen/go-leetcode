package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1}))
}

func pivotIndex(nums []int) int {
	var sum int
	leftSum := make([]int, len(nums))
	for i := range nums {
		leftSum[i] = sum
		sum += nums[i]
	}
	for i := range leftSum {
		if leftSum[i] == sum - leftSum[i] - nums[i] {
			return i
		}
	}
	return -1
}
