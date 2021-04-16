package main

func totalMoney(n int) int {
	week, res := n/7, 0
	for i := 1; i < week; i++ {
		res += i * 7
	}
	for i := week + 1; i <= week+n%7; i++ {
		res += i
	}
	return week*28 + res
}
