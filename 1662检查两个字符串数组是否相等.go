package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var str1, str2 string
	for i := range word1 {
		str1 += word1[i]
	}
	for i := range word2 {
		str2 += word2[i]
	}
	return str1 == str2
}
