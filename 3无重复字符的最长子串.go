package main

// 时间复杂度O(n)的解法:滑动窗口
// 以下标start表示滑动窗口的起始,遍历字符串s,用map记录每个字符的最新下标

func lengthOfLongestSubstring(s string) int {
	var res int
	start, m := 0, make(map[byte]int, 0)
	for i := range s {
		b := byte(s[i])
		if j, exists := m[b]; exists && start <= j {
			start = j+1
		}
		m[b] = i
		if i-start+1 > res {
			res = i-start+1
		}
	}
	return res
}
