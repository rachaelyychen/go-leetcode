package main

func largestOddNumber(num string) string {
	b := []byte(num)
	for i := len(b) - 1; i >= 0; i-- {
		if (b[i]-'0')%2 == 1 {
			return string(b[:i+1])
		}
	}
	return ""
}
