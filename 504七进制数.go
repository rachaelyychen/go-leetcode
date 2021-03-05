package main

import "strconv"

func main() {

}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var arr []int
	var tag bool
	if num < 0 {
		num = -num
		tag = true
	}
	for num != 0 {
		arr = append(arr, num%7)
		num /= 7
	}
	var res string
	for i := len(arr) - 1; i >= 0; i-- {
		res += strconv.Itoa(arr[i])
	}
	if tag {
		res = "-" + res
	}
	return res
}
