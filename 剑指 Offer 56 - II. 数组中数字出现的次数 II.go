package main

import "math"

// 除了一个数字出现一次，其它数字都出现三次，对于二进制的每一位，遍历数组算出累积和，如果能被3整除说明单独数字的该二进制位是0，否则是1；
// 得到二进制数，即可得出十进制数。

func singleNumber(nums []int) int {
	var x int
	t, n := 1, 0
	for n < 65 {
		n += 1
		var sum int
		for i := range nums {
			sum += t & nums[i]
		}
		t <<= 1
		if sum%3 > 0 {
			x += int(math.Pow(float64(2), float64(n-1)))
		}
	}
	return x
}
