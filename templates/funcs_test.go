package templates

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewSlice(t *testing.T) {
	s := newSlice()
	if len(s) != 0 {
		t.Fatal("expected empty slice")
	}
}

func TestAppendSlice(t *testing.T) {
	s := appendSlice([]string{}, "element")
	if len(s) != 1 {
		t.Fatal("expected their to be one element")
	}
	if strings.Compare(s[0], "element") != 0 {
		t.Fatal("expected first item to be element")
	}
}

func TestEscapeLatex(t *testing.T) {
	var checks map[string]string = map[string]string{
		"&": "\\&",
		"%": "\\%",
		"$": "\\$",
		"#": "\\#",
		"_": "\\_",
		"{": "\\{",
		"}": "\\}",
		"~": "{\\raise.17ex\\hbox{$\\scriptstyle\\mathtt{\\sim}$}}",
		"^": "\\char`\\^",
	}
	for v, s := range checks {
		e := escapeLatex(v)
		if strings.Compare(e, s) != 0 {
			t.Fatalf("improper escape for %s, got %s", v, e)
		}
	}
}

func TestDateFormat(t *testing.T) {
	var date string = "2021-11-12"
	var checks map[string]string = map[string]string{
		"2006": "2021",
		"01":   "11",
		"02":   "12",
	}
	for r, v := range checks {
		d := dateFormat(date, r)
		if d != v {
			t.Fatalf("expected format %s to be %s but got %s", r, v, d)
		}
	}
}

func TestDateFormatInvalid(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected a panic")
		}

		e := "parsing time \"2021-12-12T12:12:12Z\": extra text: \"T12:12:12Z\""
		if strings.Compare(fmt.Sprintf("%s", r), e) != 0 {
			t.Fatalf("panic message was unexpected:\n%s\n%s", fmt.Sprintf("%s", r), e)
		}
	}()
	dateFormat("2021-12-12T12:12:12Z", "2006")
}

func TestFuncsMap(t *testing.T) {
	if len(funcsMap()) != 5 {
		t.Fatalf("expected 5 funcs to map to templates, go %d", len(funcsMap()))
	}
}
