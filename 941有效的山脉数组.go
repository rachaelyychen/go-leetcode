package main

func validMountainArray(arr []int) bool {
	i, j := 0, len(arr)-1
	for i+1 < len(arr) {
		if arr[i+1] > arr[i] {
			i += 1
		} else if arr[i+1] == arr[i] {
			return false
		} else {
			break
		}
	}
	if i == 0 {
		return false
	}
	for j-1 >= 0 {
		if arr[j-1] > arr[j] {
			j -= 1
		} else if arr[j-1] == arr[j] {
			return false
		} else {
			break
		}
	}
	if j == len(arr)-1 {
		return false
	}
	return i == j
}


