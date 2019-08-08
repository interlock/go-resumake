package templates

import (
	"github.com/interlock/go-resumake/jsonresume"
	"errors"
	"bytes"
	"github.com/interlock/go-resumake/templates/awesome-cv"
)

// Temporary until we breake templates into exec

func Render(template string, data jsonresume.JSONResume) (*bytes.Buffer, error) {
	var output *bytes.Buffer = bytes.NewBuffer([]byte{})
	var err error

	switch(template) {
	case "awesome-cv":
		err = awesomecv.Render(data, output)
		if (err != nil) {
			return nil, err
		}
		break
	default:
		return nil, errors.New("Unknown template")
	}
	return output, nil
}