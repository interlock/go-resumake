package templates

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/interlock/go-resumake/jsonresume"
	awesomecv "github.com/interlock/go-resumake/templates/awesome-cv"
	packedcv "github.com/interlock/go-resumake/templates/packed-cv"
)

// Temporary until we breake templates into exec

func Render(tmplName string, data *jsonresume.JSONResume) (*bytes.Buffer, error) {
	var output *bytes.Buffer = bytes.NewBuffer([]byte{})
	var err error
	tmpl := template.New("document").Funcs(funcsMap())
	switch tmplName {
	case "awesome-cv":
		err = awesomecv.Render(tmpl, data, output)
		if err != nil {
			return nil, err
		}
		break
	case "packed-cv":
		err = packedcv.Render(tmpl, data, output)
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, errors.New("Unknown template")
	}
	return output, nil
}
