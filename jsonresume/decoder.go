package jsonresume

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func DecodeFile(path string) (jsonResume *JSONResume, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	jsonResume = &JSONResume{}
	err = json.Unmarshal(fBytes, jsonResume)
	if err != nil {
		return nil, err
	}

	return jsonResume, nil
}
