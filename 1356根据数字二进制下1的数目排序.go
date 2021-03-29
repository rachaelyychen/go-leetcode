package main

import "sort"

func sortByBits(arr []int) []int {
	type NumBit struct {
		num int
		bit int
	}
	var numBits []NumBit
	for i := range arr {
		b, t := 0, arr[i]
		for t != 0 {
			if t%2 == 1 {
				b += 1
			}
			t /= 2
		}
		numBits = append(numBits, NumBit{num: arr[i], bit: b})
	}
	sort.Slice(numBits, func(i, j int) bool {
		if numBits[i].bit != numBits[j].bit {
			return numBits[i].bit < numBits[j].bit
		}
		return numBits[i].num < numBits[j].num
	})
	var res []int
	for i := range numBits {
		res = append(res, numBits[i].num)
	}
	return res
}


