package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < n; i++ {
		var ok bool
		for j := 0; j < len(flowerbed); j++ {
			if flowerbed[j] == 1 {
				continue
			}
			if j == 0 && (j+1 < len(flowerbed) && flowerbed[j+1] == 0 || len(flowerbed) == 1) {
				flowerbed[j] = 1
				ok = true
				break
			}
			if j == len(flowerbed)-1 && j-1 >= 0 && flowerbed[j-1] == 0 {
				flowerbed[j] = 1
				ok = true
				break
			}
			if j-1 >= 0 && j+1 < len(flowerbed) && flowerbed[j-1] == 0 && flowerbed[j+1] == 0 {
				flowerbed[j] = 1
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
