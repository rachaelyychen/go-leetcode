package main

// 时间复杂度O(logn)的解法：对排序数组使用二分查找，分别找到该数字第一个和最后一个的下标，从而得出重复的个数。

func search(nums []int, target int) int {
	i, j := searchFirstTarget(nums, target), searchLastTarget(nums, target)
	if i == -1 || j == -1 {
		return 0
	}
	return j - i + 1
}

func searchFirstTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == 0 || mid > 0 && nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}

func searchLastTarget(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == len(nums)-1 || mid < len(nums)-1 && nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}
