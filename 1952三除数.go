package main

import "math"

func isThree(n int) bool {
	t1 := int(math.Sqrt(float64(n)))
	var ok bool
	for i := 2; i <= t1; i++ {
		if n%i != 0 {
			continue
		}
		t2 := n / i
		if t1 != t2 {
			return false
		}
		ok = true
	}
	return ok
}
