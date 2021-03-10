package main

func backspaceCompare(S string, T string) bool {
	return backspaceCompare2(S) == backspaceCompare2(T)
}

func backspaceCompare2(s string) string {
	arr, res := []byte(s), make([]byte, 0)
	for i := range arr {
		if arr[i] == '#' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, arr[i])
		}
	}
	return string(res)
}


