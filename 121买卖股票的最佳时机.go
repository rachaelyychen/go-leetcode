package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var res int
	if len(prices) == 0 {
		return res
	}
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] - min > res {
			res = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}
