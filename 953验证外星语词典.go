package main

func isAlienSorted(words []string, order string) bool {
	sortArr := make([]int, 30)
	for i := range order {
		sortArr[order[i]-'a'] = i
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]) || j < len(words[i+1]); j++ {
			if j == len(words[i+1]) {
				return false
			}
			if j == len(words[i]) {
				break
			}
			if words[i][j] != words[i+1][j] {
				if sortArr[words[i][j]-'a'] > sortArr[words[i+1][j]-'a'] {
					return false
				} else if sortArr[words[i][j]-'a'] < sortArr[words[i+1][j]-'a'] {
					break
				}
			}
		}
	}
	return true
}
