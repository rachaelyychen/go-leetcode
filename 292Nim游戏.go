package main

func canWinNim(n int) bool {
	if n < 4 {
		return true
	}
	if n > 4 {
		if (n-1)%4==0 || (n-2)%4==0 || (n-3)%4== 0 {
			return true
		}
	}
	return false
}
