package main

import "sort"

func frequencySort(nums []int) []int {
	type numFreq struct {
		num  int
		freq int
	}
	arr1, arr2 := make([]int, 205), make([]numFreq, 0)
	for i := range nums {
		arr1[100+nums[i]] += 1
	}
	for i := range arr1 {
		if arr1[i] > 0 {
			arr2 = append(arr2, numFreq{num: i - 100, freq: arr1[i]})
		}
	}
	sort.Slice(arr2, func(i, j int) bool {
		if arr2[i].freq != arr2[j].freq {
			return arr2[i].freq < arr2[j].freq
		}
		return arr2[i].num > arr2[j].num
	})
	var res []int
	for i := range arr2 {
		for arr2[i].freq > 0 {
			res = append(res, arr2[i].num)
			arr2[i].freq -= 1
		}
	}
	return res
}
