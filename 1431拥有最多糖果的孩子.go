package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxCandy int
	for i := range candies {
		if candies[i] > maxCandy {
			maxCandy = candies[i]
		}
	}
	var res []bool
	for i := range candies {
		if candies[i]+extraCandies >= maxCandy {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}


