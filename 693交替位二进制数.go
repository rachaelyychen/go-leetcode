package main

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(5))
}

func hasAlternatingBits(n int) bool {
	var arr []int
	for n != 0 {
		arr = append(arr, n%2)
		n /= 2
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
	}
	return true
}
