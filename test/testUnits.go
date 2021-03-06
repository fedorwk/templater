package test

import (
	"bytes"
	"io"
	"os"
)

var (
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
