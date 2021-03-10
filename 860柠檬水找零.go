package main

func lemonadeChange(bills []int) bool {
	// 5不用找，10要找1*5，20要找3*5或1*10+1*5
	cnt5, cnt10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			cnt5 += 1
			continue
		}
		if cnt5 == 0 {
			return false
		}
		if bill == 10 {
			cnt5 -= 1
			cnt10 += 1
		}
		if bill == 20 {
			if cnt5 < 3 && cnt10 == 0 {
				return false
			}
			if cnt10 > 0 {
				cnt5 -= 1
				cnt10 -= 1
				continue
			}
			if cnt5 >= 3 {
				cnt5 -= 3
			}
		}
	}
	return true
}


