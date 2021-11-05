package main

// 有一个数字出现的次数超过一半，首先用第一个数字做partition，将小于等于它的数字移到左边，大于它的数字移到右边，如果这个数字来到数组中间，它就是要找的，返回即可；
// 如果这个数字在数组后半部分，说明要找的数字小于它，在它的左边继续查找；如果这个数字在数组前半部分，说明要找的数字大于它，在它的右边继续查找。
// 时间复杂度O(n)，空间复杂度O(1)，注意这个解法会修改原数组。

func majorityElement(nums []int) int {
	left, right := 0, len(nums)-1
	t := partition(nums, left, right)
	for t != len(nums)/2 {
		if t > len(nums)/2 {
			right = t - 1
		} else {
			left = t + 1
		}
		t = partition(nums, left, right)
	}
	return nums[t]
}

func partition(nums []int, left, right int) int {
	t := nums[left]
	for left < right {
		for left < right && nums[right] > t {
			right -= 1
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= t {
			left += 1
		}
		nums[right] = nums[left]
	}
	nums[left] = t
	return left
}
