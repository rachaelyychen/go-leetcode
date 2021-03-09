package main

func largestTriangleArea(points [][]int) float64 {
	type point struct {
		x int
		y int
	}
	// （x1,y1）、（x2,y2）、（x3,y3） area=1/2|(x1y2+x2y3+x3y1-x1y3-x2y1-x3y2)|
	var res int
	for i := 0; i < len(points); i++ {
		point1 := point{x:points[i][0], y:points[i][1]}
		for j := 0; j < len(points); j++ {
			point2 := point{x:points[j][0], y:points[j][1]}
			if j != i {
				for k := 0; k < len(points); k++ {
					point3 := point{x:points[k][0], y:points[k][1]}
					if k != j && k != i {
						area := point1.x*point2.y + point2.x*point3.y + point3.x*point1.y - point1.x*point3.y - point2.x*point1.y - point3.x*point2.y
						if area < 0 {
							area = -area
						}
						if area > res {
							res = area
						}
					}
				}
			}
		}
	}
	return float64(res)/float64(2)
}


