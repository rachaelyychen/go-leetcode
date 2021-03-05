package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	var res int
	var k int
	for i := 0; i < len(g); i++ {
		for j := k; j < len(s); j++ {
			if s[j] >= g[i] {
				k = j + 1
				res += 1
				break
			}
		}
	}
	return res
}
