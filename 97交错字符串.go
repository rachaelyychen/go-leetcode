package main

// dp[i][j]表示s1取出的前i个元素和s2取出的前j个元素组成的字符串,是否和s3的i+j子串相等
// 边界条件:dp[0][0]=true,dp[i][0],dp[0][i]

func isInterleave(s1 string, s2 string, s3 string) bool {
	len1, len2, len3 := len(s1), len(s2), len(s3)
	if len1+len2 != len3 {
		return false
	}
	if len1 == 0 {
		return s2 == s3
	}
	if len2 == 0 {
		return s1 == s3
	}
	var dp [][]bool
	for i := 0; i <= len1; i++ {
		dp = append(dp, make([]bool, len2+1))
	}
	dp[0][0] = true
	if s1[0] == s3[0] {
		dp[1][0] = true
	}
	if s2[0] == s3[0] {
		dp[0][1] = true
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if dp[i-1][j-1] && (s1[i-1] == s3[i+j-2] && s2[j-1] == s3[i+j-1] || s1[i-1] == s3[i+j-1] && s2[j-1] == s3[i+j-2]) ||
				dp[i][j-1] && s2[j-1] == s3[i+j-1] || dp[i-1][j] && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[len1][len2]
}
