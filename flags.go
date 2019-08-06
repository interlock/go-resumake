package main

import "flag"

var (
	flagJsonResume string
	flagVerbose bool
	flagTemplate string
	flagLatex bool
)

func init() {
	flag.StringVar(&flagJsonResume, "input", "", "input json resume")
	flag.BoolVar(&flagVerbose, "verbose", false, "enable verbose")
	flag.StringVar(&flagTemplate, "template", "", "template")
	flag.BoolVar(&flagLatex, "latex", false, "latex")
}