package main

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	arr := []byte(s)
	var res []byte
	res = append(res, arr[0], arr[1])
	for i := range arr {
		if i > 1 {
			if arr[i-1] == arr[i-2] && arr[i-1] == arr[i] {
				continue
			}
			res = append(res, arr[i])
		}
	}
	return string(res)
}
