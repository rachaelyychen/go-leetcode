package main

func findNthDigit(n int) int {
	cnt, i := 0, 0
	for i = 1; i <= n; i++ {
		t := i * getDigitCount(i)
		if cnt+t > n {
			break
		}
		cnt += t
	}
	if i == 1 {
		return n
	}
	min := 1
	for j := 0; j < i-1; j++ {
		min *= 10
	}
	num := (n-cnt)/i + min
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	return arr[len(arr)-1-(n-cnt)%i]
}

func getDigitCount(n int) int {
	min, max := 1, 0
	for i := 0; i < n-1; i++ {
		min *= 10
	}
	if n == 1 {
		min = 0
	}
	for i := 0; i < n; i++ {
		max = max*10 + 9
	}
	return max - min + 1
}
