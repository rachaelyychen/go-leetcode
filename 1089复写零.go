package main

func duplicateZeros(arr []int)  {
	var t []int
	for i := 0; i < len(arr); i++ {
		t = append(t, arr[i])
	}
	arr = arr[:0]
	for i := 0; i < len(t); i++ {
		arr = append(arr, t[i])
		if len(arr) == len(t) {
			return
		}
		if t[i] == 0 {
			arr = append(arr, 0)
		}
	}
}


