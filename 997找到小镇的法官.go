package main

func findJudge(N int, trust [][]int) int {
	arrTrust, arrBeTrusted := make([]int, N+1), make([]int, N+1)
	for i := range trust {
		arrTrust[trust[i][0]] += 1
		arrBeTrusted[trust[i][1]] += 1
	}
	var res []int
	for i := range arrTrust {
		if i > 0 && arrTrust[i] == 0 && arrBeTrusted[i] == N-1 {
			res = append(res, i)
		}
	}
	if len(res) == 1 {
		return res[0]
	}
	return -1
}


