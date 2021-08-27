package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	if len(nums) >= 4 {
		for a := 0; a < len(nums); a++ {
			if a > 0 && nums[a-1] == nums[a] {
				continue
			}
			for b := a + 1; b < len(nums); b++ {
				if b-1 > a && nums[b-1] == nums[b] {
					continue
				}
				c, d := b+1, len(nums)-1
				for c < d {
					sum := nums[a] + nums[b] + nums[c] + nums[d]
					if sum == target {
						res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
						c += 1
						for c < d && nums[c-1] == nums[c] {
							c += 1
						}
						d -= 1
						for c < d && nums[d+1] == nums[d] {
							d -= 1
						}
					} else if sum < target {
						c += 1
						for c < d && nums[c-1] == nums[c] {
							c += 1
						}
					} else {
						d -= 1
						for c < d && nums[d+1] == nums[d] {
							d -= 1
						}
					}
				}
			}
		}
	}
	return res
}
