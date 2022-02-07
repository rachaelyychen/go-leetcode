package main

import . "go-leetcode/kit"

// 是否存在满足条件 i < j < k 且 nums[i] < nums[k] < nums[j] 的子序列
// 时间复杂度O(n)的解法:单调栈
// 求任何位置的左边最小的元素 nums[i] ，可以提前遍历一次得到；
// 使用「单调递减栈」，把 nums[j]  入栈时，需要把栈里面比它小的元素全都 pop 出来，由于越往栈底越大，所以 pop 出的最后一个元素，就是比 3 小的最大元素 nums[k] ;
// 判断如果 nums[i] < nums[k] ，那就说明得到了一个 132 模式。
// 因为单调栈是建立在 3 的右边的，因此从右向左遍历。

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	leftMinArr := make([]int, len(nums))
	leftMin := nums[0]
	for i := 1; i < len(nums); i++ {
		leftMinArr[i] = leftMin
		if leftMin > nums[i] {
			leftMin = nums[i]
		}
	}
	rightMaxStack := NewStack()
	rightMaxStack.Push(nums[len(nums)-1])
	for j := len(nums) - 2; j > 0; j-- {
		rightMax := leftMinArr[j]
		for !rightMaxStack.IsEmpty() {
			t := rightMaxStack.Top().(int)
			if t < nums[j] {
				rightMaxStack.Pop()
				rightMax = t
			} else {
				break
			}
		}
		if leftMinArr[j] < rightMax {
			return true
		}
		rightMaxStack.Push(nums[j])
	}
	return false
}

// 时间复杂度O(n*n)的解法:暴力+贪婪

func find132pattern2(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	leftMin := nums[0]
	for j := 1; j < len(nums)-1; j++ {
		if leftMin >= nums[j] {
			leftMin = nums[j]
			continue
		}
		for k := j + 1; k < len(nums); k++ {
			if leftMin < nums[k] && nums[k] < nums[j] {
				return true
			}
		}
	}
	return false
}
