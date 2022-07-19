package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/interlock/go-resumake/jsonresume"
)

func decodeFile(path string) (jsonresume.JSONResume, error) {
	var jsonResume jsonresume.JSONResume

	f, err := os.Open(path)
	if err != nil {
		return jsonResume, err
	}
	defer f.Close()

	fBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return jsonResume, err
	}

	json.Unmarshal(fBytes, &jsonResume)

	return jsonResume, nil
}
