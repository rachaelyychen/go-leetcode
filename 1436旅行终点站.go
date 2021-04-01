package main

func destCity(paths [][]string) string {
	m := make(map[string]bool, 0)
	for i := range paths {
		fromCity, toCity := paths[i][0], paths[i][1]
		m[fromCity] = false
		if _, exists := m[toCity]; !exists {
			m[toCity] = true
		}
	}
	for city, isDest := range m {
		if isDest {
			return city
		}
	}
	return ""
}


