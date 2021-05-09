package main

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	arr1, arr2, t1, t2 := []byte(s1), []byte(s2), -1, -1
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			if t1 < 0 {
				t1 = i
			} else if t2 < 0 {
				t2 = i
			} else {
				return false
			}
		}
	}
	if t1 < 0 && t2 < 0 || t1 >= 0 && t2 >= 0 && arr1[t1] == arr2[t2] && arr1[t2] == arr2[t1] {
		return true
	}
	return false
}
