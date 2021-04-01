package main

func maxScore(s string) int {
	// 左 子字符串中 0 的数量加上 右 子字符串中 1 的数量
	arr := []byte(s)
	var cnt0, cnt1, tot1, res int
	for i := range arr {
		if arr[i] == '1' {
			tot1 += 1
		}
	}
	for i := range arr {
		if i > 0 {
			if arr[i-1] == '0' {
				cnt0 += 1
			} else {
				cnt1 += 1
			}
			if cnt0+tot1-cnt1 > res {
				res = cnt0 + tot1 - cnt1
			}
		}
	}
	return res
}
