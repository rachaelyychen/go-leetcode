package main

func divisorGame(N int) bool {
	// a拿到1，b输了；a拿到2，b输了；a拿到3，b赢了
	if N == 2 {
		return true
	}
	if N ==1 || N == 3 {
		return false
	}
	for x := 1; x < N && N%x == 0; x++ {
		if divisorGame(N-x) == false {
			return true
		}
	}
	return false
}
