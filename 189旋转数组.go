package main

func rotate(nums []int, k int) {
	// arr := make([]int, len(nums))
	// for i := range nums {
	// 	arr[(i+k)%len(nums)] = nums[i]
	// }
	// for i := range arr {
	// 	nums[i] = arr[i]
	// }

	var cnt int
	for start := 0; start < len(nums); start++ {
		si, sn, ok := start, nums[start], false
		for {
			if si == start && ok {
				break
			}
			ti, tn := (si+k)%len(nums), nums[(si+k)%len(nums)]
			nums[ti] = sn
			sn = tn
			si = ti
			cnt += 1
			if cnt == len(nums) {
				return
			}
			if !ok {
				ok = true
			}
		}
	}
}
