package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseBits(00000010100101000001111010011100))
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// 二进制最后一位bit
		bit := num&1
		// 右移一位
		num = num >> 1
		// 左移一位加上bit
		res = (res << 1) + bit
	}
	return res
}
