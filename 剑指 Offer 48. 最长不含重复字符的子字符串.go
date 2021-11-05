package main

// dp[i]表示以s[i]结尾的最长无重复字符的子串长度，s[j]表示距s[i]最近的左边相同字符，动态方程：dp[i] = dp[i-1]+1, dp[i-1]<i-j  i-j, dp[i-1]>=i-j

func lengthOfLongestSubstring(s string) int {
	dp := make([]int, len(s))
	m := make(map[byte]int, 0)
	for i := range s {
		j := -1
		if t, exists := m[s[i]]; exists {
			j = t
		}
		m[s[i]] = i
		if i == 0 {
			dp[i] = 1
		} else {
			if dp[i-1] < i-j {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = i - j
			}
		}
	}
	var res int
	for i := range dp {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
