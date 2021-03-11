package main

func numSpecialEquivGroups(A []string) int {
	var res int
	arr := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if arr[i] > 0 {
			continue
		}
		for j := i+1; j < len(A); j++ {
			if arr[j] == 0 && isSpecialEquiv(A[i], A[j]) {
				arr[j] = 1
			}
		}
		arr[i] = 1
		res += 1
	}
	return res
}

func isSpecialEquiv(s1, s2 string) bool {
	arr1, arr2 := []byte(s1), []byte(s2)
	arr3, arr4 := make([]int, 30), make([]int, 30)
	for i := 0; i < len(arr1); i++ {
		if i % 2 == 1 {
			arr3[arr1[i]-'a'] += 1
			arr3[arr2[i]-'a'] -= 1
		} else {
			arr4[arr1[i]-'a'] += 1
			arr4[arr2[i]-'a'] -= 1
		}
	}
	for i := 0; i < 30; i++ {
		if arr3[i] != 0 || arr4[i] != 0 {
			return false
		}
	}
	return true
}


