package main

func reformat(s string) string {
	arr1, arr2 := make([]byte, 0), make([]byte, 0)
	for _, b := range []byte(s) {
		if b >= 'a' && b <= 'z' {
			arr1 = append(arr1, b)
		} else {
			arr2 = append(arr2, b)
		}
	}
	if len(arr1) != len(arr2) && len(arr1)+1 != len(arr2) && len(arr2)+1 != len(arr1) {
		return ""
	}
	var res string
	for i := 0; i < len(arr2); i++ {
		res += string(arr2[i])
		if i < len(arr1) {
			res += string(arr1[i])
		}
	}
	if len(arr2) < len(arr1) {
		res = string(arr1[len(arr1)-1]) + res
	}
	return res
}
