package main

func hasGroupsSizeX(deck []int) bool {
	deckCntMap := make(map[int]int, 0)
	minCnt := len(deck)
	for i := range deck {
		if cnt, exists := deckCntMap[deck[i]]; !exists {
			deckCntMap[deck[i]] = 1
		} else {
			deckCntMap[deck[i]] = cnt + 1
		}
	}
	for _, cnt := range deckCntMap {
		if cnt < minCnt {
			minCnt = cnt
		}
	}
	for size := 2; size <= minCnt; size++ {
		ok := true
		for _, cnt := range deckCntMap {
			if cnt%size > 0 {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
