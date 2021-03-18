package main

func maxNumberOfBalloons(text string) int {
	arr := make([]int, 30)
	for _, b := range []byte(text) {
		arr[b-'a'] += 1
	}
	aCnt, bCnt, lCnt, oCnt, nCnt := 1, 1, 2, 2, 1
	var res int
	if arr[0] == 0 || arr[0] < aCnt {
		return 0
	} else {
		res = arr[0] / aCnt
	}
	if arr[1] == 0 || arr[1] < bCnt {
		return 0
	} else if res > arr[1] / bCnt {
		res = arr[1] / bCnt
	}
	if arr['l'-'a'] == 0 || arr['l'-'a'] < lCnt {
		return 0
	} else if res > arr['l'-'a'] / lCnt {
		res = arr['l'-'a'] / lCnt
	}
	if arr['o'-'a'] == 0 || arr['o'-'a'] < oCnt {
		return 0
	} else if res > arr['o'-'a'] / oCnt {
		res = arr['o'-'a'] / oCnt
	}
	if arr['n'-'a'] == 0 || arr['n'-'a'] < nCnt {
		return 0
	} else if res > arr['n'-'a'] / nCnt {
		res = arr['n'-'a'] / nCnt
	}
	return res
}


