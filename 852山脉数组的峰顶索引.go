package main

func peakIndexInMountainArray(arr []int) int {
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] < arr[i+1] {
			i += 1
		}
		if arr[j] < arr[j-1] {
			j -= 1
		}
	}
	return i
}


