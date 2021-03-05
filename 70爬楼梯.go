package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
}

var m = make(map[int]int, 0)

func climbStairs(n int) int {
	if n <= 2 {
		m[n] = n
		return n
	}
	var res int
	if v, exists := m[n]; !exists {
		res = climbStairs(n-1)+climbStairs(n-2)
		m[n] = res
	} else {
		res = v
	}
	return res
}
