package test

import (
	"testing"

	"github.com/fedorwk/templater"
)

var (
	BasicTemplate    = "{objects} are {adjective}"
	OneSidedTemplate = "$objects are $adjective"
	BasicItems       = []map[string]string{
		{"objects": "Roses", "adjective": "red"},
		{"objects": "violets", "adjective": "blue"},
	}
	TextExpected = []string{
		"Roses are red",
		"violets are blue",
	}
)

func TestBasicTemplateReplacer(t *testing.T) {
	template := BasicTemplate
	items := BasicItems
	replacer := templater.NewReplacer(items, "{", "}")
	templater := templater.NewTemplater(template, replacer)
	result := templater.AllStrings() // Get all strings with replacement by item data
	for i := range result {
		if result[i] != TextExpected[i] {
			t.Errorf("unexpected result\nwant:%s, \ngot:%s", TextExpected[i], result[i])
		}
	}
}

func TestOneSidedTemplateReplacer(t *testing.T) {
	template := OneSidedTemplate
	items := BasicItems
	replacer := templater.NewReplacer(items, "$", "")
	templater := templater.NewTemplater(template, replacer)
	// Get strigns one by one with replacement by item data
	for i := range items {
		if result := templater.ExecuteToString(i); result != TextExpected[i] {
			t.Errorf(
				"unexpected result\nwant:%s, \ngot:%s",
				TextExpected[i],
				result)
		}
	}
}
