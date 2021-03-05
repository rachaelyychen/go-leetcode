package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	s2t, t2s := make(map[rune]rune, 0), make(map[rune]rune, 0)
	rs, rt := []rune(s), []rune(t)
	var res []rune
	for i:= 0; i < len(rs); i++ {
		if v, ok := s2t[rs[i]]; !ok {
			if _, ok1 := t2s[rt[i]]; !ok1 {
				s2t[rs[i]] = rt[i]
				t2s[rt[i]] = rs[i]
				res = append(res, rt[i])
			} else {
				return false
			}
		} else {
			res = append(res, v)
		}
	}
	return strings.EqualFold(string(res), t)
}
