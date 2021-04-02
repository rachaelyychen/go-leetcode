package main

func finalPrices(prices []int) []int {
	var res []int
	for i := 0; i < len(prices); i++ {
		t := prices[i]
		for j := i+1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				t -= prices[j]
				break
			}
		}
		res = append(res, t)
	}
	return res
}


