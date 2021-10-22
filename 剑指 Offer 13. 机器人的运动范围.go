package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/20/21 11:29 AM
**/

// 遍历数组+回溯法，以每个方格为起始位置，得到机器人能够到达的最多方格数。需要使用visited数组标记已访问过的方格。

func movingCount(m int, n int, k int) int {
	var res int
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var t int
			moving(m, n, k, i, j, visited, &t)
			if t > res {
				res = t
			}
		}
	}
	return res
}

func moving(m, n, k, i, j int, visited [][]bool, cnt *int) {
	if !checkMoving(m, n, k, i, j, visited) {
		return
	}
	visited[i][j] = true
	*cnt += 1
	moving(m, n, k, i-1, j, visited, cnt)
	moving(m, n, k, i+1, j, visited, cnt)
	moving(m, n, k, i, j-1, visited, cnt)
	moving(m, n, k, i, j+1, visited, cnt)
}

func checkMoving(m, n, k, i, j int, visited [][]bool) bool {
	if i < 0 || i == m || j < 0 || j == n || visited[i][j] {
		return false
	}
	var t int
	for i > 0 {
		t += i % 10
		i /= 10
	}
	for j > 0 {
		t += j % 10
		j /= 10
	}
	if t > k {
		return false
	}
	return true
}
