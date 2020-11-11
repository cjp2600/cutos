package utils

import (
	"bytes"
	"github.com/bxcodec/faker/v3"
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
		return faker.Phonenumber()
	case "productName":
		return faker.Name()
	case "url":
		return faker.URL()
	case "name":
		return faker.Name()
	case "lastname":
		return faker.LastName()
	case "firstname":
		return faker.FirstName()
	case "secondname":
		return faker.LastName()
	case "login":
		if isEmailValid(content) {
			return faker.Email()
		}
		return "userLogin"
	case "email":
		return faker.Email()
	case "password", "pwd", "pswrd":
		return faker.Password()
	default:
		return PreviewContent(content, 80)
	}
}
