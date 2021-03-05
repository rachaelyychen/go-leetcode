package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	commonPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if strings.HasPrefix(strs[i], commonPrefix) {
			continue
		}
		arr := []rune(strs[i])
		var t string
		for j := 1; j <= len(arr); j++ {
			if strings.HasPrefix(commonPrefix, string(arr[:j])) {
				t = string(arr[:j])
			} else {
				break
			}
		}
		commonPrefix = t
	}
	return commonPrefix
}
