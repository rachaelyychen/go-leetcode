package main

func findRestaurant(list1 []string, list2 []string) []string {
	stringIndexMap1 := make(map[string]int, 0)
	stringIndexMap2 := make(map[string]int, 0)
	for i := range list1 {
		stringIndexMap1[list1[i]] = i
	}
	for i := range list2 {
		if index, exists := stringIndexMap1[list2[i]]; exists {
			stringIndexMap2[list2[i]] = index + i
		}
	}
	var res []string
	min := len(list1) + len(list2)
	for s, index := range stringIndexMap2 {
		if index < min {
			min = index
			res = make([]string, 0)
			res = append(res, s)
		} else if index == min {
			res = append(res, s)
		}
	}
	return res
}
