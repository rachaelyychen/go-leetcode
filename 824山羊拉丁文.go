package main

import "strings"

func toGoatLatin(S string) string {
	var res string
	arr := strings.Split(S, " ")
	for i := range arr {
		if i > 0 {
			res += " "
		}
		if strings.HasPrefix(arr[i], "a") || strings.HasPrefix(arr[i], "A") ||
			strings.HasPrefix(arr[i], "e") || strings.HasPrefix(arr[i], "E") ||
			strings.HasPrefix(arr[i], "i") || strings.HasPrefix(arr[i], "I") ||
			strings.HasPrefix(arr[i], "o") || strings.HasPrefix(arr[i], "O") ||
			strings.HasPrefix(arr[i], "u") || strings.HasPrefix(arr[i], "U") {
			res = res + arr[i] + "ma"
		} else {
			b := []byte(arr[i])
			res = res + string(b[1:]) + string(b[0]) + "ma"
		}
		res += strings.Repeat("a", i+1)
	}
	return res
}


