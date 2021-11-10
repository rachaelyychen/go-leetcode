package main

// 先对所有数字进行一次异或，得到两个出现一次的数字的异或值，在异或结果中找到任意为 1 的位；
// 根据这一位对所有的数字进行分组，在每个组内进行异或操作，得到两个数字。

func singleNumbers(nums []int) []int {
	var x int
	for i := 0; i < len(nums); i++ {
		x ^= nums[i]
	}
	t := 1
	for t&x == 0 {
		t <<= 1
	}
	a, b := 0, 0
	for i := range nums {
		if t&nums[i] == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}
