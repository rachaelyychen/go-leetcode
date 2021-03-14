package main

func commonChars(A []string) []string {
	arr := make([]int, 30)
	for i := range A {
		t := make([]int, 30)
		for _, b := range []byte(A[i]) {
			t[b-'a'] += 1
		}
		if i == 0 {
			for j := range t {
				arr[j] = t[j]
			}
		} else {
			for j := range arr {
				if arr[j] > 0 {
					if t[j] == 0 {
						arr[j] = 0
					} else if t[j] < arr[j] {
						arr[j] = t[j]
					}
				}
			}
		}
	}
	var res []string
	for i := range arr {
		for arr[i] > 0 {
			res = append(res, string(i+'a'))
			arr[i] -= 1
		}
	}
	return res
}
