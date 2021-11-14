package main

// 双指针指向连续正数序列的首尾两数，初始位置在1、2，优先向右移动right（增大连续和），其次向右移动left（减小连续和）。

func findContinuousSequence(target int) [][]int {
	if target < 3 {
		return [][]int{}
	}
	var res [][]int
	left, right, sum := 1, 2, 3
	for left < right {
		if left > target/2 {
			break
		}
		if sum == target {
			var t []int
			for i := left; i <= right; i++ {
				t = append(t, i)
			}
			res = append(res, t)
			right += 1
			sum += right
		} else if sum < target {
			right += 1
			sum += right
		} else {
			sum -= left
			left += 1
		}
	}
	return res
}
