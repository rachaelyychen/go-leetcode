package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcdefg", 3))
}

var arr []byte

func reverseStr(s string, k int) string {
	arr = []byte(s)
	var res string
	reverseStr541(0, k, &res)
	return res
}

func reverseStr541(from, k int, res *string) {
	if from == len(arr) || k == 0 {
		return
	}
	to := from + 2*k
	if to > len(arr) {
		to = len(arr)
	}
	var t []byte
	if to-from < k {
		// 全部反转
		for i := to-1; i >= from; i-- {
			t = append(t, arr[i])
		}
	}
	if to-from >= k {
		// 反转前k个
		for i := from+k-1; i >= from; i-- {
			t = append(t, arr[i])
		}
		t = append(t, arr[from+k:to]...)
	}
	*res += string(t)
	reverseStr541(to, k, res)
}
