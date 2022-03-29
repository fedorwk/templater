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
	result := templater.AllStrings()
	for i := range result {
		if result[i] != BasicItemsExpected[i] {
			t.Errorf("unexpected result\nwant:%s, \ngot:%s", BasicItemsExpected[i], result[i])
		}
	}
}

func TestEmailReplacer(t *testing.T) {
	template := GetEmailTemplate()
	items := EmailItems
	replacer := templater.NewMultiReplacer(items, "{", "}")
	templater := templater.NewTemplater(template, replacer)
	fmt.Println(templater.AllStrings())
}
