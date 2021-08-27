package main

func subarraysDivByK(nums []int, k int) int {
	arr, m := make([]int, len(nums)), make(map[int]int, 0)
	var res int
	// (x-y)%m = x%m - y%m = 0
	for i := range nums {
		if i == 0 {
			arr[i] = nums[i]
		} else {
			arr[i] = arr[i-1] + nums[i]
		}
		t := (arr[i]%k + k) % k
		if cnt, exists := m[t]; !exists {
			m[t] = 1
		} else {
			m[t] = cnt + 1
		}
	}
	for k, v := range m {
		if k == 0 {
			res += v
		}
		if v > 1 {
			res += v * (v - 1) / 2
		}
	}
	return res
}
