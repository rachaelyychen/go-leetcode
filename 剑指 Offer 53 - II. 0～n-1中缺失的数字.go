package main

// 时间复杂度O(logn)的解法：对排序数组使用二分查找，如果当前数字等于数组下标，说明缺失的数字在右边，继续往右边查找；
// 如果当前数字大于数组下标，而且前一个数字等于下标，这个数字减去1就是缺失的数字，否则继续往左边查找;
// 如果没有找到，说明数组里的数字和下标都是相等的，缺失的数字是最大的数字。

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if mid == nums[mid] {
			left = mid + 1
		} else if mid < nums[mid] {
			if mid == 0 || mid > 0 && mid-1 == nums[mid-1] {
				return nums[mid] - 1
			}
			right = mid - 1
		}
	}
	return len(nums)
}
