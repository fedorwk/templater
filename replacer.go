package templater

import (
	"bytes"
	"io"
	"strings"

	"github.com/fedorwk/templater/basic"
)

type Replacer struct {
	items          []map[string]string
	startDelimiter string
	endDelimiter   string
}

func NewReplacer(items []map[string]string, stratDelimiter, endDelimiter string) *Replacer {
	return &Replacer{
		items:          items,
		startDelimiter: stratDelimiter,
		endDelimiter:   endDelimiter,
	}
}

func (repl *Replacer) Len() int {
	return len(repl.items)
}

func (repl *Replacer) replacementSliceAt(n int) []string {
	if n < 0 || n > len(repl.items) {
		return nil
	}
	if !repl.withDelimiters() {
		return basic.UnwrapMap(repl.items[n])
	}
	res := make([]string, 0, len(repl.items[n])*2)
	for key, val := range repl.items[n] {
		res = append(
			res,
			strings.Join([]string{repl.startDelimiter, key, repl.endDelimiter}, ""),
			val,
		)
	}
	return res
}

func (repl *Replacer) executeToString(template string, n int) string {
	var buf bytes.Buffer
	_, err := repl.executeToStream(template, n, &buf)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (repl *Replacer) executeToStream(template string, n int, w io.Writer) (int, error) {
	replacer := strings.NewReplacer(repl.replacementSliceAt(n)...)
	bytesWritten, err := replacer.WriteString(w, template)
	if err != nil {
		return bytesWritten, err
	}
	return bytesWritten, nil
}

func (repl *Replacer) withDelimiters() bool {
	return repl.startDelimiter != "" || repl.endDelimiter != ""
}
