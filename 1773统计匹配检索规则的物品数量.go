package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var res int
	for i := range items {
		switch ruleKey {
		case "type":
			if items[i][0] == ruleValue {
				res += 1
			}
		case "color":
			if items[i][1] == ruleValue {
				res += 1
			}
		case "name":
			if items[i][2] == ruleValue {
				res += 1
			}
		}
	}
	return res
}
