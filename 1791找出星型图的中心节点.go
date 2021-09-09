package main

func findCenter(edges [][]int) int {
	arr := make([]int, 1e5+1)
	for i := range edges {
		arr[edges[i][0]] += 1
		if arr[edges[i][0]] == len(edges) {
			return edges[i][0]
		}
		arr[edges[i][1]] += 1
		if arr[edges[i][1]] == len(edges) {
			return edges[i][1]
		}
	}
	return 0
}
