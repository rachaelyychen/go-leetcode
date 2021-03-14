package main

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{0}, 100))
}

func addToArrayForm(A []int, K int) []int {
	B := integerToArray(int64(K))
	var t []int
	i, j, c := len(A)-1, len(B)-1, 0
	for i >= 0 || j >= 0 || c > 0 {
		x := c
		if i >= 0 {
			x += A[i]
			i -= 1
		}
		if j >= 0 {
			x += B[j]
			j -= 1
		}
		t = append(t, x%10)
		c = x/10
	}
	var res []int
	for i := len(t)-1; i >= 0; i-- {
		res = append(res, t[i])
	}
	return res
}

func arrayToInteger(array []int) (res int64) {
	length := len(array)
	for i := range array {
		res += int64(array[i]) * Pow10(length-i-1)
	}
	return
}

func integerToArray(integer int64) (array []int) {
	length := 1
	for integer%Pow10(length) != integer {
		length += 1
	}
	array = make([]int, length)
	i := length - 1
	for i >= 0 {
		array[i] = int(integer) % 10
		integer /= 10
		i -= 1
	}
	return
}

func Pow10(p int) int64 {
	res := int64(1)
	for p > 0 {
		res = int64(10*res)
		p -= 1
	}
	return res
}
