package main

func nextGreatestLetter(letters []byte, target byte) byte {
	for i := range letters {
		if i == 0 && letters[i] > target {
			return letters[i]
		}
		if i+1 < len(letters) && letters[i] <= target && letters[i+1] > target {
			return letters[i+1]
		}
		if i == len(letters)-1 && letters[i] <= target {
			return letters[0]
		}
	}
	return letters[0]
}


