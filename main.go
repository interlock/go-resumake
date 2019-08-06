package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	jsonResume, err := decodeFile(flagJsonResume)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(2)
	}
	if (flagVerbose) {
		fmt.Printf("%s\n", jsonResume)
	}
	if (strings.Compare(flagTemplate, "") == 0) {
		fmt.Println("Template was empty")
		os.Exit(3)
	}
}