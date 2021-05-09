package main

func mergeAlternately(word1 string, word2 string) string {
	var res []byte
	arr1, arr2, i, j := []byte(word1), []byte(word2), 0, 0
	for i < len(arr1) || j < len(arr2) {
		if i == len(arr1) {
			res = append(res, arr2[j:]...)
			break
		}
		if j == len(arr2) {
			res = append(res, arr1[i:]...)
			break
		}
		res = append(res, arr1[i])
		i += 1
		res = append(res, arr2[j])
		j += 1
	}
	return string(res)
}
