package main

func restoreString(s string, indices []int) string {
	arr := []byte(s)
	res := make([]byte, len(arr))
	for i := 0; i < len(arr); i++ {
		res[indices[i]] = arr[i]
	}
	return string(res)
}


