package main

func subtractProductAndSum(n int) int {
	var arr []int
	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	t1, t2 := 1, 0
	for i := range arr {
		t1 *= arr[i]
		t2 += arr[i]
	}
	return t1-t2
}


