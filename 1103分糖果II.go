package main

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	var round int
	for {
		for i := 1; i <= num_people; i++ {
			t := round*num_people + i
			if candies >= t {
				res[i-1] += t
				candies -= t
			} else {
				res[i-1] += candies
				candies = 0
			}
			if candies == 0 {
				return res
			}
		}
		round += 1
	}
}


