package main

func countStudents(students []int, sandwiches []int) int {
	cnt1, cnt2 := 0, 0
	for i := range students {
		if students[i] == 0 {
			cnt1 += 1
		} else {
			cnt2 += 1
		}
	}
	for i := range sandwiches {
		if sandwiches[i] == 0 {
			if cnt1 > 0 {
				cnt1 -= 1
			} else {
				break
			}
		} else {
			if cnt2 > 0 {
				cnt2 -= 1
			} else {
				break
			}
		}
	}
	return cnt1 + cnt2
}
