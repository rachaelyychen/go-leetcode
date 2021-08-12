package main

import "math"

func myAtoi(s string) int {
	// 去掉前导空格【描述1】
	i := 0
	for ; i < len(s); i ++ {
		if s[i] != ' ' {
			break
		}
	}
	s = s[i:]
	ans := 0
	sign := 1 // 默认为‘+’【描述2里的默认值】
	for i, v := range s {
		if v >= '0' && v <= '9' { // 转整数【描述4】
			ans = ans*10 + int(v-'0')
		} else if v == '-' && i == 0 { // 第一个字符(必须第一个)是‘+’或者‘-’【描述2】
			sign = -1
		} else if v == '+' && i == 0 { // 第一个字符(必须第一个)是‘+’或者‘-’【描述2】
			sign = 1
		} else { // 【描述3】
			break
		}

		// 边界处理【描述5】
		if sign == 1 && ans > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && -ans < math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * ans // 【描述6】
}
