package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
}

func licenseKeyFormatting(S string, K int) string {
	S = strings.ReplaceAll(strings.ToUpper(S), "-", "")
	arr := []byte(S)
	var res string
	for i := len(arr) - 1; i >= 0; i -= K {
		if i-K+1 > 0 {
			res = "-" + string(arr[i-K+1:i+1]) + res
		} else if i-K+1 == 0 {
			res = string(arr[i-K+1:i+1]) + res
		} else {
			res = string(arr[0:i+1]) + res
		}
	}
	return res
}
