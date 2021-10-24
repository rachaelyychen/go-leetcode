package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/22/21 10:34 PM
**/

func isNumber(s string) bool {
	var (
		start, i int // start块的开始下标
		n        = len(s)
	)

	// 数值第一部分，(空格*x)
	for ; i < n && s[i] == ' '; i++ {
	}
	if i == n { // 只有空格
		return false
	}

	// 数值第二部分，(小数/整数)
	var pot bool // 有`.`就是小数
	start = i
	for ; i < n && (s[i] >= '0' && s[i] <= '9' || s[i] == '+' || s[i] == '-' || s[i] == '.'); i++ {
		if s[i] == '.' {
			pot = true
		}
	}
	var ans bool
	if pot {
		ans = checkDec(&s, start, i-1) // 检测是否为合法的小数
	} else {
		ans = checkNum(&s, start, i-1) // 检测是否为合法的整数
	}
	if !ans {
		return false
	}

	// 数值第三部分，[e/E整数]
	if i < n && (s[i] == 'E' || s[i] == 'e') {
		i++
		start = i
		for ; i < n && (s[i] >= '0' && s[i] <= '9' || s[i] == '+' || s[i] == '-'); i++ {
		}
		if !checkNum(&s, start, i-1) { // 检测e/E后面是否为合法的整数
			return false
		}
	}

	// 数值第四部分，(空格*x)
	for ; i < n && s[i] == ' '; i++ {
	}

	// i != n，不是合法的数值
	if i != n {
		return false
	}
	return true
}

func checkDec(s *string, start, end int) bool {
	var addSub, num, pot bool
	for i := start; i <= end; i++ {
		switch (*s)[i] {
		case '+', '-':
			if addSub || num || pot {
				return false // +/- 只能出现在第一个位置
			}
			addSub = true
		case '.':
			if pot {
				return false // 小数点只能出现一次
			}
			pot = true
		default:
			num = true
		}
	}
	return num // 必须要有数字
}

func checkNum(s *string, start, end int) bool {
	var addSub, num bool
	for i := start; i <= end; i++ {
		switch (*s)[i] {
		case '+', '-':
			if addSub || num {
				return false // +/- 只能出现在第一个位置
			}
			addSub = true
		default:
			num = true
		}
	}
	return num // 必须要有数字
}
