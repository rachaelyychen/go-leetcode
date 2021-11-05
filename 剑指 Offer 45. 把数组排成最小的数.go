package main

import (
	"sort"
	"strconv"
)

// 时间复杂度O(nlogn)的解法：将非负整数都转换成字符串，定义字符串的比较规则并排序，最后拼接所有字符串得到结果。
// 这里两个字符串的比较规则是，如何拼接起来得到的数最小，也就是两个拼接字符串的排序。

func minNumber(nums []int) string {
	var arr []string
	for i := range nums {
		arr = append(arr, strconv.Itoa(nums[i]))
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] < arr[j]+arr[i]
	})
	var res string
	for i := range arr {
		res += arr[i]
	}
	return res
}
