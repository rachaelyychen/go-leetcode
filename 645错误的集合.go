package main

import (
	"fmt"
)

func main() {
	fmt.Println(findErrorNums([]int{3,2,3,4,6,5}))
}

func findErrorNums(nums []int) []int {
	var res []int
	arr := make([]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if arr[nums[i]] == true {
			res = append(res, nums[i])
		} else {
			arr[nums[i]] = true
		}
	}
	for i := range arr {
		if i >0 && arr[i] == false {
			res = append(res, i)
			break
		}
	}
	return res
}
