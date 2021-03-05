package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(constructRectangle(5))
}

func constructRectangle(area int) []int {
	mid := int(math.Floor(math.Sqrt(float64(area))))
	low := mid
	for low > 0 {
		if area%low == 0 {
			high := area/low
			if high >= low {
				return []int{high, low}
			}
		}
		low -= 1
	}
	return nil
}
