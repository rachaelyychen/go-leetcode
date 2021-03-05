package main

import "fmt"

func main() {
	fmt.Println(reverse(-120))
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var res int
	for ; x !=0; {
		res = res*10 + x%10
		x /= 10
	}
	if res != int(int32(res)) {
		return 0
	}
	return res
}
