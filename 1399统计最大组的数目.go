package main

func countLargestGroup(n int) int {
	arr := make([]int, 10005)
	for i := 1; i <= n; i++ {
		t, a := i, 0
		for t != 0 {
			a += t%10
			t /= 10
		}
		arr[i] = a
	}
	maxCnt := 0
	m := make(map[int]int, 0)
	for i := range arr {
		if i > 0 && arr[i] == 0 {
			break
		}
		if i > 0 {
			if cnt, exists := m[arr[i]]; !exists {
				m[arr[i]] = 1
			} else {
				m[arr[i]] = cnt+1
			}
			if m[arr[i]] > maxCnt {
				maxCnt = m[arr[i]]
			}
		}
	}
	var res int
	for _, cnt := range m {
		if cnt == maxCnt {
			res += 1
		}
	}
	return res
}


