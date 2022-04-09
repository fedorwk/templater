package templater

import "io"

type Templater struct {
	Template string
	Replacer *Replacer
}

func NewTemplater(Template string, items []map[string]string,
	startDelimiter string, endDelimiter string) *Templater {

	Replacer := NewReplacer(items, startDelimiter, endDelimiter)
	return &Templater{Template, Replacer}
}

// AllStrings() returns list of strings with replacements executed
func (t *Templater) AllStrings() []string {
	res := make([]string, 0, t.Replacer.Len())

	for i := 0; i < t.Replacer.Len(); i++ {
		res = append(res, t.Replacer.executeToString(t.Template, i))
	}
	return res
}

// returns string with executed replacement with data at given index in MultiReplacer
func (t *Templater) ExecuteToString(i int) string {
	return t.Replacer.executeToString(t.Template, i)
}

// writes text with executed replacement with item data at given index in MultiReplacer
// returns bytes written and error if occurred
func (t *Templater) ExecuteToStream(i int, w io.Writer) (int, error) {
	return t.Replacer.executeToStream(t.Template, i, w)
}
