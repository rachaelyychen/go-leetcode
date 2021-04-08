package main

import "fmt"

func main() {
	fmt.Println(mostVisited(2, []int{2,1,2,1,2,1,2,1,2}))
}

func mostVisited(n int, rounds []int) []int {
	arr := make([]int, 105)
	maxCnt, res := 0, make([]int, 0)
	j, i := 0, rounds[0]
	for {
		if j == len(rounds) {
			break
		}
		t := i%(n+1)
		arr[t] += 1
		if arr[t] > maxCnt {
			maxCnt = arr[t]
		}
		if t == rounds[j] {
			j += 1
		}
		i += 1
	}
	for i := range arr {
		if i > 0 && arr[i] == maxCnt {
			res = append(res, i)
		}
	}
	return res
}


