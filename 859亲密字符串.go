package main

func buddyStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		return hasDuplicateChar(a)
	}
	diffCnt, diff1, diff2 := 0, -1, -1
	arrA, arrB := []byte(a), []byte(b)
	for i := 0; i < len(arrA); i++ {
		if arrA[i] != arrB[i] {
			diffCnt += 1
			if diffCnt > 2 {
				return false
			}
			if diffCnt == 1 {
				diff1 = i
			}
			if diffCnt == 2 {
				diff2 = i
			}
		}
	}
	if diffCnt == 2 && arrA[diff1] == arrB[diff2] && arrA[diff2] == arrB[diff1] {
		return true
	}
	return false
}

func hasDuplicateChar(a string) bool {
	arr := make([]int, 30)
	for _, b := range []byte(a) {
		if arr[b-'a'] == 1 {
			return true
		}
		arr[b-'a'] = 1
	}
	return false
}
