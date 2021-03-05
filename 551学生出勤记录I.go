package main

func checkRecord(s string) bool {
	// 不超过一个A && 不超过两个连续的L
	return checkRecordA([]byte(s)) && checkRecordL([]byte(s))
}

func checkRecordA(arr []byte) bool {
	var aCnt int
	for i := range arr {
		if arr[i] == 'A' {
			aCnt += 1
			if aCnt == 2 {
				return false
			}
		}
	}
	return true
}

func checkRecordL(arr []byte) bool {
	var lCnt int
	for i := range arr {
		if arr[i] == 'L' {
			if i-1 >= 0 && arr[i-1]=='L' {
				lCnt += 1
				if lCnt > 2 {
					return false
				}
			} else {
				lCnt = 1
			}
		} else {
			lCnt = 0
		}
	}
	return true
}
