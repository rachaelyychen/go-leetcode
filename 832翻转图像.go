package main

func flipAndInvertImage(image [][]int) [][]int {
	for i := range image {
		var t []int
		for j := len(image[i])-1; j >= 0; j-- {
			if image[i][j] == 0 {
				t = append(t, 1)
			} else {
				t = append(t, 0)
			}
		}
		image[i] = t
	}
	return image
}


