package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	// "9001 discuss.leetcode.com"
	domainCntMap := make(map[string]int, 0)
	for _, cpdomain := range cpdomains {
		arr := strings.Split(cpdomain, " ")
		cnt, _ := strconv.Atoi(arr[0])
		domainArr := strings.Split(arr[1], ".")
		var s string
		for i := len(domainArr)-1; i >= 0; i-- {
			if i == len(domainArr)-1 {
				s = domainArr[i]
			} else {
				s = domainArr[i]+"."+s
			}
			if v, exists := domainCntMap[s]; !exists {
				domainCntMap[s] = cnt
			} else {
				domainCntMap[s] = v + cnt
			}
		}
	}
	var res []string
	for domain, cnt := range domainCntMap {
		res = append(res, strconv.Itoa(cnt)+" "+domain)
	}
	return res
}


