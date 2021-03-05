package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	var primeArr []int
	arr := make([]int, n)
	for i := 2; i < n; i++ {
		if arr[i] == 0 {
			primeArr = append(primeArr, i)
		}
		for j := i*2; j < n; j+=i {
			arr[j] = 1
		}
	}
	return len(primeArr)
}
