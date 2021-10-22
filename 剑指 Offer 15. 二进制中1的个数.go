package main

import "fmt"

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/21/21 4:28 PM
**/

// 三种解法：
// 第一种解法：将二进制数不断右移一位，直到为0，只需要判断最右边一位是否为1就可以统计出一共有多少个1；
// 判断最右边一位是不是1，可以和1做与运算，如果结果是1，说明是1，否则是0；
// 需要注意的是，这种方法不能用于负数，因为负数右移会补位1。
// 第二种解法：将1不断左移一位，直到为0，检查每次与运算的结果是否为0判断二进制数该位是否为1，如果结果不为0，说明是1，否则是0；
// 这种解法可用于负数，但二进制数有多少位就需要判断多少次。
// 第三种解法：将二进制数减1再与自身做与运算，可以将它最右边的1去掉，不断重复这个过程（1110->1100->1000->0000）直到为0，能重复多少次说明有多少个1；
// 这种解法有多少个1就判断多少次。

// 第三种解法
func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		num = (num - 1) & num
		res += 1
	}
	return res
}

// 第一种解法
func hammingWeight1(num uint32) int {
	var res int
	for num != 0 {
		if num&1 == 1 {
			res += 1
		}
		num >>= 1
	}
	return res
}

// 一条语句判断一个数是不是2的整数次方：(num-1)&num==0
// 两个整数m和n，改变m二进制的多少位才能得到n：先把m和n做异或(m^n)，相同为0相异为1，再计算异或结果中1的个数。
func main() {
	m, n := 10, 13
	fmt.Println(hammingWeight(uint32(m ^ n)))
}
