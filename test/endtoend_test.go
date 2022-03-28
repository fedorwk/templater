package test

import (
	"fmt"
	"testing"

	"github.com/fedorwk/templater"
)

func TestBasicReplacer(t *testing.T) {
	template := BasicTemplate
	items := BasicItems
	multireplacer := templater.NewMultiReplacer(items, "{", "}")
	templater := templater.NewTemplater(template, multireplacer)
	fmt.Println(templater.AllStrings())
}

func TestEmailReplacer(t *testing.T) {
	template := GetEmailTemplate()
	items := EmailItems
	replacer := templater.NewMultiReplacer(items, "{", "}")
	templater := templater.NewTemplater(template, replacer)
	fmt.Println(templater.AllStrings())
}
