package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	var res int
	for i := range boxTypes {
		if truckSize >= boxTypes[i][0] {
			res += boxTypes[i][1]*boxTypes[i][0]
			truckSize -= boxTypes[i][0]
		} else {
			res += boxTypes[i][1]*truckSize
			truckSize = 0
		}
	}
	return res
}


