package basic

import (
	"io"
	"strings"
)

func ReplaceString(template string, replacements map[string]string, dst io.Writer) (n int, err error) {
	replacementSlice := UnwrapMap(replacements)
	replacer := strings.NewReplacer(replacementSlice...)
	n, err = replacer.WriteString(dst, template)
	if err != nil {
		return n, err
	}
	return n, nil
}

func UnwrapMap(inputMap map[string]string) []string {
	res := make([]string, 0, len(inputMap)*2)
	for key, val := range inputMap {
		res = append(res, key, val)
	}
	return res
}
