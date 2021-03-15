package main

func canThreePartsEqualSum(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	var sum int
	for i := range arr {
		sum += arr[i]
	}
	if sum % 3 != 0 {
		return false
	}
	i, j := 0, len(arr)-1
	sum1, sum2 := arr[i], arr[j]
	for i+1 < j {
		if sum1 == sum/3 && sum2 == sum/3 {
			return true
		}
		if sum1 != sum/3 {
			i += 1
			sum1 += arr[i]
		}
		if sum2 != sum/3 {
			j -= 1
			sum2 += arr[j]
		}
	}
	return false
}


