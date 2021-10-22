package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/20/21 9:38 AM
**/

// 遍历数组+回溯法，以每个元素为起始字符，检查是否存在连续字符满足条件。由于一个字符不能使用两次，还需要visited数组标记已使用的字符。

func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 {
		return false
	}
	col := len(board[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if checkExist(board, []byte(word), i, j, 0, row, col, visited) {
				return true
			}
		}
	}
	return false
}

func checkExist(board [][]byte, word []byte, i, j, k, row, col int, visited [][]bool) bool {
	if i >= 0 && i < row && j >= 0 && j < col && visited[i][j] == false {
		if board[i][j] == word[k] {
			visited[i][j] = true
			if k == len(word)-1 {
				return true
			}
			if !(checkExist(board, word, i-1, j, k+1, row, col, visited) ||
				checkExist(board, word, i+1, j, k+1, row, col, visited) ||
				checkExist(board, word, i, j-1, k+1, row, col, visited) ||
				checkExist(board, word, i, j+1, k+1, row, col, visited)) {
				visited[i][j] = false
			} else {
				return true
			}
		}
	}
	return false
}
