package main

func rotateString(A string, B string) bool {
	if A == B {
		return true
	}
	if len(A) == 0 && len(B) != 0 || len(A) != 0 && len(B) == 0 {
		return false
	}
	arr := []byte(A)
	cnt := 0
	for cnt < len(A) {
		arr = append(arr[1:], arr[0])
		if string(arr) == B {
			return true
		}
		cnt += 1
	}
	return false
}


