package main

func numberOfLines(widths []int, s string) []int {
	maxWidth, minLines, curLineWidth := 100, 1, 0
	for _, b := range []byte(s) {
		if curLineWidth+widths[b-'a'] > maxWidth {
			minLines += 1
			curLineWidth = 0
		}
		curLineWidth += widths[b-'a']
	}
	return []int{minLines, curLineWidth}
}
