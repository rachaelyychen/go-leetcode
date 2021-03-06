package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(longestWord([]string{"yo","ew","fc","zrc","yodn","fcm","qm","qmo","fcmz","z","ewq","yod","ewqz","y"}))
}

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) < len(words[j])
		} else {
			return strings.Compare(words[i], words[j]) > 0
		}
	})
	m := make(map[int][]string, 0)
	for i := range words {
		if v, exists := m[len(words[i])]; !exists {
			m[len(words[i])] = []string{words[i]}
		} else {
			m[len(words[i])] = append(v, words[i])
		}
	}
	// fmt.Println(m)
	var res string
	for i := len(words)-1; i >= 0; i-- {
		ok := true
		for j := 1; j < len(words[i]); j++ {
			if v, exists := m[j]; !exists {
				ok = false
				break
			} else {
				ok2 := false
				for k := range v {
					if strings.HasPrefix(words[i], v[k]) {
						ok2 = true
						break
					}
				}
				if ok2 == false {
					ok = false
					break
				}
			}
		}
		if ok {
			res = words[i]
			break
		}
	}
	return res
}
