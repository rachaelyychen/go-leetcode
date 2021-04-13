package main

func canFormArray(arr []int, pieces [][]int) bool {
	piecesMark := make([]int, len(pieces))
	for i := 0; i < len(arr); i++ {
		t := i
		hasFound := false
		for j := 0; j < len(pieces); j++ {
			if piecesMark[j] == 0 {
				ok := true
				for k := 0; k < len(pieces[j]); k++ {
					if pieces[j][k] != arr[t] {
						ok = false
						break
					} else {
						t += 1
					}
				}
				if ok {
					piecesMark[j] = 1
					hasFound = true
					break
				}
			}
		}
		if !hasFound {
			return false
		}
		i = t - 1
	}
	return true
}
