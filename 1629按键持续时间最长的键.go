package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	arr := make([]int, 30)
	var max int
	for i := range releaseTimes {
		var t int
		if i == 0 {
			t = releaseTimes[i]
		} else {
			t = releaseTimes[i] - releaseTimes[i-1]
		}
		if arr[keysPressed[i]-'a'] < t {
			arr[keysPressed[i]-'a'] = t
			if t > max {
				max = t
			}
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == max {
			return byte('a' + i)
		}
	}
	return 'a'
}
