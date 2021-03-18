package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)
	var ans1, ans2 int
	for i := start; i != destination; i = (i+1)%n {
		ans1 += distance[i]
	}
	for i := destination; i != start; i = (i+1)%n {
		ans2 += distance[i]
	}
	if ans1 < ans2 {
		return ans1
	}
	return ans2
}


