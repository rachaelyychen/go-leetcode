package main

func removeDuplicates(S string) string {
	arr := []byte(S)
	if !hasDuplicates(arr) {
		return S
	}
	for i := 0; i < len(arr); i++ {
		cnt, j := 1, 0
		for j = i+1; j <= len(arr); j++ {
			if j == len(arr) || arr[j] != arr[i] {
				break
			} else {
				cnt += 1
			}
		}
		if cnt % 2 == 0 {
			var t []byte
			if i > 0 {
				t = append(t, arr[:i]...)
			}
			if j < len(arr) {
				t = append(t, arr[j:]...)
			}
			arr = t
		}
	}
	return removeDuplicates(string(arr))
}

func hasDuplicates(arr []byte) bool {
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}


