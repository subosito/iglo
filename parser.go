package iglo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"

	version "github.com/hashicorp/go-version"
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

func ParseMarkdown(r io.Reader) ([]byte, error) {
	path, err := drafter()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewReader(b)
	cmd := exec.Command(path, "--format", "json", "--type", "ast")
	cmd.Stdin = buf

	return cmd.Output()
}

func CheckVersion(v string) error {
	mv, _ := version.NewVersion("1.0.0")
	ov, err := version.NewVersion(v)
	if err != nil {
		return err
	}

	if ov.LessThan(mv) {
		return errors.New(fmt.Sprintf("You are using drafter version %s. Minimum version should be %s", ov, mv))
	}

	return nil
}

func drafter() (string, error) {
	path, err := exec.LookPath("drafter")
	if err != nil {
		return "", errors.New("Couldn't find drafter. Please install it first https://github.com/apiaryio/drafter")
	}

	return path, nil
}
