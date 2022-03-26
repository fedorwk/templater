package basic_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fedorwk/templater/basic"
)

func TestReplaceString(t *testing.T) {
	template := "{name} is {job}"
	rows := []map[string]string{
		{
			"{name}": "Jack",
			"{job}":  "Sales",
		},
		{
			"{name}": "Sam",
			"{job}":  "Admin",
		},
	}
	for _, row := range rows {
		fmt.Println("Processing row: ", row)
		basic.ReplaceString(template, row, os.Stdout)
		fmt.Fprintln(os.Stdout)
	}
}
