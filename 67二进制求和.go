package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("1111", "1111"))
}

func addBinary(a string, b string) string {
	var res string
	ra := []rune(a)
	rb := []rune(b)
	i := len(ra) - 1
	j := len(rb) - 1
	x := 0
	var na int
	var nb int
	for i >= 0 || j >= 0 {
		if i < 0 {
			na = 0
		} else {
			na, _ = strconv.Atoi(string(ra[i]))
		}
		if j < 0 {
			nb = 0
		} else {
			nb, _ = strconv.Atoi(string(rb[j]))
		}
		c := na + nb + x
		if c > 1 {
			c = c%2
			x = 1
		} else {
			x = 0
		}
		res += strconv.Itoa(c)
		i -= 1
		j -= 1
	}
	if x > 0 {
		res += strconv.Itoa(x)
	}
	return reverseString(res)
}

func reverseString(s string) string {
	r := []rune(s)
	var res string
	for i := len(r)-1; i >= 0; i-- {
		res += string(r[i])
	}
	return res
}
