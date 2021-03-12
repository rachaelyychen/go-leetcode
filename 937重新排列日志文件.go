package main

import (
	"sort"
	"strings"
	"unicode"
)

type logFile struct {
	t   string
	log string
}

func (l *logFile) isLet() bool {
	return unicode.IsLetter(rune(l.log[0]))
}

func reorderLogFiles(logs []string) []string {
	var letArr []logFile
	var digArr []string
	for _, log := range logs {
		i := strings.Index(log, " ")
		l := logFile{t:log[:i], log:log[i+1:]}
		if l.isLet() {
			letArr = append(letArr, l)
		} else {
			digArr = append(digArr, log)
		}
	}
	sort.Slice(letArr, func(i, j int) bool {
		if letArr[i].log != letArr[j].log {
			return letArr[i].log < letArr[j].log
		} else {
			return letArr[i].t < letArr[j].t
		}
	})
	var res []string
	for _, l := range letArr {
		res = append(res, l.t+" "+l.log)
	}
	res = append(res, digArr...)
	return res
}


