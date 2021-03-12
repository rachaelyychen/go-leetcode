package main

import "fmt"

func main() {
	fmt.Printf("%%%s", "a")
}

func reverseOnlyLetters(S string) string {
	arr := []byte(S)
	var s string
	for _, b := range arr {
		if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' {
			s += "%s"
		} else {
			if b == '%' {
				s += "%"
			}
			s += string(b)
		}
	}
	var reverseArr []interface{}
	for i := len(arr)-1; i >= 0; i-- {
		if arr[i] >= 'a' && arr[i] <= 'z' || arr[i] >= 'A' && arr[i] <= 'Z' {
			reverseArr = append(reverseArr, string(arr[i]))
		}
	}
	return fmt.Sprintf(s, reverseArr...)
}


