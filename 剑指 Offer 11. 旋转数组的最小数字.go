package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/19/21 3:47 PM
**/

// 时间复杂度O(logn)的解法：利用旋转数组的特性（左侧增序数组+右侧增序数组，最小数字在连接处），使用二分查找不断缩小范围，
// 如果mid数字比left数字大，说明mid处于左侧排序数组中，最小数字在mid右边，于是令left=mid；
// 如果mid数字比right数字小，说明mid处于右侧排序数组中，最小数字在mid左边，于是令right=mid；
// 不断重复查找过程，直到left在right左边，返回right数字，即最小数字。
// 另外需要处理两种情况：1、非旋转数组；2、当mid、left、right三个数字相等时，无法判断最小数字在哪边，只能改换顺序遍历查找。

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if numbers[0] < numbers[len(numbers)-1] {
		return numbers[0]
	}
	left, right := 0, len(numbers)-1
	for left < right-1 {
		mid := (left + right) / 2
		if numbers[mid] == numbers[left] && numbers[mid] == numbers[right] {
			t := numbers[0]
			for i := 1; i < len(numbers); i++ {
				if t > numbers[i] {
					t = numbers[i]
				}
			}
			return t
		}
		if numbers[mid] >= numbers[left] {
			left = mid
			continue
		}
		if numbers[mid] <= numbers[right] {
			right = mid
		}
	}
	return numbers[right]
}
