package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("", "a"))
}

func isSubsequence(s string, t string) bool {
	sArr := []byte(s)
	tArr := []byte(t)
	loc := 0
	for i := 0; i < len(sArr); i++ {
		var ok bool
		for j := loc; j < len(tArr); j++ {
			if sArr[i] == tArr[j] {
				loc = j+1
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
