package main

import "strconv"

func freqAlphabets(s string) string {
	// 字符（'a' - 'i'）分别用（'1' - '9'）表示。
	// 字符（'j' - 'z'）分别用（'10#' - '26#'）表示。
	var res string
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		var t int
		if (arr[i] == '1' || arr[i] == '2') && i+2 < len(arr) && arr[i+2] == '#' {
			x, _ := strconv.Atoi(string(arr[i]))
			y, _ := strconv.Atoi(string(arr[i+1]))
			t = x*10 + y
			i = i + 2
		} else {
			t, _ = strconv.Atoi(string(arr[i]))
		}
		res += string(t - 1 + 'a')
	}
	return res
}
