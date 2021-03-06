package main

func isOneBitCharacter(bits []int) bool {
	var tag bool
	for i := 0; i < len(bits); i++ {
		if i == len(bits)-1 && tag == false {
			return true
		}
		if bits[i] == 1 && tag == false {
			tag = true
			continue
		}
		if tag == true {
			tag = false
		}
	}
	return false
}
