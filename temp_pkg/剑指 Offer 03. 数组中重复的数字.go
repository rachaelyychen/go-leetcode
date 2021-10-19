package temp_pkg

import "fmt"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/17/21 8:06 PM
**/

// 长度为n的数组中数字都在0～n-1范围内，存在至少一个数字是重复的，找出其中任意一个重复的数字。
// 时间复杂度O(n)，空间复杂度O(1)的解法：遍历数组，目标是将每个数字放在对应下标，
// 判断如果数字和下标相同，则继续遍历，如果数字和下标不同，则交换此下标和对应下标的两个数字，重复判断，
// 过程中发现相同的数字则返回。

func main() {
	fmt.Println(duplicateNumber([]int{4, 3, 1, 0, 2, 3}))
}

func duplicateNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			t := nums[nums[i]]
			nums[nums[i]] = nums[i]
			nums[i] = t
			i -= 1
		}
	}
	return -1
}
