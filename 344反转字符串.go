package main

func reverseString(s []byte)  {
	i, j := 0, len(s)- 1
	for i < j {
		t := s[i]
		s[i] = s[j]
		s[j] = t
		i += 1
		j -= 1
	}
	return
}
