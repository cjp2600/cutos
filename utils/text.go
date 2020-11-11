package utils

import (
	"bytes"
	"regexp"
	"strings"
)

func PreviewContent(content string, num int) string {
	runes := bytes.Runes([]byte(content))
	if len(runes) > num {
		return string(runes[:num]) + " ..."
	}
	return string(runes)
}

func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func Clean(field, content string) string {
	switch strings.ToLower(field) {
	case "phone":
		runes := bytes.Runes([]byte(content))
		return string(runes[:6]) + "****"
	case "login":
		if isEmailValid(content) {
			return "example@example.com"
		}
		return "userLogin"
	case "email":
		return "example@example.com"
	case "password", "pwd", "pswrd":
		return "SECReEtPa$$WORD"
	default:
		return PreviewContent(content, 80)
	}
}
