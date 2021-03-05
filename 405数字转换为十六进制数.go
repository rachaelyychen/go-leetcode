package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var res string
	hexArr := []byte("0123456789abcdef")
	for num != 0 && len(res) < 8 {
		res = string(hexArr[num & 0xf]) + res
		num >>= 4
	}
	return res
}
