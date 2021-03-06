package main

import "fmt"

func main() {
	// 1 1 1
	// 1 1 0
	// 1 0 1
	fmt.Println(floodFill([][]int{{1,1,1}, {1,1,0}, {1,0,1}}, 1, 1, 2))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		floodFill2(&image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func floodFill2(image *[][]int, sr, sc, val, newColor int) {
	(*image)[sr][sc] = newColor
	if sr > 0 && (*image)[sr-1][sc] == val {
		floodFill2(image, sr-1, sc, val, newColor)
	}
	if sr < len(*image)-1 && (*image)[sr+1][sc] == val {
		floodFill2(image, sr+1, sc, val, newColor)
	}
	if sc > 0 && (*image)[sr][sc-1] == val {
		floodFill2(image, sr, sc-1, val, newColor)
	}
	if sc < len((*image)[0])-1 && (*image)[sr][sc+1] == val {
		floodFill2(image, sr, sc+1, val, newColor)
	}
}


