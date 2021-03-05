package main

func findLUSlength(a string, b string) int {
	lenA, lenB := len(a), len(b)
	if lenA > lenB {
		return lenA
	} else if lenA < lenB {
		return lenB
	} else if a != b {
		return lenA
	}
	return -1
}
