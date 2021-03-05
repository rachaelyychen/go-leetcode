package main

import "fmt"

func main() {
	fmt.Println(readBinaryWatch(5))
}

var timeArr = [][]int{{8, 4, 2, 1}, {32, 16, 8, 4, 2, 1}}

func readBinaryWatch(num int) []string {
	var res []string
	readBinary(num, &res, 0, 0, 0, 0)
	return res
}

func readBinary(num int, res *[]string, index1, index2, hour, minute int) {
	if num == 0 {
		*res = append(*res, fmt.Sprintf("%d:%02d", hour, minute))
		return
	}
	if index1 == 1 && index2 > 5 {
		return
	}
	if index1 == 0 && index2 > 3 {
		index1 = 1
		index2 = 0
	}
	readBinary(num, res, index1, index2+1, hour, minute)
	if index1 == 0 {
		if timeArr[index1][index2]+hour > 11 {
			return
		} else {
			hour += timeArr[index1][index2]
		}
	}
	if index1 == 1 {
		if timeArr[index1][index2]+minute > 59 {
			return
		} else {
			minute += timeArr[index1][index2]
		}
	}
	readBinary(num-1, res, index1, index2+1, hour, minute)
}
