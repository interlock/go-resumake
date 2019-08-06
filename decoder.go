package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/interlock/go-resumake/jsonresume"
)

func decodeFile(path string) (*jsonresume.JSONResume, error) {
	f, err := os.Open(path)
	if (err != nil) {
		return nil, err
	}
	defer f.Close()

	fBytes, err := ioutil.ReadAll(f)
	if (err != nil) {
		return nil, err
	}

	var jsonResume jsonresume.JSONResume

	json.Unmarshal(fBytes, &jsonResume)

	return &jsonResume, nil
}