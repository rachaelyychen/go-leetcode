package main

func findSpecialInteger(arr []int) int {
	for i := range arr {
		cnt := 1
		for j := i+1; j < len(arr); j++ {
			if arr[j] == arr[i] {
				cnt += 1
			} else {
				i = j-1
				break
			}
		}
		if 4*cnt > len(arr) {
			return arr[i]
		}
	}
	return 0
}


