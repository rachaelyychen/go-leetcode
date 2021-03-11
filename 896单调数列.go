package main

func isMonotonic(A []int) bool {
	var tag int
	for i := 0; i < len(A); i++ {
		if i == 0 && i+1 < len(A) {
			if A[i] < A[i+1] {
				tag = 1
			} else if A[i] > A[i+1] {
				tag = -1
			}
		}
		if i > 1 {
			if A[i-1] > A[i] {
				if tag == 1 {
					return false
				}
				if tag == 0 {
					tag = -1
				}
			}
			if A[i-1] < A[i] {
				if tag == -1 {
					return false
				}
				if tag == 0 {
					tag = 1
				}
			}
		}
	}
	return true
}


