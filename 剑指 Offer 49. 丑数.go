package main

// 利用丑数的递推性质：丑数只包含因子2、3、5，因此有 “丑数 == 某较小丑数 × 某因子” ，从1开始不断递推出下一个丑数。

func init() {
	nthUglyNumberArr = make([]int, 1690)
	nthUglyNumberArr[0] = 1
	i, j, k, index := 0, 0, 0, 1
	for index < len(nthUglyNumberArr) {
		t := nthUglyNumberArr[i] * 2
		if t > nthUglyNumberArr[j]*3 {
			t = nthUglyNumberArr[j] * 3
		}
		if t > nthUglyNumberArr[k]*5 {
			t = nthUglyNumberArr[k] * 5
		}
		if t == nthUglyNumberArr[i]*2 {
			i += 1
		}
		if t == nthUglyNumberArr[j]*3 {
			j += 1
		}
		if t == nthUglyNumberArr[k]*5 {
			k += 1
		}
		nthUglyNumberArr[index] = t
		index += 1
	}
}

func nthUglyNumber(n int) int {
	return nthUglyNumberArr[n-1]
}

var nthUglyNumberArr []int
