package main

func minOperations(s string) int {
	var cnt1, cnt2 int
	for i, b := range []byte(s) {
		if b != byte(i%2)+'0' {
			cnt1 += 1
		} else {
			cnt2 += 1
		}
	}
	if cnt1 < cnt2 {
		return cnt1
	}
	return cnt2
}
