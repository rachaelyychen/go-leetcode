package main

func intToRoman(num int) string {
	valueSymbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var res string
	for num > 0 {
		for i := range valueSymbols {
			if valueSymbols[i].value <= num {
				res += valueSymbols[i].symbol
				num -= valueSymbols[i].value
				break
			}
		}
	}
	return res
}
