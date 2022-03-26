package basic

import (
	"io"
	"strings"

	"github.com/fedorwk/templater/util"
)

func ReplaceString(template string, replacements map[string]string, dst io.Writer) (n int, err error) {
	replacementSlice := util.UnwrapMap(replacements)
	replacer := strings.NewReplacer(replacementSlice...)
	n, err = replacer.WriteString(dst, template)
	if err != nil {
		return n, err
	}
	return n, nil
}
