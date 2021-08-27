package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] < nums[len(nums)-1] {
		return binarySearch(nums, 0, len(nums)-1, target)
	} else {
		var k int
		for i := 0; i < len(nums)-1; i++ {
			if i+1 < len(nums) && nums[i] > nums[i+1] {
				k = i
				break
			}
		}
		if t := binarySearch(nums, 0, k, target); t > -1 {
			return t
		}
		return binarySearch(nums, k+1, len(nums)-1, target)
	}
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
