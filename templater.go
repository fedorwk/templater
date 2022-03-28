package templater

import "io"

type Templater struct {
	template string
	replacer *MultiReplacer
}

func NewTemplater(template string, replacer *MultiReplacer) *Templater {
	return &Templater{template, replacer}
}

// AllStrings() returns list of strings with replacements executed
func (t *Templater) AllStrings() []string {
	res := make([]string, 0, len(t.replacer.replacements))

	for i := 0; i < t.replacer.Len(); i++ {
		res = append(res, t.replacer.executeToString(t.template, i))
	}
	return res
}

// returns string with executed replacement with data at given index in MultiReplacer
func (t *Templater) ExecuteToString(i int) string {
	return t.replacer.executeToString(t.template, i)
}

// writes text with executed replacement with data at given index in MultiReplacer
// returns bytes written and error if occurred
func (t *Templater) ExecuteToStream(i int, w io.Writer) (int, error) {
	return t.replacer.executeToStream(t.template, i, w)
}
