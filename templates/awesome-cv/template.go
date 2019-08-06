package awesomecv

import (
	"github.com/interlock/go-resumake/jsonresume"
	"io"
)

func Render(resume jsonresume.JSONResume, output io.Writer) error {
	output.Write([]byte("meow\n"))
	return nil
}