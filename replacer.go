package templater

import (
	"bytes"
	"io"
	"strings"

	"github.com/fedorwk/templater/util"
)

type MultiReplacer struct {
	replacements   []map[string]string
	startDelimiter string
	endDelimiter   string
}

func NewMultiReplacer(replacements []map[string]string, stratDelim, endDelim string) *MultiReplacer {
	return &MultiReplacer{
		replacements:   replacements,
		startDelimiter: stratDelim,
		endDelimiter:   endDelim,
	}
}

func (repl *MultiReplacer) Len() int {
	return len(repl.replacements)
}

func (repl *MultiReplacer) replacementSliceAt(n int) []string {
	if n < 0 || n > len(repl.replacements) {
		return nil
	}
	if !repl.withDelimiters() {
		return util.UnwrapMap(repl.replacements[n])
	}
	res := make([]string, 0, len(repl.replacements[n])*2)
	for key, val := range repl.replacements[n] {
		res = append(
			res,
			strings.Join([]string{repl.startDelimiter, key, repl.endDelimiter}, ""),
			val,
		)
	}
	return res
}

func (repl *MultiReplacer) executeToString(template string, n int) string {
	var buf bytes.Buffer
	_, err := repl.executeToStream(template, n, &buf)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (repl *MultiReplacer) executeToStream(template string, n int, w io.Writer) (int, error) {
	replacer := strings.NewReplacer(repl.replacementSliceAt(n)...)
	bytesWritten, err := replacer.WriteString(w, template)
	if err != nil {
		return bytesWritten, err
	}
	return bytesWritten, nil
}

func (repl *MultiReplacer) withDelimiters() bool {
	return !(repl.startDelimiter == "" && repl.endDelimiter == "")
}
