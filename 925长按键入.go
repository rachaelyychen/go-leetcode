package main

func isLongPressedName(name string, typed string) bool {
	if len(name) > len(typed) {
		return false
	}
	i, j := 0, 0
	for i < len(name) || j < len(typed) {
		if i == len(name) {
			if typed[j] != name[i-1] {
				return false
			}
			j += 1
			continue
		}
		if j == len(typed) {
			return false
		}
		if name[i] != typed[j] {
			if i == 0 || typed[j] != name[i-1] {
				return false
			}
			j += 1
			continue
		}
		i += 1
		j += 1
	}
	return true
}


