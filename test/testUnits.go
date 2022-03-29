package test

import (
	"bytes"
	"io"
	"os"
)

var BasicTemplate = "Hi, {name}! You look {adjective}!"

var (
	BasicItems = []map[string]string{
		{"name": "Fedor", "adjective": "cool"},
		{"name": "Kate", "adjective": "smart"},
		{"name": "Sam", "adjective": "good"},
	}
	BasicItemsExpected = []string{
		"Hi, Fedor! You look cool!",
		"Hi, Kate! You look smart!",
		"Hi, Sam! You look good!",
	}

	EmailItems = []map[string]string{
		{"email": "exmpl@exmpl.com", "subject": "exmpl", "text": "some text"},
		{"email": "one@one.ru", "subject": "onesubjext", "text": "new\nline\ntest"},
		{"email": "EMAIL", "subject": "SUBJECT", "text": "TEXT"},
	}
)

func GetEmailTemplate() string {
	file, err := os.Open("draftemail.eml")
	if err != nil {
		return err.Error()
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	return buf.String()
}
