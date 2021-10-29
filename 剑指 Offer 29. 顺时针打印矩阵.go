package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/24/21 10:22 AM
**/

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	num := row * col
	res, index, cnt := make([]int, num), 0, 0
	for {
		rowFrom, rowTo, colFrom, colTo := cnt, row-cnt, cnt, col-cnt
		if rowFrom >= rowTo || colFrom >= colTo {
			break
		}
		i, j := rowFrom, colFrom
		for ; i < rowTo && j < colTo; j += 1 {
			if index == num {
				break
			}
			res[index] = matrix[i][j]
			index += 1
		}
		i, j = i+1, colTo-1
		for ; i < rowTo && j < colTo; i += 1 {
			if index == num {
				break
			}
			res[index] = matrix[i][j]
			index += 1
		}
		i, j = rowTo-1, j-1
		for ; i >= rowFrom && j >= colFrom; j -= 1 {
			if index == num {
				break
			}
			res[index] = matrix[i][j]
			index += 1
		}
		i, j = i-1, colFrom
		for ; i > rowFrom && j >= colFrom; i -= 1 {
			if index == num {
				break
			}
			res[index] = matrix[i][j]
			index += 1
		}
		cnt += 1
	}
	return res
}
