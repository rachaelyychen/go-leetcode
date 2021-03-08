package main

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[byte]bool, 0)
	for _, b := range []byte(jewels) {
		jewelsMap[b] = true
	}
	var res int
	for _, b := range []byte(stones) {
		if _, exists := jewelsMap[b]; exists {
			res += 1
		}
	}
	return res
}


