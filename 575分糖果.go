package main

func distributeCandies(candyType []int) int {
	diffCandies, halfCandies := 0, len(candyType)/2
	candyTypeMap := make(map[int]int, 0)
	for i := range candyType {
		if _, exists := candyTypeMap[candyType[i]]; !exists {
			candyTypeMap[candyType[i]] = 1
		}
	}
	diffCandies = len(candyTypeMap)
	if diffCandies < halfCandies {
		return diffCandies
	} else {
		return halfCandies
	}
}
