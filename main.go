package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/interlock/go-resumake/templates"
)

func main() {
	flag.Parse()
	jsonResume, err := decodeFile(flagJsonResume)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if flagVerbose {
		fmt.Printf("%s\n", jsonResume)
	}
	if strings.Compare(flagTemplate, "") == 0 {
		fmt.Println("Template was empty")
		os.Exit(3)
	}
	output, err := templates.Render(flagTemplate, jsonResume)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	if len(flagLatex) > 0 {
		fp, err := os.Create(flagLatex)
		if err != nil {
			fmt.Println("%s\n", err)
			os.Exit(5)
		}

		_, err = fp.Write(output.Bytes())
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(6)
		}
		fp.Close()
	} else {
		fmt.Printf("%s\n", output)
	}
}
