package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/22/21 10:52 PM
**/

// 时间复杂度	O(n)，空间复杂度O(1)的解法：双指针分别指向数组头和尾，开始遍历数组，
// 左指针指向奇数、右指针指向偶数时同时移动，左指针指向奇数、右指针指向奇数时移动左指针，
// 左指针指向偶数、右指针指向偶数时移动右指针，左指针指向偶数、右指针指向奇数时交换两数再同时移动。

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		a, b := nums[i]&0x1, nums[j]&0x1
		if a == 1 && b == 0 {
			i += 1
			j -= 1
		} else if a == 1 && b == 1 {
			i += 1
		} else if a == 0 && b == 0 {
			j -= 1
		} else {
			t := nums[i]
			nums[i] = nums[j]
			nums[j] = t
			i += 1
			j -= 1
		}
	}
	return nums
}
