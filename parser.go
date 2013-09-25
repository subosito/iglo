package iglo

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os/exec"
)

func ParseJSON(r io.Reader) (*API, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	api := new(API)
	err = json.Unmarshal(b, &api)
	if err != nil {
		return nil, err
	}

	return api, nil
}

func ParseMarkdown() ([]byte, error) {
	path, err := exec.LookPath("snowcrash")
	if err != nil {
		return nil, errors.New("Couldn't find snowcrash. Please install it first https://github.com/apiaryio/snowcrash")
	}

	cmd := exec.Command(path, "--format", "json", "API.md")

	return cmd.Output()
}
