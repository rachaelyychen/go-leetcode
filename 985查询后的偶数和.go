package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var res []int
	for i := range queries {
		var sum int
		if i == 0 {
			for j := range A {
				if A[j]%2 == 0 {
					sum += A[j]
				}
			}
		} else {
			sum = res[len(res)-1]
		}
		if A[queries[i][1]]%2 == 0 {
			sum -= A[queries[i][1]]
		}
		A[queries[i][1]] += queries[i][0]
		if A[queries[i][1]]%2 == 0 {
			sum += A[queries[i][1]]
		}
		res = append(res, sum)
	}
	return res
}



