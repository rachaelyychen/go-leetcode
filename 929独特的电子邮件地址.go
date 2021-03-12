package main

import "strings"

func numUniqueEmails(emails []string) int {
	emailMap := make(map[string]bool, 0)
	for _, email := range emails {
		arr := strings.Split(email, "@")
		if len(arr) == 2 {
			email = arr[0]
			if strings.Contains(email, "+") {
				email = email[:strings.Index(email, "+")]
			}
			email = strings.ReplaceAll(email, ".", "")
			email += "@" + arr[1]
			if _, exists := emailMap[email]; !exists {
				emailMap[email] = true
			}
		}
	}
	return len(emailMap)
}


