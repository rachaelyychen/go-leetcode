package main

func longestPalindrome(s string) string {
	arr, dp := []byte(s), [1005][1005]int{}
	var res string
	var ans int
	for i := 0; i < len(arr); i++ {
		dp[i][i] = 1
		if ans == 0 {
			res = string(arr[i])
			ans = 1
		}
		if i+1 < len(arr) && arr[i] == arr[i+1] {
			dp[i][i+1] = 1
			if ans < 2 {
				res = string(arr[i : i+2])
				ans = 2
			}
		}
	}
	for l := 3; l <= len(arr); l++ {
		for i := 0; i+l-1 < len(arr); i++ {
			j := i + l - 1
			if dp[i+1][j-1] == 1 && arr[i] == arr[j] {
				dp[i][j] = 1
				if ans < l {
					res = string(arr[i : i+l])
					ans = l
				}
			}
		}
	}
	return res
}
