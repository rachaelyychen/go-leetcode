package main

func minOperations(logs []string) int {
	var res int
	for i := range logs {
		switch logs[i] {
		case "./":
		case "../":
			if res > 0 {
				res -= 1
			}
		default:
			res += 1
		}
	}
	return res
}
