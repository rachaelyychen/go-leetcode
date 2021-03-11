package main

func sortArrayByParity(A []int) []int {
	arr1, arr2 := make([]int, 5001), make([]int, 5002)
	for i := range A {
		if A[i] % 2 == 0 {
			arr1[A[i]] += 1
		} else {
			arr2[A[i]] += 1
		}
	}
	var res []int
	for i := range arr1 {
		for arr1[i] > 0 {
			res = append(res, i)
			arr1[i] -= 1
		}
	}
	for i := range arr2 {
		for arr2[i] > 0 {
			res = append(res, i)
			arr2[i] -= 1
		}
	}
	return res
}


