package main

func binaryGap(n int) int {
	maxDist, last1 := 0, -1
	var arr []int
	for n != 0 {
		arr = append(arr, n%2)
		n /= 2
		if arr[len(arr)-1] == 1{
			if last1 > -1 {
				dist := len(arr)-1-last1
				if dist > maxDist {
					maxDist = dist
				}
			}
			last1 = len(arr)-1
		}
	}
	return maxDist
}


