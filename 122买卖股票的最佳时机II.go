package main

import "fmt"

func main() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit2(prices []int) int {
	var res int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}
