package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/19/21 3:43 PM
**/

// 青蛙跳问题、小矩形覆盖问题都是斐波那契数列的变形。
// 另一种青蛙跳问题：青蛙一次可以跳1级台阶，可以跳2级台阶...可以跳n级台阶，那么它跳上n级台阶共有多少种跳法？
// 数学归纳法可以证明f(n)=2^(n-1) (n>0)

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	first, second, t, m := 1, 1, 0, 1000000007
	for i := 2; i <= n; i++ {
		t = (first%m + second%m) % m
		first = second
		second = t
	}
	return t
}
