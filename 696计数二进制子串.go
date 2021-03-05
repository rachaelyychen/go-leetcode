package main

func countBinarySubstrings(s string) int {
	// 先统计连续的0和1分别有多少个，如：111100011000，得到4323；在4323中的任意相邻两个数字，取小的一个加起来，就是3+2+2 = 7.
	var arr []int
	b := []byte(s)
	var t int
	for i := 0; i < len(b); i++ {
		t += 1
		if i == len(b)-1 || b[i] != b[i+1] {
			arr = append(arr, t)
			t = 0
		}
	}
	var res int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			res += arr[i]
		} else {
			res += arr[i+1]
		}
	}
	return res
}
