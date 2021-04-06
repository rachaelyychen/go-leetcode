package main


func isPathCrossing(path string) bool {
	type PathPoint struct {
		x int
		y int
	}
	m := make(map[PathPoint]bool, 0)
	m[PathPoint{0,0}] = true
	var t PathPoint
	for _, b := range []byte(path) {
		switch b {
		case 'N':
			t.y += 1
		case 'S':
			t.y -= 1
		case 'E':
			t.x += 1
		case 'W':
			t.x -= 1
		}
		if _, exists := m[t]; exists {
			return true
		} else {
			m[t] = true
		}
	}
	return false
}


