package main

import (
	"sort"
)

func tictactoe(moves [][]int) string {
	var arrA, arrB [][]int
	for i := range moves {
		if i%2 == 0 {
			arrA = append(arrA, moves[i])
		} else {
			arrB = append(arrB, moves[i])
		}
	}
	if checkTictactoeWin(arrA) {
		return "A"
	}
	if checkTictactoeWin(arrB) {
		return "B"
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func checkTictactoeWin(arr [][]int) bool {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] != arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		if arr[i][1] != arr[j][1] {
			return arr[i][1] < arr[j][1]
		}
		return true
	})
	t1, t2 := 0, 0
	for i := 0; i < len(arr); i++ {
		if i+2 < len(arr) {
			if arr[i][0] == arr[i+1][0] && arr[i+1][0] == arr[i+2][0] {
				return true
			}
		}
		// 正对角线：[0,2],[1,1],[2,0]
		// 负对角线：[0,0],[1,1],[2,2]
		if arr[i][0] == 0 && arr[i][1] == 2 || arr[i][0] == 1 && arr[i][1] == 1 || arr[i][0] == 2 && arr[i][1] == 0 {
			t1 += 1
			if t1 == 3 {
				return true
			}
		}
		if arr[i][0] == 0 && arr[i][1] == 0 || arr[i][0] == 1 && arr[i][1] == 1 || arr[i][0] == 2 && arr[i][1] == 2 {
			t2 += 1
			if t2 == 3 {
				return true
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] != arr[j][1] {
			return arr[i][1] < arr[j][1]
		}
		if arr[i][0] != arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		return true
	})
	for i := 0; i < len(arr); i++ {
		if i+2 < len(arr) {
			if arr[i][1] == arr[i+1][1] && arr[i+1][1] == arr[i+2][1] {
				return true
			}
		}
	}
	return false
}
