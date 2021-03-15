package main

func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	var num int
	for i := range A {
		// 最末尾数字为0或5才能被5整除
		num = (num*2 + A[i])%10
		res[i] = num%5 == 0
	}
	return res
}


