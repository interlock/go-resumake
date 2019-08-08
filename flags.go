package main

import "flag"

var (
	flagJsonResume string
	flagVerbose bool
	flagTemplate string
	flagLatex string
)

func init() {
	flag.StringVar(&flagJsonResume, "input", "", "input json resume")
	flag.BoolVar(&flagVerbose, "verbose", false, "enable verbose")
	flag.StringVar(&flagTemplate, "template", "", "template")
	flag.StringVar(&flagLatex, "latex", "", "latex")
}