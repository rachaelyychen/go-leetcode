package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/19/21 3:23 PM
**/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	first, second, t, m := 0, 1, 0, 1000000007
	for i := 2; i <= n; i++ {
		t = (first%m + second%m) % m
		first = second
		second = t
	}
	return t
}
