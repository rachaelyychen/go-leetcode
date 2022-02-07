package main

import (
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		len1, len2 := len(dictionary[i]), len(dictionary[j])
		if len1 != len2 {
			return len1 > len2
		}
		return dictionary[i] < dictionary[j]
	})
	for i := range dictionary {
		if len(dictionary[i]) > len(s) {
			continue
		}
		j, k := 0, 0
		for j < len(dictionary[i]) && k < len(s) {
			for k < len(s) && s[k] != dictionary[i][j] {
				k += 1
			}
			for j < len(dictionary[i]) && k < len(s) && s[k] == dictionary[i][j] {
				k += 1
				j += 1
			}
		}
		if j == len(dictionary[i]) {
			return dictionary[i]
		}
	}
	return ""
}
