package main

func uniqueMorseRepresentations(words []string) int {
	arr := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morseMap := make(map[string]int, 0)
	for _, word := range words {
		var morse string
		for _, b := range []byte(word) {
			morse += arr[b-'a']
		}
		if cnt, exists := morseMap[morse]; !exists {
			morseMap[morse] = 1
		} else {
			morseMap[morse] = cnt + 1
		}
	}
	return len(morseMap)
}
