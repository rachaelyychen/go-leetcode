package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numPrimeArrangements(100))
}

func numPrimeArrangements(n int) int {
	var cnt int
	for i := range prime {
		if prime[i] > n {
			cnt = i
			break
		}
	}
	m, x, y := int(math.Pow(float64(10), float64(9))+7), 1, 1
	for i := 2; i <= cnt; i++ {
		x = (x % m * i % m) % m
	}
	for i := 2; i <= n-cnt; i++ {
		y = (y % m * i % m) % m
	}
	fmt.Println(cnt, n-cnt, x, y)
	return (x % m * y % m) % m
}

var prime []int

func init() {
	mark := make([]int, 105)
	for i := 2; i <= 101; i++ {
		if mark[i] == 0 {
			prime = append(prime, i)
			for j := 2 * i; j <= 101; j += i {
				mark[j] = 1
			}
		}
	}
}
