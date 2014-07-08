package iglo

import (
	"encoding/json"
	"errors"
	"fmt"
	version "github.com/hashicorp/go-version"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
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
	path, err := snowcrash()
	if err != nil {
		return nil, err
	}

	err = detectVersion()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	echo := exec.Command("echo", string(b))
	out, err := echo.StdoutPipe()
	if err != nil {
		return nil, err
	}

	echo.Start()

	cmd := exec.Command(path, "--format", "json")
	cmd.Stdin = out

	return cmd.Output()
}

func CheckVersion(v string) error {
	mv, _ := version.NewVersion("0.11.0")
	ov, err := version.NewVersion(v)
	if err != nil {
		return err
	}

	if ov.LessThan(mv) {
		return errors.New(fmt.Sprintf("You are using snowcrash version %s. Minimum version should be %s", ov, mv))
	}

	return nil
}

func detectVersion() error {
	v, err := snowcrashVersion()
	if err != nil {
		return err
	}

	err = CheckVersion(v)
	if err != nil {
		return err
	}

	return nil
}

func snowcrash() (string, error) {
	path, err := exec.LookPath("snowcrash")
	if err != nil {
		return "", errors.New("Couldn't find snowcrash. Please install it first https://github.com/apiaryio/snowcrash")
	}

	return path, nil
}

func snowcrashVersion() (string, error) {
	path, err := snowcrash()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(path, "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.Replace(string(output), "v", "", 1)), nil
}
