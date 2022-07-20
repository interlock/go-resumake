package templates

import (
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"time"
)

func newSlice() []string {
	return make([]string, 0)
}

func appendSlice(in []string, v string) []string {
	return append(in, v)
}

// escapeLatex
// & % $ # _ { } ~ ^ \
func escapeLatex(s string) string {
	re := regexp.MustCompile(`[&$%#_\{\}~^\\]`)
	return re.ReplaceAllStringFunc(s, func(str string) string {
		if strings.Compare(str, "\\") == 0 {
			return "\\textbackslash{}"
		}
		if strings.Compare(str, "~") == 0 {
			// return "\\char`\\~"
			return "{\\raise.17ex\\hbox{$\\scriptstyle\\mathtt{\\sim}$}}"
		}
		if strings.Compare(str, "^") == 0 {
			return "\\char`\\^"
		}
		return fmt.Sprintf("\\%s", str)
	})
}

func dateFormat(d string, f string) string {
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		panic(err)
	}
	return t.Format(f)
}

func funcsMap() map[string]any {
	return template.FuncMap{
		"join":        strings.Join,
		"newSlice":    newSlice,
		"appendSlice": appendSlice,
		"escape":      escapeLatex,
		"dateFormat":  dateFormat,
	}
}
