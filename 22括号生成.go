package main

func generateParenthesis(n int) []string {
	var res []string
	dfs(0, 0, "", 2*n, &res)
	return res
}

func dfs(left, right int, t string, l int, res *[]string) {
	if len(t) == l {
		if left == right {
			*res = append(*res, t)
		}
		return
	}
	if left < right {
		return
	}
	dfs(left+1, right, t+"(", l, res)
	dfs(left, right+1, t+")", l, res)
}
