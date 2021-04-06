package main

import "sort"

func average(salary []int) float64 {
	sort.Slice(salary, func(i, j int) bool {
		return salary[i] < salary[j]
	})
	var sum int
	for i := 1; i < len(salary)-1; i++ {
		sum += salary[i]
	}
	return float64(sum) / float64(len(salary)-2)
}


