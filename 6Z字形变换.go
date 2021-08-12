package main

func convert(s string, numRows int) string {
	arr1, arr2, t := [1005]string{}, []byte(s), numRows*2-2
	if numRows == 1 {
		t = 1
	}
	for i := 0; i < len(arr2); i++ {
		if i%t < numRows {
			arr1[i%t] += string(arr2[i])
		} else {
			arr1[t-i%t] += string(arr2[i])
		}
	}
	var res string
	for i := 0; i < numRows; i++ {
		res += arr1[i]
	}
	return res
}
