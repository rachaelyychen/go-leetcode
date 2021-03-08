package main

var m = map[int]bool{2:true, 3:true, 5:true, 7:true, 11:true, 13:true, 17:true, 19:true}

func countPrimeSetBits(L int, R int) int {
	var res int
	for i := L; i <= R; i++ {
		j, cnt := i, 0
		for j != 0 {
			if j%2 == 1 {
				cnt += 1
			}
			j /= 2
		}
		if isPrime(cnt) {
			res += 1
		}
	}
	return res
}

func isPrime(x int) bool {
	_, exists := m[x]
	return exists
}
