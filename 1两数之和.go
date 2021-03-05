package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

type NumIndex struct {
	Num   int
	Index int
}

func twoSum(nums []int, target int) []int {
	var arr []NumIndex
	for i := range nums {
		arr = append(arr, NumIndex{Num: nums[i], Index: i})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Num < arr[j].Num
	})
	i := 0
	j := len(nums) - 1
	for ; i < j; {
		tot := arr[i].Num + arr[j].Num
		if tot == target {
			return []int{arr[i].Index, arr[j].Index}
		} else if tot < target {
			i += 1
		} else {
			j -= 1
		}
	}
	return nil
}
