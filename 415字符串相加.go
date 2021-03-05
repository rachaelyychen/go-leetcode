package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
	arr1, arr2 := []byte(num1), []byte(num2)
	i, j, t := len(arr1)-1, len(arr2)-1, 0
	var arr3 []int
	for i >= 0 || j >= 0 || t > 0 {
		var n1, n2 int
		if i < 0 {
			n1 = 0
		} else {
			n1, _ = strconv.Atoi(string(arr1[i]))
		}
		if j < 0 {
			n2 = 0
		} else {
			n2, _ = strconv.Atoi(string(arr2[j]))
		}
		if n1+n2+t > 9 {
			arr3 = append(arr3, (n1+n2+t)%10)
			t = 1
		} else {
			arr3 = append(arr3, n1+n2+t)
			t = 0
		}
		i -= 1
		j -= 1
	}
	var res string
	for i := len(arr3)-1; i >= 0; i-- {
		res += strconv.Itoa(arr3[i])
	}
	return res
}
