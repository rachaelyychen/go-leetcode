package main

import "fmt"

func main() {
	fmt.Println(findLHS([]int{1, 2, 2, 1}))
}

func findLHS(nums []int) int {
	numCntMap := make(map[int]int, 0)
	for i := range nums {
		if cnt, exists := numCntMap[nums[i]]; !exists {
			numCntMap[nums[i]] = 1
		} else {
			numCntMap[nums[i]] = cnt + 1
		}
	}
	var res int
	for num, cnt := range numCntMap {
		if cnt1, exists := numCntMap[num+1]; exists {
			if res < cnt+cnt1 {
				res = cnt + cnt1
			}
		}
	}
	return res
}
