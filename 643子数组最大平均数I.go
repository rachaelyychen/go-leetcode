package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	var maxSum int
	var sum int
	for i := 0; i < len(nums); i++ {
		j := i + k - 1
		if j == len(nums) {
			break
		}
		if i == 0 {
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			maxSum = sum
		} else {
			sum = sum - nums[i-1] + nums[j]
		}
		fmt.Println(i, j, sum)
		if maxSum < sum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}
