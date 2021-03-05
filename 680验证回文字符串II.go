package main

func validPalindrome(s string) bool {
	arr := []byte(s)
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			return validPalindrome2(arr, i+1, j) || validPalindrome2(arr, i, j-1)
		}
		i += 1
		j -= 1
	}
	return true
}

func validPalindrome2(arr []byte, i, j int) bool {
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
